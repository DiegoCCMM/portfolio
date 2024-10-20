/*
Ejecutable de replica/servidor de almacenamiento
*/
package main

import (
	"os"
	"servistas/v2/internal/msgsys"
	"servistas/v2/internal/srvalm"
)

func main() {
	// obtener host:puerto de esta réplica desde argumentos
	me := os.Args[1]

	// segundo parámetro: host:puerto del gestor de vistas
	gv := os.Args[2]

	// Poner en marcha replica de almacenamiento
	sa := srvalm.Make(msgsys.HostPuerto(me), msgsys.HostPuerto(gv))
	sa.Start()
}
