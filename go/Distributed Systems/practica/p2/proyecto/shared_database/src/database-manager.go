package main

import (
	"fmt"
	ms "messagesystem"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	msg := ms.New(1, "users.txt")
	f, err := os.Create("operations.log")
	checkError(err)
	exit := false
	for !exit {
		m := msg.Receive()
		if m.Operation == ms.START_WRITE {
			_, err := f.WriteString("Operation from " + strconv.Itoa(m.Sender) + ": " + ms.START_WRITE + "\n")
			checkError(err)
		} else if m.Operation == ms.END_WRITE {
			_, err := f.WriteString("Operation from " + strconv.Itoa(m.Sender) + ": " + ms.END_WRITE + "\n")
			checkError(err)
		} else if m.Operation == ms.START_READ {
			_, err := f.WriteString("Operation from " + strconv.Itoa(m.Sender) + ": " + ms.START_READ + "\n")
			checkError(err)
		} else if m.Operation == ms.END_READ {
			_, err := f.WriteString("Operation from " + strconv.Itoa(m.Sender) + ": " + ms.END_READ + "\n")
			checkError(err)
		} else if m.Operation == ms.READ {
			_, err := f.WriteString("Operation from " + strconv.Itoa(m.Sender) + ": " + ms.READ + "\n")
			checkError(err)
		} else if m.Operation == ms.WRITE {
			_, err := f.WriteString("Operation from " + strconv.Itoa(m.Sender) + ": " + ms.WRITE + "\n")
			checkError(err)
		} else if m.Operation == ms.CLOSE {
			f.Close()
			exit = true
		} else {
			_, err := f.WriteString("Unknown Operation\n")
			checkError(err)
		}
	}
	msg.Stop()
}
