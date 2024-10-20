/*AUTORES: nombres y apellidos
NIAs: números de identificaci'on de los alumnos
FICHERO: nombre del fichero
FECHA: fecha de realizaci’on
TIEMPO: tiempo en horas de codificación
DESCRIPCION: breve descripcion del contenido del fichero*/

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



func servidorSecuencial() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	client.CheckError(err)
	fmt.Println("escucho")

	for {

		conn, err := listener.Accept()
		fmt.Println("acepto")
		if err != nil {
			continue
		}
		dec := gob.NewDecoder(conn) // Will read from network.

		fmt.Println("decoderEND")

		//nPrimos := <-ch:
		var nPrimos int
		var clientaddress string
		err = dec.Decode(&nPrimos)
		client.CheckError(err)
		err = dec.Decode(&clientaddress)
		client.CheckError(err)


		fmt.Println("desbloqueo")

		primes := prime.FindPrimes(nPrimos)

		enc, conn2 := client.Encoder(clientaddress)

		enc.Encode(primes)

		//respuesta <- primes
		client.CheckError(err)
		conn2.Close()
		fmt.Println("respondo")

		conn.Close()
	}
}

func main() {

	servidorSecuencial()

}
