package main

import (
	"fmt"
	"os"
	ms "p2GoMod/messagesystem"
	"strconv"
	ra "p2GoMod/ricartAgrawala"
	"sync"
    "github.com/DistributedClocks/GoVector/govec"
)

//varibles globales
var nProcess, myPid, clockglobal, lrd, state *int
var delay *[]int


func request(ms  *ms.MessageSystem, permission *int, wg *sync.WaitGroup, Log *govec.GoLog){

	for {
		*clockglobal = ra.Request(ms, *myPid, ra.READER, *lrd, *clockglobal, permission, state, delay, Log)
		//fmt.Println("clock global tras request ", *clockglobal)
//		wg.Done()
	}

}

func main() {
	var snProcess, smyPid, sclockglobal, sldr, sstate, permission int
	var sdelay []int
	var wg sync.WaitGroup

	//inicializamos el reloj a 0
	sldr = 0
	sclockglobal = 0

	//variables globales
	 nProcess = &snProcess
	 myPid = &smyPid
	 clockglobal = &sclockglobal
	 lrd = &sldr
	 state = &sstate
	 delay = &sdelay
	//overhead del paso de parametros
	argsWithoutProg := os.Args[1:]  //número de procesos a los que esperar + quien soy

	*nProcess, _ = strconv.Atoi(argsWithoutProg[0])
	*myPid, _ = strconv.Atoi(argsWithoutProg[1])

	Logger := govec.InitGoVector("MyProcess "+strconv.Itoa(*myPid), "LogFile", govec.GetDefaultConfig())

	msg := ms.New(*myPid, "users.txt")	//ms

	//sincronización
	if(*myPid != 3) {
		msg.Receive()
	}else{
		for i := 2; i<=*nProcess +1; i++ {
			if(i != *myPid) {
				msg.Send(i, ms.Message{Operation: "RequestWriter", Sender: *myPid})
			}
		}
	}


	//PREPROTOCOL

	//for i:=1 ; i !=3; i++ {
	//prueba clock

		permission = *nProcess - 1 //waiting
		wg.Add(permission)
		*state = 1 //trying
		fmt.Println("intento entrar")
		*lrd = *clockglobal + 1                    //mi reloj local
		go request(&msg, &permission, &wg, Logger) //proceso request

		for i := 2; i <= *nProcess+1; i++ { //mensaje a todos
			if (i != *myPid) {

				//msg.Send(i, ms.Message{Operation: "RequestReader " + strconv.Itoa(*lrd), Sender: *myPid})
				vectorclockmessage := Logger.PrepareSend("Sending Message", "RequestReader "+strconv.Itoa(*lrd), govec.GetDefaultLogOptions())
				msg.Send(i, ms.Message{Operation: string(vectorclockmessage), Sender: *myPid})
				fmt.Println(" he enviado request", i)

			}
		}

		for permission != 0 { //espero
		}

		fmt.Println("entro en SC")
		*state = 2 //in
		//SC
		msg.Send(1, ms.Message{Operation: ms.START_READ, Sender: *myPid}) //luego la creación de un mensaje que es una operation y un sender (mi pid global)
		msg.Send(1, ms.Message{Operation: ms.READ, Sender: *myPid})
		msg.Send(1, ms.Message{Operation: ms.READ, Sender: *myPid})
		msg.Send(1, ms.Message{Operation: ms.END_READ, Sender: *myPid})
		//msg.Send(1,ms.Message{Operation: ms.CLOSE, Sender: 2}) //para terminar la ejecución

		//POST PROTOCOL
		fmt.Println("acabé")
		*state = 3
		for i := len(sdelay)-1; i  > -1; i-- {

				vectorclockmessage := Logger.PrepareSend("Sending Message", "Permission "+strconv.Itoa(*lrd), govec.GetDefaultLogOptions())
				msg.Send(sdelay[i], ms.Message{Operation: string(vectorclockmessage), Sender: *myPid})
			sdelay = sdelay[:len(sdelay)-1]

		}
	//}
	//prueba clock

	for{
		wg.Add(1)
		wg.Wait()
	}
}
