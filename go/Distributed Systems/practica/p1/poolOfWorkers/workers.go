package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"p1Ordenado/client"
	"p1Ordenado/prime"
	"strconv"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

const (
	CONN_HOST = "0.0.0.0"
	CONN_TYPE = "tcp"
)

func workerP(local string,  masterAddr string, port string) {

	encServer, connenc := client.Encoder(masterAddr+":22223")

	err := encServer.Encode("newWorker")
	CheckError(err)
	err = encServer.Encode(local)
	CheckError(err)
	connenc.Close()
	fmt.Println("escucho server")
	listener, err := net.Listen(CONN_TYPE, "0.0.0.0"+port)
	client.CheckError(err)

	defer connenc.Close()


	for {
		//acepta conexión
		conndec, err := listener.Accept()
		if err != nil {
			continue
		}

		decServer := gob.NewDecoder(conndec) // Will read from network.

		//nPrimos := <-ch:
		var nPrimosCliente int
		var clientAddr string
		err = decServer.Decode(&nPrimosCliente)
		CheckError(err)
		err = decServer.Decode(&clientAddr)
		CheckError(err)

		primes := prime.FindPrimes(nPrimosCliente)

		encCl, connCl := client.Encoder(clientAddr)

		err = encCl.Encode(primes)
		CheckError(err)

		//signal
		encServer, connenc := client.Encoder(masterAddr+":22223")
		err = encServer.Encode("end")
		CheckError(err)
		err = encServer.Encode(local)
		CheckError(err)

		connCl.Close()
		conndec.Close()
		connenc.Close()

	}
}

func crearPoolWorkers(local string, masterAddr string, n int) {
	for i := n; i > 1; i-- {
		go workerP(local+":2323"+strconv.Itoa(i-1), masterAddr, ":2323"+strconv.Itoa(i-1))
	}
	workerP(local+":23230", masterAddr, ":23230")
}

func main() {

	var local string
	var IP string
	var n int


	fmt.Print("Enter master IP addres: ")
	fmt.Scanf("%s", &IP)

	fmt.Print("Enter local IP addres: ")
	fmt.Scanf("%s", &local)

	fmt.Print("Enter num workers: ")
	fmt.Scanf("%d", &n)

	crearPoolWorkers(local, IP,  n)

}
