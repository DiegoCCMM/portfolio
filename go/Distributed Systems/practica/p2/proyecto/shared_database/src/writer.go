package main

import (
	ms "messagesystem"
)

func main() {
	msg := ms.New(3, "users.txt")
	msg.Send(1, ms.Message{Operation: ms.START_WRITE, Sender: 3})
	msg.Send(1, ms.Message{Operation: ms.WRITE, Sender: 3})
	msg.Send(1, ms.Message{Operation: ms.WRITE, Sender: 3})
	msg.Send(1, ms.Message{Operation: ms.WRITE, Sender: 3})
	msg.Send(1, ms.Message{Operation: ms.WRITE, Sender: 3})
	msg.Send(1, ms.Message{Operation: ms.WRITE, Sender: 3})
	msg.Send(1, ms.Message{Operation: ms.WRITE, Sender: 3})
	msg.Send(1, ms.Message{Operation: ms.WRITE, Sender: 3})
	msg.Send(1, ms.Message{Operation: ms.END_WRITE, Sender: 3})
}
