package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"p1Ordenado/client"
	"p1Ordenado/prime"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "22223"
	CONN_TYPE = "tcp"
)

func worker(nPrimos int, enc *gob.Encoder, conn net.Conn) {
	primos := prime.FindPrimes(nPrimos)
	//respuesta <- primos
	err := enc.Encode(primos)
	client.CheckError(err)

	fmt.Println("respondo")

}

func servidorConcurrentePet() {
	//crea escucha
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	client.CheckError(err)

	for {
		fmt.Println("escucho")

		//acepta conexión
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		dec := gob.NewDecoder(conn) // Will read from network.

		//nPrimos := <-ch:
		var clientAddr string
		var nPrimos int
		err = dec.Decode(&nPrimos)
		client.CheckError(err)

		err = dec.Decode(&clientAddr)
		client.CheckError(err)

		enc, conn2 := client.Encoder(clientAddr)

		go worker(nPrimos, enc, conn2)

		conn.Close()

	}
}

func main() {

	servidorConcurrentePet()

}
