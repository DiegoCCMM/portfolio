package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"p1Ordenado/client"
)

import "container/list"

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "22223"
	CONN_TYPE = "tcp"
)



func master(){

	workers := list.New()
	peticionesCliente := list.New()


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

		var mensaje string
		err = dec.Decode(&mensaje)
		client.CheckError(err)

		if(mensaje == "peticion Cliente") {	//petición cliente
			fmt.Println("peticion Cliente")
			//nPrimos := <-ch:
			var nPrimos int
			var clientaddress string
			err = dec.Decode(&nPrimos)
			client.CheckError(err)
			err = dec.Decode(&clientaddress)
			client.CheckError(err)

			if(workers.Len()>0) {
				//si hay workers, les mando la petición y los borro de la lista
				addrW := fmt.Sprintf("%v", workers.Back().Value)
				encW, connW := client.Encoder(addrW)

				err = encW.Encode(nPrimos)
				client.CheckError(err)
				err = encW.Encode(clientaddress)
				client.CheckError(err)

				workers.Remove(workers.Back())
				connW.Close()
			}else{//si no hay workers, no les mando la petición
				peticionesCliente.PushFront(clientaddress)

			}

		}else if(mensaje == "newWorker"){	//register
			//apunto al worker
			fmt.Println("newWorker")
			fmt.Println(workers.Len())
			var workerIP string
			err = dec.Decode(&workerIP)
			client.CheckError(err)

			workers.PushFront(workerIP)

		}else if(mensaje == "end"){ //signal worker
			fmt.Println("end")
			//reapunto al worker
			var workerIP string
			err = dec.Decode(&workerIP)
			client.CheckError(err)

			//vaciar peticiones encoladas
			if(peticionesCliente.Len() > 0){
				addr := fmt.Sprintf("%v", peticionesCliente.Front().Value)

				encW, connW := client.Encoder(workerIP)

				err = encW.Encode(60000)
				client.CheckError(err)
				err = encW.Encode(addr)
				client.CheckError(err)

				peticionesCliente.Remove(peticionesCliente.Front())
				connW.Close()
			}else{
				workers.PushFront(workerIP)
			}
		}
		conn.Close()
	}

}

func main() {
	master()
}
