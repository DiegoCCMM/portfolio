package main

import (
	"fmt"
	"net"
	"os"
)

func client(x int, y int)  {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()
	
	buf := bytes.Buffer{}
	_, err = conn.Read(buf)
	checkError(err)

	fmt.Println("recibidos: " + buf.String())

	os.Exit(0)

}

