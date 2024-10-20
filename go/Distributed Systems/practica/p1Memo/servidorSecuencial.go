package main

import (
	"fmt"
	"./prime"
)

func servidorSecuencial(ch chan int, respuesta chan []int) {
	for{
		select{
			case nPrimos:= <-ch:
				primes:=FindPrimes(nPrimos)
				respuesta<- primes
		}
	}
}

func clienteLocal(ch chan int, respuesta chan []int, load int){
	for  i := load; i > 0; i-- {
		ch<- 60000
	}
	for  i := load; i > 0; i-- {
		primes:= <-respuesta
		fmt.Println(primes[i])
	}
}

func main() {
	ch := make(chan int, 2)
	respuesta:= make(chan []int)

	go servidorSecuencial(ch, respuesta)
	go clienteLocal(ch, respuesta, 1)


}
