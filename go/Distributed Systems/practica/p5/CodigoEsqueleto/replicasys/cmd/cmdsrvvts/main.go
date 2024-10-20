/*
Ejecutable del gestor de vistas
*/
package main

import (
	"os"
	"replicasys/internal/msgsys"
	"replicasys/internal/srvvts"
)

func main() {
	// obtener host:puerto de este gestor de vistas desde argumentos
	me := os.Args[1]

	// Poner en marcha gestor de vistas
	sv := srvvts.Make(msgsys.HostPuerto(me))
	sv.Start()
}
