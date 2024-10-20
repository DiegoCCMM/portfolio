package ricartAgrawala

import (
	"fmt"
	"github.com/DistributedClocks/GoVector/govec"
	"math"
	msl "p2GoMod/messagesystem"
	"strconv"
	"strings"
)



const(
	READER = 1
	WRITER = 0
)

func exclude(op1 int, op2 int) (bool){
	if((op1 + op2 ) == 2) {	//ambos dos readers, no se excluye
		return false;
	}else{									//cualquier otro caso, son o reader writer o writer writer y se excluyen entre ellos
		return true;
	}
}

func Send(ms *msl.MessageSystem, sended int, sender int, operation int, clock int){

	if(operation == READER){
		ms.Send(sended, msl.Message{Operation: "RequestReader", Sender: sender})

	}else{	//Writer
		ms.Send(sended, msl.Message{Operation: "RequestWriter", Sender: sender})
	}
	sclock := strconv.Itoa( clock)
	ms.Send(1,msl.Message{Operation: sclock,Sender: sender})

}

func Request(ms *msl.MessageSystem, sender int, whatami int, myclock int, globalclock int ,
	permission *int, state *int, delay *[]int, Logger *govec.GoLog)(clocksal int){

	var received msl.Message

	received =ms.Receive()
	fmt.Println("recibo mensaje")

	var messagepayload string
	Logger.UnpackReceive("Receiving Message", []byte(received.Operation), &messagepayload, govec.GetDefaultLogOptions())

	vectstring := strings.Split(messagepayload, " ")
	fmt.Println(messagepayload)
	sclock:=vectstring[1]

	iclock, _ := strconv.Atoi(sclock)
	
	clocksal = int(math.Max(float64(globalclock), float64(iclock)))

	var operationtype int
	var  j int
	var typeMessage string

	j=received.Sender

	fmt.Println("PID sender",received.Sender)
	fmt.Println("clock sender", iclock)
	fmt.Println("my clock", myclock)

	if(vectstring[0] == "RequestReader"){
		operationtype = READER
		typeMessage = "Request"
	}else if(vectstring[0] == "RequestWriter"){
		operationtype = WRITER
		typeMessage = "Request"
	}else if(vectstring[0] == "Permission"){
		typeMessage = "Permission"
		fmt.Print(j)
		fmt.Println( " me ha dicho que si")
		*permission = *permission - 1;
	}

	var tuples bool

	if(myclock == iclock){ tuples = sender <j } else { tuples = myclock < iclock}

	if(typeMessage == "Request") {
		if(*state != 3 && exclude(operationtype,whatami) && tuples ){
			*delay = append(*delay, j)
			fmt.Print(j)
			fmt.Println( " le deniego")
		}else{
			vectorclockmessage := Logger.PrepareSend("Sending Message", "Permission "+ strconv.Itoa(myclock), govec.GetDefaultLogOptions())
			ms.Send(j, msl.Message{Operation: string(vectorclockmessage), Sender: sender})
			//ms.Send(j,msl.Message{Operation: strconv.Itoa(myclock), Sender: sender} )
			fmt.Print(sender)
			fmt.Print( " le doy permiso a ")
			fmt.Println(j)
		}
	}
	return clocksal
}
