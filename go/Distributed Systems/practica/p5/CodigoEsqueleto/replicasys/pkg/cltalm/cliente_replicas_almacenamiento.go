/*
Package cltalm : Cliente de almacenamiento local con msgsys y puerto único
*/
package cltalm

import (
	"log"
	"replicasys/internal/comun"
	"replicasys/internal/msgsys"
)

type ClienteAlm struct {
	msgsys.MsgSys
	me, gv, primario msgsys.HostPuerto
}

func Start(me, gv msgsys.HostPuerto) ClienteAlm {
	// Registrar tipos de mensaje de gestión d vistas
	comun.RegistrarMensajesSA()

	return ClienteAlm{
		MsgSys:   msgsys.MakeMsgSys(me),
		gv:       gv,
		primario: msgsys.HOSTINDEFINIDO,
	}
}
func (ca *ClienteAlm) Stop() {
	ca.CloseMessageSystem()
	log.Println("Cliente de almacenamiento DETENIDO : ", ca.Me())
}

func (ca *ClienteAlm) GetPrimario() msgsys.HostPuerto {

	// Completar

	return ca.primario
}

// escritura a modificar para fallos de comunicaciones y operaciones
func (ca *ClienteAlm) Write(clave, valor string) (res string, ok bool) {

	// Completar
	return res, ok
}

func (ca *ClienteAlm) Read(clave string) (res string, ok bool) {

	// Completar
	return res, ok
}
