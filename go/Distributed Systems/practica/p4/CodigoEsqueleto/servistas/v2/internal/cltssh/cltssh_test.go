package cltssh

import (
	"fmt"
	"servistas/v2/internal/gvcomun"
	"servistas/v2/internal/msgsys"
	"testing"
	"time"
)

var host string

func TestBasico(t *testing.T) {
	host = "155.210.154.196:23232"
	cmd :=
		"cd /home/a738712/SSDD/p4/v2/cmd ; /usr/local/go/bin/go run cmdsrvvts/main.go " + host
	r := make(chan string, 1000)
	ExecMutipleHosts(cmd, []string{"155.210.154.196"}, r, "/home/a738712/.ssh/id_rsa")	//local es /home/diego/.ssh/id_rsa

	time.Sleep(5000 * time.Millisecond)

	msgsys.Registrar([]msgsys.Message{gvcomun.MsgFin{}})
	//buzonReadTests, doneReceivers :=
	ms := msgsys.MakeMsgSys("0.0.0.0:22444")
	ms.Send(msgsys.HostPuerto(host), gvcomun.MsgFin{})

	// esperar parada se servidor remoto el tiempo suficiente
	// para volcar salida de ejecuciones ssh en cmdOutput
	time.Sleep(100 * time.Millisecond)
	ms.CloseMessageSystem()
	close(r) //para que termine el bucle for siguiente, en lugar de bloquear

	for s := range r {
		fmt.Println(s)
	}

}
