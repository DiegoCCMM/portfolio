package messagesystem

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Message struct {
	Operation string
	Clock     int
}

type MessageSystem struct {
	mbox  chan Message
	peers []string
	done  chan bool
	me    int
}

const (
	REQUEST     = "REQUEST"
	REPLY       = "REPLY"
	MAXMESSAGES = 10000
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func parsePeers(path string) (lines []string) {
	file, err := os.Open(path)
	checkError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func (ms MessageSystem) Send(pid int, msg Message) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ms.peers[pid-1])
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(msg)
	conn.Close()
}

func (ms *MessageSystem) Receive() (msg Message) {
	msg = <-ms.mbox
	return msg
}

func New(whoIam int, usersFile string) (ms MessageSystem) {
	ms.me = whoIam
	ms.peers = parsePeers(usersFile)
	ms.mbox = make(chan Message, MAXMESSAGES)
	ms.done = make(chan bool)
	go func() {
		listener, err := net.Listen("tcp", ms.peers[ms.me-1])
		checkError(err)
		fmt.Println("Process listening at " + ms.peers[ms.me-1])
		defer close(ms.mbox)
		for {
			select {
			case <-ms.done:
				return
			default:
				conn, err := listener.Accept()
				checkError(err)
				decoder := gob.NewDecoder(conn)
				var msg Message
				err = decoder.Decode(&msg)
				conn.Close()
				ms.mbox <- msg
			}
		}
	}()
	return ms
}

func (ms *MessageSystem) Stop() {
	ms.done <- true
}
