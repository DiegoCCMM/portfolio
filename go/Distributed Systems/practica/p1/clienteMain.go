package main

import (
	"fmt"
	"p1Ordenado/client"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var num, yorN int
	var IP, IPserv string
	fmt.Print("Enter client number: ")
	fmt.Scanf("%d", &num)

	fmt.Print("Enter local IP addres: ")
	fmt.Scanf("%s", &IP)

	fmt.Print("Enter server IP addres: ")
	fmt.Scanf("%s", &IPserv)

	fmt.Print("masterWorker?: ")
	fmt.Scanf("%d", &yorN)


	for i:=1; i<num;i++ {

		s := strconv.Itoa(22220 + i)
		go client.ClienteLocal(IPserv+":22223", IP+":"+s, yorN, &wg)
		wg.Add(1)
   }
	wg.Add(1)
	client.ClienteLocal(IPserv+":22223", IP+":22220", yorN, &wg)

	wg.Wait()

}
