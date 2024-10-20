/*
Package srvvts implementa servicio de vistas para gestión de réplicas
*/
package srvvts

import (
	"log"
	"replicasys/internal/comun"
	"replicasys/internal/msgsys"
	"time"
	// paquetes adicionales si necesario
)

type ServVistas struct {
	msgsys.MsgSys
	doneTicker chan struct{}

	// Campos adicionales del estado de información que debe guardar
	// este gestor de vistas
}

func Make(me msgsys.HostPuerto) ServVistas {
	// poner microsegundos y PATH completo de fichero codigo en logs
	comun.ChangeLogPrefix()

	// Registrar tipos de mensaje de gestión d vistas
	comun.RegistrarMensajesGV()

	// crear estructura datos estado del servidor
	return ServVistas{
		// Crear channel msgsys
		MsgSys:     msgsys.MakeMsgSys(me),
		doneTicker: make(chan struct{}),

		// Other fields

	}
}

// Poner en marcha el servidor de vistas/gestor de vistas
func (sv *ServVistas) Start() {

	comun.ChangeLogPrefix()
	log.Println("Arrancando gestor de vistas: ", sv.Me())

	// crea generador de eventos ticker en periodos de envio de latido
	go sv.ticker()

	// Tratar mensajes de clientes y ticks periodicos de procesado de situacion
	// Termina ejecución cuando recibe mensaje tipo MsgFin.
	// Este metodo esta en paquete msgsys
	sv.ProcessAllMsg(sv.procesaMensaje)
}

func (sv *ServVistas) Stop() {
	// primero cerrar ticker para que pare de enviar MsgTickInterno a buzon
	close(sv.doneTicker)

	// Esperar que goroutina ticker llegue a done por el time.Sleep
	time.Sleep(2 * comun.INTERVALOLATIDOS * time.Millisecond)

	//
	sv.CloseMessageSystem()
}

// ticker to run through all the execution
func (sv *ServVistas) ticker() {
iteracion:
	for {
		select {
		case <-sv.doneTicker:
			log.Println("Done with receiving Messages from TICKER!!!")
			break iteracion // Deja ya de generar  Ticks !!!

		default:
			time.Sleep(comun.INTERVALOLATIDOS * time.Millisecond)
			sv.InternalSend(comun.MsgTickInterno{})
		}
	}
}

func (sv *ServVistas) procesaMensaje(m msgsys.Message) {
	// log.Printf("----Recibido mensaje : %#v\n", m)

	// Completar
}
