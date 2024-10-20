package srvalm

import (
	"fmt"
	"log"
	"os"
	"servistas/v2/internal/gvcomun"
	"servistas/v2/internal/msgsys"
	"time"

	// paquetes adicionales si necesario
)

type OperacionPedida struct{
	operacion string
	clave string
	valor string
}

type ServAlm struct {
	msgsys.MsgSys
	gv msgsys.HostPuerto
	doneTicker chan struct{}

	vista gvcomun.Vista

	comunicacion chan  string
	info map[string] string
	ultimaOperacionPedida OperacionPedida
	estado string //estado primario/copia/espera
	// Campos adicionales del estado de información que debe guardar
	// este servidor de almacenamiento ???
}

func Make(me, gv msgsys.HostPuerto) ServAlm {
	// poner microsegundos y PATH completo de fichero codigo en logs
	gvcomun.ChangeLogPrefix()

	// Registrar tipos de mensaje de gestión d vistas
	gvcomun.RegistrarMensajesSA()

	// crear estructura datos estado del servidor
	return ServAlm{
		// Crear channel msgsys
		MsgSys: msgsys.MakeMsgSys(me),
		gv:     gv,
		doneTicker: make(chan struct{}),

		vista: gvcomun.Vista{0, msgsys.HOSTINDEFINIDO, msgsys.HOSTINDEFINIDO},

		info: make(map[string]string),
		ultimaOperacionPedida: OperacionPedida{"", "", ""},
		estado: "ESPERA",

		comunicacion: make(chan string),

		//TODO INICIALIZAR
	}
}

// Poner en marcha el servidor de vistas/gestor de vistas
func (sa *ServAlm) Start() {

	gvcomun.ChangeLogPrefix()
	log.Println("Arrancando servidor_almacenamiento: ", sa.Me())

	go sa.ticker()

	// Tratar mensajes de clientes
	// Termina ejecución cuando recibe mensaje tipo MsgFin.
	// Este metodo esta en paquete msgsys
	sa.ProcessAllMsg(sa.procesaMensaje)
}

func (sa *ServAlm) Stop() {

	close(sa.doneTicker)

	// Esperar que goroutina ticker llegue a done por el time.Sleep
	time.Sleep(2 * gvcomun.INTERVALOLATIDOS * time.Millisecond)
	sa.CloseMessageSystem()
}

func (sa *ServAlm) procesaMensaje(m msgsys.Message) {
	//log.Printf("----Recibido mensaje : %#v\n", m)
	switch x := m.(type) {
	case gvcomun.MsgVistaTentativa:
		sa.cambiaVista(x)

	case gvcomun.MsgWrite:
		log.Println("Recibido Write")
		if x.Clave != sa.ultimaOperacionPedida.clave || 				//si no es la misma operación de antes
					sa.ultimaOperacionPedida.operacion != "WRITE" ||
					x.Valor != sa.ultimaOperacionPedida.valor {
			go sa.procesaWrite(x)
		}

	case gvcomun.MsgRead:
		log.Println("Recibido Read")
		if x.Clave != sa.ultimaOperacionPedida.clave ||					//si no es la misma operación de antes
					sa.ultimaOperacionPedida.operacion != "READ" ||
					"" != sa.ultimaOperacionPedida.valor {
			sa.procesaRead(x)
		}
	case gvcomun.MsgRespuestaOperacion:
		log.Println("Recibido RO copia ")
		sa.procesaRespuesta(x)
	case gvcomun.MsgWriteTotal:
		sa.procesaWriteTotal(x)
	case gvcomun.MsgTickInterno:
		sa.procesaTickInterno()

	case gvcomun.MsgFin:
		log.Println("Recibido FIN")
		sa.Stop() // Eliminar el servidor de almacenamiento !!
		os.Exit(0)
	default:
		log.Printf(
			"Llega mensaje TIPO DESCONOCIDO en servidor_vistas.go!! %#v\n", m)
	}

}


// salida en el mensaje del tipo MsgRespuestaOperacion

func (sa *ServAlm) cambiaVista(x gvcomun.MsgVistaTentativa) {

	if(sa.MsgSys.Me() == x.Vista.Primario){	//si soy el primario

		sa.estado = "PRIMARIO"
		if(x.Vista.Copia != msgsys.HOSTINDEFINIDO && sa.vista.NumVista != x.Vista.NumVista) { 	//si tengo que replicar

			sa.MsgSys.Send(x.Vista.Copia, gvcomun.MsgWriteTotal{Mapa: sa.info, Remitente: sa.Me()})
		}
	}else if(sa.MsgSys.Me() == x.Vista.Copia){	//si soy copia

		sa.estado = "COPIA"

		//fmt.Println("Me he convertido en copia.")

	}else{

		sa.estado = "ESPERA"

		//fmt.Println("Me he convertido en espera.")

	}
	if(x.Vista.NumVista != 1) {
		sa.vista = x.Vista
	}else{	//vista == 1 envio -1 para no confirmar
		sa.vista.NumVista = -1
	}

}

func (sa *ServAlm) procesaWrite(x gvcomun.MsgWrite) {
	var mesg string

	if(sa.estado == "PRIMARIO"){
		if(sa.vista.Copia != msgsys.HOSTINDEFINIDO) {

			sa.Send(sa.vista.Copia, gvcomun.MsgWrite{Clave: x.Clave, Valor: x.Valor, Remitente: sa.MsgSys.Me()}) //propaga el cambio
			fmt.Println(sa.vista.Copia)
			select {
			 case mesg = <-sa.comunicacion:
				if (mesg == "HECHO"){

					sa.info[x.Clave] = x.Valor                                                              //realiza el cambio
					sa.Send(x.Remitente, gvcomun.MsgRespuestaOperacion{Valor: x.Valor, Remitente: sa.Me()}) //devuelve el valor
					sa.ultimaOperacionPedida = OperacionPedida{	//se apunta la última operación
						operacion: "WRITE",
						clave:     x.Clave,
						valor:     x.Valor,
					}
				}
			case <-time.After(2*gvcomun.INTERVALOLATIDOS * time.Millisecond):
				mesg = "nil"
				sa.Send(x.Remitente, gvcomun.MsgRespuestaOperacion{Valor: "FALLO", Remitente: sa.Me()})
			}
			fmt.Println(mesg)

		}

	}else if(sa.estado == "COPIA"){
		sa.info[x.Clave] = x.Valor		//realiza el cambio
		sa.Send(x.Remitente,  gvcomun.MsgRespuestaOperacion{Valor: "HECHO", Remitente: sa.Me()})	//confirmación
		fmt.Println("hago Write enviado por primario: ", x.Remitente)
	}
}

func (sa *ServAlm) procesaRead(x gvcomun.MsgRead) {

	if(sa.estado == "PRIMARIO") {
		sa.Send(x.Remitente, gvcomun.MsgRespuestaOperacion{Valor: sa.info[x.Clave], Remitente: sa.Me()}) //devuelve el contenido
	}
	sa.ultimaOperacionPedida = OperacionPedida{	//se apunta la última operación
		operacion: "READ",
		clave:     x.Clave,
		valor:     "",
	}
}


func (sa *ServAlm) procesaTickInterno() {
	sa.MsgSys.Send(sa.gv, gvcomun.MsgLatido{sa.vista.NumVista, sa.MsgSys.Me()})
}

// ticker to run through all the execution
func (sa *ServAlm) ticker() {
iteracion:
	for {
		select {
		case <-sa.doneTicker:
			//log.Println("Done with receiving Messages from TICKER!!!")
			break iteracion // Deja ya de generar  Ticks !!!

		default:
			time.Sleep(gvcomun.INTERVALOLATIDOS/3 * time.Millisecond)
			sa.InternalSend(gvcomun.MsgTickInterno{})
		}
	}
}

func (sa *ServAlm) procesaWriteTotal(x gvcomun.MsgWriteTotal) {
	sa.info = x.Mapa
	fmt.Println("procesado write total")
}

func (sa *ServAlm) procesaRespuesta(x gvcomun.MsgRespuestaOperacion) {
	sa.comunicacion <- x.Valor
	fmt.Println("Write realizado en copia")
}



