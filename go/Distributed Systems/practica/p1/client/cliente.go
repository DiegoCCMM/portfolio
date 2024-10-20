package client

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "22222"
	CONN_TYPE = "tcp"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func Encoder(address string) (*gob.Encoder, *net.TCPConn) { //addres := addres:port
	//crear conexión a la dirección
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	CheckError(err)
	fmt.Println("resuelvo")
	//connect
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)
	fmt.Println("conecto")
	//utilizamos gob para generalizar el uso de los bytes
	enc := gob.NewEncoder(conn) // Will write to network.
	fmt.Println("encoderEND")

	return enc, conn
}

func decoder(address string) *gob.Decoder {	//addres := addres:port
	listener, err := net.Listen("tcp", address)
	CheckError(err)
	conn, err := listener.Accept()

	dec := gob.NewDecoder(conn) // Will read from network.

	return dec
}


func ClienteLocal(address string, local string, mW int, wg *sync.WaitGroup) {

	var start time.Time

	enc,conn := Encoder(address)
	defer conn.Close()

	if(mW == 1)  {
		err := enc.Encode("peticion Cliente")
		CheckError(err)
	}

	start = time.Now() //empezamos a contar en el envío del mensaje
	//ch <- 60000
	nPrimosCliente := 60000
	fmt.Println("envio")
	err := enc.Encode(nPrimosCliente)
	CheckError(err)
	err = enc.Encode(local)
	CheckError(err)
	fmt.Println("envioEND")

	dec := decoder(local)

	//<-respuesta
	var primes []int
	err = dec.Decode(&primes)
	CheckError(err)

	final := time.Now()
	elapsed := final.Sub(start)
	fmt.Println("recibo y fin cliente Tcliente: ", elapsed)
	wg.Done()
}
