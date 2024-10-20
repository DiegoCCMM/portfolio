package srvalm

import (
	"log"
	"replicasys/internal/comun"
	"replicasys/internal/msgsys"
	// paquetes adicionales si necesario
)

type ServAlm struct {
	msgsys.MsgSys
	gv msgsys.HostPuerto

	// Campos adicionales del estado de información que debe guardar
	// este servidor de almacenamiento ???
}

func Make(me, gv msgsys.HostPuerto) ServAlm {
	// poner microsegundos y PATH completo de fichero codigo en logs
	comun.ChangeLogPrefix()

	// Registrar tipos de mensaje de gestión d vistas
	comun.RegistrarMensajesSA()

	// crear estructura datos estado del servidor
	return ServAlm{
		// Crear channel msgsys
		MsgSys: msgsys.MakeMsgSys(me),
		gv:     gv,

		// campos adicionales ???

	}
}

// Poner en marcha el servidor de vistas/gestor de vistas
func (sa *ServAlm) Start() {

	comun.ChangeLogPrefix()
	log.Println("Arrancando servidor_almacenamiento: ", sa.Me())

	// Código adicional si necesario ...........................

	// Tratar mensajes de clientes
	// Termina ejecución cuando recibe mensaje tipo MsgFin.
	// Este metodo esta en paquete msgsys
	sa.ProcessAllMsg(sa.procesaMensaje)
}

func (sa *ServAlm) Stop() {

	// Código adicional si necesario .......................

	sa.CloseMessageSystem()
}

func (sa *ServAlm) procesaMensaje(m msgsys.Message) {
	//log.Printf("----Recibido mensaje : %#v\n", m)

}
