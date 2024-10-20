/*
Ejecutable del gestor de vistas con parámetro host:puerto
*/
package main

import (
	"log"
	"os"
	"servistas/v2/internal/gvcomun"
	"servistas/v2/internal/msgsys"
	"servistas/v2/internal/srvvts"
)

func main() {
	// obtener "host:puerto" de este gestor de vistas desde argumentos
	me := os.Args[1]

	gvcomun.ChangeLogPrefix()
	log.Println("Arrancando servidor : ", me)

	// Poner en marcha gestor de vistas
	sv := srvvts.Make(msgsys.HostPuerto(me))
	sv.Start()
}
