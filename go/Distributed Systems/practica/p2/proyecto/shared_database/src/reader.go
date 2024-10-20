package main

import (
	ms "messagesystem"
)

func main() {
	msg := ms.New(2, "users.txt")
	msg.Send(1, ms.Message{Operation: ms.START_READ, Sender: 2})
	msg.Send(1, ms.Message{Operation: ms.READ, Sender: 2})
	msg.Send(1, ms.Message{Operation: ms.READ, Sender: 2})
	msg.Send(1, ms.Message{Operation: ms.END_READ, Sender: 2})
}
