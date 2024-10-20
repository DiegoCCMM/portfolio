/*
Package cltalm : Cliente de almacenamiento local con msgsys y puerto único
*/
package cltalm

import (
	"fmt"
	"log"
	"servistas/v2/internal/gvcomun"
	"servistas/v2/internal/msgsys"
	"time"
)

type ClienteAlm struct {
	msgsys.MsgSys
	me, gv, primario msgsys.HostPuerto
}

func Start(me, gv msgsys.HostPuerto) ClienteAlm {
	// Registrar tipos de mensaje de gestión d vistas
	gvcomun.RegistrarMensajesSA()

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

	ca.Send(ca.gv, gvcomun.MsgPeticionPrimario{
		Remitente: ca.Me(),
	})

	re := ca.Receive()

	switch x := re.(type) {
		case gvcomun.MsgPrimario:
			ca.primario = msgsys.HostPuerto(x)
	}

	return ca.primario
}
func (ca *ClienteAlm)ObtenerPrimario(){
	fmt.Println("El valor de primario de este cliente es: ", ca.primario)
}
// escritura a modificar para fallos de comunicaciones y operaciones
func (ca *ClienteAlm) Write(clave, valor string) (res string, ok bool) {

	re, ok := ca.SendReceive(ca.primario, gvcomun.MsgWrite{
		Clave:     clave,
		Valor:     valor,
		Remitente: ca.Me(),
	},
		gvcomun.ANSWERWAITTIME*time.Millisecond,
	)

	switch x := re.(type) {
	case gvcomun.MsgRespuestaOperacion:
		res = x.Valor
	}

	return res, ok
}

func (ca *ClienteAlm) Read(clave string) (res string, ok bool) {

	re, ok := ca.SendReceive(ca.primario, gvcomun.MsgRead{
		Clave:     clave, Remitente: ca.Me(),
	},
		gvcomun.ANSWERWAITTIME*time.Millisecond,
	)
	switch x := re.(type) {
	case gvcomun.MsgRespuestaOperacion:
		res = x.Valor
	}

	return res, ok
}
