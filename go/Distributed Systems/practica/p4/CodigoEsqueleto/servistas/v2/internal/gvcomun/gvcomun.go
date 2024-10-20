/*
   paquete "comun" donde reside el código común compartido por
   los diferentes paquetes de este modulo
*/
package gvcomun

import (
	"log"
	"servistas/v2/internal/msgsys"
)

type HostPuerto string // "nombredns:puerto" o "numIP:puerto"

const (
	LATIDOSFALLIDOS  = 4  // nº látidos fallidos para decidir fallo definitivo
	INTERVALOLATIDOS = 50 // en milisegundos
	ANSWERWAITTIME   = 35 // en milisegundos
)

type Vista struct {
	NumVista int
	Primario msgsys.HostPuerto
	Copia    msgsys.HostPuerto
}

// Tipos mensajes concretos para cliente y servidor:
//      Latido, VistaTentativa, VistaValida, DameVistaValida
// Mensajes de cliente
type MsgLatido struct {
	NumVista  int
	Remitente msgsys.HostPuerto
}

type MsgPeticionVistaValida struct {
	Remitente msgsys.HostPuerto
}

type MsgPeticionPrimario struct {
	Remitente msgsys.HostPuerto
}

type MsgFin struct{}

//Mensajes de Servidor
type MsgVistaTentativa struct {
	Vista Vista
}

type MsgVistaValida struct {
	Vista Vista
}

type MsgPrimario msgsys.HostPuerto

// Mensajes entre cliente del servicio de almacenamiento y primario
type MsgWriteTotal struct {
	Mapa      map[string]string
	Remitente msgsys.HostPuerto
}

type MsgWrite struct {
	Clave     string
	Valor     string
	Remitente msgsys.HostPuerto
}

type MsgRead struct {
	Clave     string
	Remitente msgsys.HostPuerto
}

type MsgRespuestaOperacion struct {
	Valor     string
	Remitente msgsys.HostPuerto
}


// Mensajes comunes
type MsgTickInterno struct{}

func RegistrarTiposMensajesGV() {
	// Registrar tipos mensajes de todos los tipos de servidores
	// para decodificación con Encode y Decode de red
	tiposMensaje := []msgsys.Message{MsgLatido{}, MsgVistaTentativa{},
		MsgPeticionPrimario{}, MsgPrimario(""), MsgPeticionVistaValida{},
		MsgVistaValida{}, MsgFin{}}
	msgsys.Registrar(tiposMensaje)
}

func RegistrarMensajesSA() {
	// Registrar tipos mensajes de servidores de almacenamiento
	// para decodificación con Encode y Decode de red
	tiposMensaje := []msgsys.Message{MsgLatido{}, MsgVistaTentativa{},
		MsgWrite{}, MsgRead{}, MsgRespuestaOperacion{}, MsgWriteTotal{}, MsgFin{}}

	msgsys.Registrar(tiposMensaje)
}

func RegistrarMensajesCLTSA() {
	// Registrar tipos mensajes de clientes de almacenamiento
	// para decodificación con Encode y Decode de red
	tiposMensaje := []msgsys.Message{MsgWrite{}, MsgRead{},
		MsgRespuestaOperacion{}, MsgPeticionPrimario{}, MsgPrimario(""),
		MsgPeticionVistaValida{}, MsgVistaValida{}, MsgFin{}}

	msgsys.Registrar(tiposMensaje)
}

// Funciones genericas

func CheckError(err error, comment string) {
	if err != nil {
		ChangeLogPrefix()
		log.Fatalf("Fatal error --- %s -- %s\n", err.Error(), comment)
	}
}

func ChangeLogPrefix() {
	// poner microsegundos y PATH completo de fichero codigo en logs
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}
