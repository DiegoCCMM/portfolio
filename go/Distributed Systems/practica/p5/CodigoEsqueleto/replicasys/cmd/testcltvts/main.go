/*
Ejecutable testcltvts solo para pruebas, únicamente en práctica 4,
   de un cliente muy baśico del servico de vistas
*/
package main

import (
	"fmt"
	"log"
	"os"
	"replicasys/internal/comun"
	"replicasys/internal/msgsys"
	"replicasys/pkg/cltvts"
)

func main() {
	// obtener host:puerto de este clientes desde argumentos de llamada
	me := os.Args[1]
	// obtener host:puerto del gestor de vistas desde argumentos de llamada
	gv := os.Args[2]
	// obtener host:puerto del nodo de tests
	nodotest := os.Args[3]

	// Start a test client and get its mailbox channel
	StartClienteVistas(msgsys.HostPuerto(me),
		msgsys.HostPuerto(gv), msgsys.HostPuerto(nodotest))
}

type clienteTest struct {
	cltvts.CltVts // tipo abstracto de cliente de gestor de vistas
	msgsys.MsgSys
	nodotest msgsys.HostPuerto
}

func StartClienteVistas(me, gv, nodotest msgsys.HostPuerto) {

	comun.ChangeLogPrefix()
	log.Println("Arrancando ClienteTestVistas: ", me)

	// Registrar tipos de mensaje de gestión d vistas
	comun.RegistrarMensajesGV()

	// Crear el cliente de test
	clt := clienteTest{
		CltVts:   cltvts.CltVts{Gv: gv},
		MsgSys:   msgsys.MakeMsgSys(me),
		nodotest: nodotest,
	}

	// Tratar mensajes de clientes para reencaminar a Servidor de Vistas
	// Termina ejecución cuando recibe mensaje tipo MsgFin
	clt.ProcessAllMsg(clt.procesaMensaje)
}

func (cl *clienteTest) stop() {
	cl.CloseMessageSystem()
}

func (cl *clienteTest) procesaMensaje(m msgsys.Message) {
	//fmt.Printf("----Recibido mensaje : %#v\n", m)

	switch x := m.(type) {
	case comun.MsgLatido:
		//fmt.Println("Recibido latido %#v", x)
		cl.Latido(cl, x.NumVista)

	case comun.MsgVistaTentativa:
		cl.Send(cl.nodotest, comun.MsgVistaTentativa{Vista: x.Vista})

	case comun.MsgPeticionVistaValida:
		//fmt.Println("Recibido peticion vista")
		cl.PeticionVistaValida(cl)

	case comun.MsgVistaValida:
		cl.Send(cl.nodotest, comun.MsgVistaValida{Vista: x.Vista})

	case comun.MsgPeticionPrimario:
		//fmt.Println("Recibido peticion PRIMARIO")
		cl.PeticionPrimario(cl)

	case comun.MsgPrimario:
		cl.Send(cl.nodotest, x)

	case comun.MsgFin:
		log.Println("Recibido FIN")
		cl.stop()
		os.Exit(0) // Aquí termina  la ejecución del servidor

	default:
		fmt.Printf(
			"Recibido mensaje TIPO DESCONOCIDO en servidor_vistas.go!! %#v\n", m)
	}
}
