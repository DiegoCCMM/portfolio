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

func workerP(wchannel chan int, encCh chan *gob.Encoder,connCh chan *net.TCPConn) {
	for {
		select {
		case nPrimos := <-wchannel:
			enc:= <- encCh
			conn:= <- connCh
			res := prime.FindPrimes(nPrimos)
			enc.Encode(res)

			conn.Close()
		}
	}
}

func crearPoolWorkers(wchannel chan int, enc chan *gob.Encoder, connCh chan *net.TCPConn, n int) { //TODO RESPUESTA ES CONN PARA QUE ESCRIBAN, CREAR ESTE TUNEL EN EL MAIN
	for i := n; i > 0; i-- {
		go workerP(wchannel, enc, connCh)
	}
}

func servidorConcurrentePool(wchan chan int, encCh chan *gob.Encoder, connCh chan *net.TCPConn) {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	client.CheckError(err)
	fmt.Println("escucho")

	for {

		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		dec := gob.NewDecoder(conn) // Will read from network.

		var nPrimos int
		var clientaddress string
		err = dec.Decode(&nPrimos)
		client.CheckError(err)
		err = dec.Decode(&clientaddress)
		client.CheckError(err)

		wchan <- nPrimos

		enc, conn2 := client.Encoder(clientaddress)

		encCh <- enc
		connCh <- conn2

		conn.Close()
	}
}

func main() {
	connCh := make(chan  *net.TCPConn)
	encCh := make(chan *gob.Encoder)
	wchannel := make(chan int)

	var n int
	fmt.Print("num Workers: ")
	fmt.Scanf("%d", &n)

	crearPoolWorkers(wchannel, encCh, connCh, n)
	fmt.Println("he creado el pool")
	 servidorConcurrentePool(wchannel, encCh, connCh)

}
