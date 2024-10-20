package main

import (
	ms "node/messagesystem"
	"fmt"
)

func main() {
	ms1 := ms.New(1, "users.txt")
	for i := 1; i < 1000; i++ {
		ms1.Send(2, ms.Message{Operation: "Hola", Clock: i})
	}
	ms1.Stop()
	m := ms1.Receive()
	fmt.Println(m.Operation)
}
