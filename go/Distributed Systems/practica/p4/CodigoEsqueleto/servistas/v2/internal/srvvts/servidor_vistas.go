/*
   Package srvvts implementa servicio de vistas para gestión de réplicas
*/
package srvvts

import (
	"fmt"
	"log"
	"os"
	"servistas/v2/internal/gvcomun"
	"servistas/v2/internal/msgsys"
	"time"
	// paquetes adicionales si necesario
)

type ServVistas struct {
	msgsys.MsgSys
	doneTicker chan struct{}

	actual gvcomun.Vista
	tentativa  gvcomun.Vista

	espera [] msgsys.HostPuerto
	hanLatido map[msgsys.HostPuerto] int
	//-1 ha latido, 0 no ha latido, > 0
	//no ha latido durante x intentos
	//(hasta gvcomun.LATIDOSFALLIDOS)

	hemosEmpezadoATrabajar bool


}

func Make(me msgsys.HostPuerto) ServVistas {
	// poner microsegundos y PATH completo de fichero codigo en logs
	gvcomun.ChangeLogPrefix()

	//log.Println("Puesta en marcha de servidor_gv")

	// Registrar tipos de mensaje de gestión d vistas
	gvcomun.RegistrarTiposMensajesGV()

	// crear estructura datos estado del servidor
	return ServVistas{
		// Crear channel msgsys
		MsgSys:     msgsys.MakeMsgSys(me),
		doneTicker: make(chan struct{}),

		actual: gvcomun.Vista{0, msgsys.HOSTINDEFINIDO, msgsys.HOSTINDEFINIDO},
		tentativa: gvcomun.Vista{0, msgsys.HOSTINDEFINIDO, msgsys.HOSTINDEFINIDO},

		espera: []msgsys.HostPuerto{},
		hanLatido: make(map[msgsys.HostPuerto]int),
		hemosEmpezadoATrabajar: false,
	}
}

// Poner en marcha el servidor de vistas/gestor de vistas
func (sv *ServVistas) Start() {
	// crea generador de eventos ticker en periodos de envio de latido
	go sv.ticker()

	// Tratar mensajes de clientes y ticks periodicos de procesado de situacion
	// Termina ejecución cuando recibe mensaje tipo MsgFin.
	sv.ProcessAllMsg(sv.procesaMensaje) // Este metodo esta en paquete "msgsys"
}

func (sv *ServVistas) Stop() {
	// primero cerrar terminación (done) para que ticker
	// pare de enviar MsgTickInterno a buzon
	close(sv.doneTicker)

	// Esperar que goroutina ticker llegue a done por el time.Sleep
	time.Sleep(2 * gvcomun.INTERVALOLATIDOS * time.Millisecond)

	//
	sv.CloseMessageSystem()
}

// ticker to run through all the execution
func (sv *ServVistas) ticker() {
iteracion:
	for {
		select {
		case <-sv.doneTicker:
			//log.Println("Done with receiving Messages from TICKER!!!")
			break iteracion // Deja ya de generar  Ticks !!!

		default:
			time.Sleep(gvcomun.INTERVALOLATIDOS * time.Millisecond)
			sv.InternalSend(gvcomun.MsgTickInterno{})
		}
	}
}

func (sv *ServVistas) procesaMensaje(m msgsys.Message) {
	//log.Printf("----Recibido mensaje : %#v\n", m)

	switch x := m.(type) {
	case gvcomun.MsgLatido:
		fmt.Println("**********\n recibido latido ", sv.tentativa.NumVista, sv.tentativa.Primario, sv.tentativa.Copia )
		sv.trataLatido(x)
	case gvcomun.MsgPeticionVistaValida:
		log.Println("Recibido peticion vista valida")
		sv.Send(x.Remitente, gvcomun.MsgVistaValida{sv.actual})
	case gvcomun.MsgPeticionPrimario:
		sv.trataPeticionPrimario(x.Remitente)
	case gvcomun.MsgTickInterno:
		sv.procesaSituacionReplicas()
	case gvcomun.MsgFin:
		log.Println("Recibido FIN")
		sv.Stop() // Eliminar el servidor de vistas !!
		os.Exit(0)
	default:
		log.Printf(
			"Llega mensaje TIPO DESCONOCIDO en servidor_vistas.go!! %#v\n", m)
	}
}

func (sv *ServVistas) trataLatido(x gvcomun.MsgLatido) {
	fmt.Println("El num de vista recibido de "+ x.Remitente +" es ", x.NumVista)
	if x.NumVista == 0{
		//perdida de memoria o conexión de primeras o se ha perdido la conexión y reconectamos que es lo mismo
		if x.Remitente == sv.tentativa.Primario || x.Remitente == sv.tentativa.Copia{
			//si uno de estos estaba considerado primario o copia, hemos perdido la memoria
			if x.Remitente == sv.tentativa.Primario{
				sv.caidaPrimario()
			}else{sv.caidaCopia()}

		}else if sv.tentativa.Primario == msgsys.HOSTINDEFINIDO {	//aún no estamos inicializados
			sv.tentativa.NumVista++
			sv.tentativa.Primario = x.Remitente
		}else if sv.tentativa.Copia == msgsys.HOSTINDEFINIDO && x.Remitente != sv.tentativa.Primario {
			sv.tentativa.NumVista++
			sv.tentativa.Copia = x.Remitente
		}else{	//estamos inicializados, a espera
			sv.espera = append(sv.espera, x.Remitente)
		}

	}else if (x.Remitente == sv.tentativa.Primario && sv.tentativa.Copia != msgsys.HOSTINDEFINIDO && x.NumVista == sv.tentativa.NumVista && sv.tentativa != sv.actual) {
		sv.actual = sv.tentativa
		if !sv.hemosEmpezadoATrabajar {
			sv.hemosEmpezadoATrabajar = true
		}
	}
	//en cualquier caso, apuntamos su latido y reenviamos
	sv.hanLatido[x.Remitente] = -1
	sv.Send(x.Remitente, gvcomun.MsgVistaTentativa{sv.tentativa})
}

func (sv *ServVistas) trataPeticionPrimario(remitente msgsys.HostPuerto) {
	sv.Send(remitente, gvcomun.MsgPrimario(sv.actual.Primario))
}

func (sv *ServVistas) procesaSituacionReplicas() {//recibo un tick interno, o lo que es lo mismo, han pasado 50 ms
	if(sv.tentativa.Primario != msgsys.HOSTINDEFINIDO) {
		sv.cambioEstado(sv.tentativa.Primario, -1)
	}
	if(sv.tentativa.Copia != msgsys.HOSTINDEFINIDO) {
		sv.cambioEstado(sv.tentativa.Copia, -2)
	}

	for i := 0; i < len(sv.espera); i++ {
		sv.cambioEstado(sv.espera[i], i)
	}

}

func (sv *ServVistas) cambioEstado(puerto msgsys.HostPuerto, i int ) {
	if sv.hanLatido[puerto] != -1 { //no ha latido

		if sv.hanLatido[puerto] < gvcomun.LATIDOSFALLIDOS - 1{ // no es su cuarta espera
			sv.hanLatido[puerto] ++

		} else { //es su cuarta espera
			if puerto == sv.tentativa.Primario {
				fmt.Println("se ha muerto el primario")
				sv.caidaPrimario()
			}else if puerto == sv.tentativa.Copia {
				fmt.Println("se ha muerto el copia")
				sv.caidaCopia()
			}else {	//se ha caido un nodo en espera
				sv.espera = append(sv.espera[:i], sv.espera[i+1:]...)
			}

		}
	} else { //era -1, ha latido, lo ponemos a 0
		sv.hanLatido[puerto] = 0
	}
}

func (sv *ServVistas) caidaPrimario() {
	if sv.tentativa == sv.actual  || !sv.hemosEmpezadoATrabajar {
		sv.tentativa.NumVista++
		sv.tentativa.Primario = sv.tentativa.Copia
		fmt.Println("ahora primario es "+sv.tentativa.Copia)
		if len(sv.espera) > 0{
			sv.tentativa.Copia = sv.espera[0]
			fmt.Println("ahora copia es "+sv.espera[0])
			sv.espera = append(sv.espera[:0], sv.espera[1:]...)
		}else{
			sv.tentativa.Copia = msgsys.HOSTINDEFINIDO
		}
	}else {
		fmt.Println("Error crítico: Hemos perdido los datos de un primario sin réplica")
		sv.Stop()
		os.Exit(1)
	}
}

func (sv *ServVistas) caidaCopia() {
	if len(sv.espera) > 0 {
		sv.tentativa.NumVista++
		sv.tentativa.Copia = sv.espera[0]
		sv.espera = append(sv.espera[:0], sv.espera[1:]...)
		fmt.Println("ahora copia es "+sv.tentativa.Copia)
	}else{
		sv.tentativa.Copia = msgsys.HOSTINDEFINIDO
	}
}

