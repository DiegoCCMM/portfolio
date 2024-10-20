package main

import (
	"fmt"
	ms "node/messagesystem"
)

func main() {
	ms2 := ms.New(2, "users.txt")
	for i := 1; i < 1000; i++ {
		m := ms2.Receive()
		fmt.Println("Mensaje: ", i, " ", m.Operation)
	}
	ms2.Send(1, ms.Message{Operation: "REQ", Clock: 2})
}

