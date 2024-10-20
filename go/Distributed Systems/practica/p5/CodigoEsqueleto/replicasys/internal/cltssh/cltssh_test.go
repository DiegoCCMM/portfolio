package cltssh

import (
	"fmt"
	"replicasys/internal/comun"
	"replicasys/internal/msgsys"
	"testing"
	"time"
)

func TestBasico(t *testing.T) {
	cmd :=
		"cd /home/tmp/servistas/v2/cmd/; go run cmdsrvvts/main.go 127.0.0.1:29009"
	r := make(chan string, 1000)
	ExecMutipleHosts(cmd, []string{"127.0.0.1"}, r, "/home/unai/.ssh/id_ed25519")

	time.Sleep(1000 * time.Millisecond)

	msgsys.Registrar([]msgsys.Message{comun.MsgFin{}})
	//buzonReadTests, doneReceivers :=
	ms := msgsys.MakeMsgSys("127.0.0.1:29008")
	ms.Send("127.0.0.1:29009", comun.MsgFin{})

	// esperar parada se servidor remoto el tiempo suficiente
	// para volcar salida de ejecuciones ssh en cmdOutput
	time.Sleep(500 * time.Millisecond)
	ms.CloseMessageSystem()
	close(r) //para que termine el bucle for siguiente, en lugar de bloquear

	for s := range r {
		fmt.Println(s)
	}

}
