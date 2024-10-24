/*
Paquete para test de integración únicamente del servidor de gestión de vistas
*/
package integraciongvtest

import (
	"fmt"
	"replicasys/internal/cltssh"
	"replicasys/internal/comun"
	"replicasys/internal/msgsys"
	"testing"
	"time"
)

const (
	//hosts
	MAQUINA1 = "127.0.0.1"
	MAQUINA2 = "127.0.0.1"
	MAQUINA3 = "127.0.0.1"
	MAQUINA4 = "127.0.0.1"

	//puertos
	PUERTO0 = "29000"
	PUERTO1 = "29001"
	PUERTO2 = "29002"
	PUERTO3 = "29003"
	PUERTO4 = "29004"

	//nodos
	NODOTEST     = MAQUINA1 + ":" + PUERTO0
	NODOGV       = MAQUINA2 + ":" + PUERTO1
	NODOCLIENTE1 = MAQUINA3 + ":" + PUERTO2
	NODOCLIENTE2 = MAQUINA4 + ":" + PUERTO3
	NODOCLIENTE3 = MAQUINA1 + ":" + PUERTO4

	// PATH de los ejecutables de modulo golang de servicio de vistas
	PATH = "/home/tmp/replicasys/cmd/"

	// fichero main de ejecutables relativos a PATH previo
	EXECGV      = "cmdsrvvts/main.go " + NODOGV // un parámetro
	EXECTESTCLT = "testcltvts/main.go "         // 2 parámetros en llamadas ssh

	// comandos completo a ejecutar en máquinas remota con ssh
	// cd /home/tmp/servistas/v2/cmd/; go run cmdsrvvts/main.go 127.0.0.1:29001
	SRVVTSCMD = "cd " + PATH + "; go run " + EXECGV
	// cd /home/tmp/servistas/v2/cmd/;
	// go run testcltvts/main.go 127.0.0.1:29003 127.0.0.1:29001 127.0.0.1:29000
	CLTVTSCMD = "cd " + PATH + "; go run " + EXECTESTCLT

	// Ubicar, en esta constante, el PATH completo a vuestra clave privada local
	// emparejada con la clave pública en authorized_keys de máquinas remotas

	PRIVKEYFILE = "/home/unai/.ssh/id_ed25519"
)

// TEST primer rango
func TestPrimerasPruebas(t *testing.T) { // (m *testing.M) {
	// <setup code>
	// Crear servidor de test y
	//procesos en maquinas remotas : servidor de vistas
	ts := startTestServer(NODOTEST)
	ts.startDistributedProcesses()

	// Run test sequence

	// Test1 : No debería haber ningun primario, si SV no ha recibido aún latidos
	t.Run("P=0:T1",
		func(t *testing.T) { ts.ningunPrimarioTest1(t) })

	// Test2: tenemos el primer primario correcto
	t.Run("P=1:T2",
		func(t *testing.T) { ts.primerPrimarioTest2(t) })

	// Test3: Primer nodo copia
	t.Run("C=1:T3",
		func(t *testing.T) { ts.PrimerCopiaTest3(t) })

	// tear down code
	// eliminar procesos en máquinas remotas
	ts.stopDistributedProcesses()
	ts.stop()
}

// ---------------------------------------------------------------------
// Servidor de test

type testServer struct {
	msgsys.MsgSys
	// Canal de resultados de ejecución de comadnos ssh remotos
	cmdOutput chan string
}

func startTestServer(me msgsys.HostPuerto) (ts testServer) {
	// Registrar tipos de mensaje de gestión d vistas
	comun.RegistrarMensajesGV()

	ts = testServer{
		MsgSys:    msgsys.MakeMsgSys(me),
		cmdOutput: make(chan string, 1000),
	}

	return ts
}

func (ts *testServer) stop() {
	ts.CloseMessageSystem()
	close(ts.cmdOutput)

	// Leer las salidas obtenidos de los ssh ejecutados por servidor de tests
	for s := range ts.cmdOutput {
		fmt.Println(s)
	}
}

func (ts *testServer) startDistributedProcesses() {

	// Poner en marcha servidor/gestor de vistas y 3 clientes
	// en 4 máquinas remota con ssh
	cltssh.ExecMutipleHosts(SRVVTSCMD,
		[]string{MAQUINA2}, ts.cmdOutput, PRIVKEYFILE)
	cltssh.ExecMutipleHosts(CLTVTSCMD+NODOCLIENTE1+" "+NODOGV+" "+NODOTEST,
		[]string{MAQUINA3}, ts.cmdOutput, PRIVKEYFILE)
	cltssh.ExecMutipleHosts(CLTVTSCMD+NODOCLIENTE2+" "+NODOGV+" "+NODOTEST,
		[]string{MAQUINA4}, ts.cmdOutput, PRIVKEYFILE)
	cltssh.ExecMutipleHosts(CLTVTSCMD+NODOCLIENTE3+" "+NODOGV+" "+NODOTEST,
		[]string{MAQUINA1}, ts.cmdOutput, PRIVKEYFILE)

	// ajustar si necesario para esperar al
	// tiempo de establecimiento de sesión de ssh
	time.Sleep(4000 * time.Millisecond)
}

func (ts *testServer) stopDistributedProcesses() {

	// Parar procesos distribuidos con ssh
	// una opción :
	ts.Send(NODOGV, comun.MsgFin{})
	ts.Send(NODOCLIENTE1, comun.MsgFin{})
	ts.Send(NODOCLIENTE2, comun.MsgFin{})
	ts.Send(NODOCLIENTE3, comun.MsgFin{})

	// esperar parada se servidores remotos el tiempo suficiente
	// para volcar salida de ejecuciones ssh en cmdOutput
	time.Sleep(100 * time.Millisecond)
}

// --------------------------------------------------------------------------
// FUNCIONES DE SUBTESTS

// No debería haber primario
func (ts *testServer) ningunPrimarioTest1(t *testing.T) {
	fmt.Println(t.Name(), ".....................")

	// obten la respuesta a  la petición de primario
	p, ok := ts.SendReceive(NODOCLIENTE1,
		comun.MsgPeticionPrimario{Remitente: NODOCLIENTE1},
		comun.ANSWERWAITTIME*time.Millisecond)

	if !ok {
		fmt.Printf("TIMEOUT SENDRECEIVE NINGUN PRIMERO TEST: %s", NODOCLIENTE1)
		t.Fatalf(
			"Ha saltado timeout esperando respuesta de Gestor de Vistas %#v",
			t.Name())
	}

	if p != comun.MsgPrimario(msgsys.HOSTINDEFINIDO) {
		t.Fatalf("Primario = %#v; DESEABLE DESCONOCIDO = %#v",
			p, msgsys.HOSTINDEFINIDO)
	}

	fmt.Println(".............", t.Name(), "Superado")
}

// No debería haber primario
func (ts *testServer) primerPrimarioTest2(t *testing.T) {
	// t.Skip("SKIPPED primerPrimarioTest2")

	fmt.Println(t.Name(), ".....................")

	// Primer cliente por primera vez :
	// 		latido 0 y vista tentativa por respuesta en tiempo razonable
	vTentativa := ts.clienteLatido0(t, NODOCLIENTE1)

	// Preparar las vistas a comparar entre recibida y vista esperada
	vac := vistasAcomparar{t: t,
		recibido: vTentativa,
		referencia: comun.Vista{Primario: NODOCLIENTE1,
			Copia:    msgsys.HOSTINDEFINIDO,
			NumVista: 1},
	}

	// Comprobar vista tentativa recibida
	vac.comprobar()

	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) PrimerCopiaTest3(t *testing.T) {
	//t.Skip("SKIPPED PrimerCopiaTest3")

	// solo nos interesa la vista tentativa devuelta por latido a Gestor Vistas
	_, ok := ts.SendReceive(NODOCLIENTE1,
		comun.MsgLatido{NumVista: -1, Remitente: NODOCLIENTE1},
		comun.ANSWERWAITTIME*time.Millisecond,
	)
	if !ok {
		t.Fatal("Salta timeout esperando respuesta a latido -1 de cliente")
	}

	// Segundo cliente por primera vez:
	// 		latido 0 y vista tentativa por respuesta en tiempo razonable
	vTentativa := ts.clienteLatido0(t, NODOCLIENTE2)

	// Preparar las vistas a comparar entre recibida y vista esperada
	vac := vistasAcomparar{t: t,
		recibido: vTentativa,
		referencia: comun.Vista{Primario: NODOCLIENTE1,
			Copia:    NODOCLIENTE2,
			NumVista: 2},
	}

	// Comprobar vista tentativa recibida
	vac.comprobar()

	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) clienteLatido0(t *testing.T,
	nodoCliente msgsys.HostPuerto) comun.Vista {

	// solo nos interesa la vista tentativa devuelta por latido a Gestor Vistas
	m, ok := ts.SendReceive(nodoCliente,
		comun.MsgLatido{NumVista: 0, Remitente: nodoCliente},
		comun.ANSWERWAITTIME*time.Millisecond,
	)
	if !ok {
		t.Fatal("Salta timeout esperando latido 0 de cliente")
	}

	switch x := m.(type) {
	case comun.MsgVistaTentativa:
		return x.Vista // salida correcta
	default:
		t.Fatalf(t.Name(),
			"Mensaje recibido INCORRECTO en primerPrimario: %#v", x)
	}

	// no debería llegar a ejecutarse, pero se pone por error compilacion
	return comun.Vista{}
}

// --------------------------------------------------------------------------
// FUNCIONES DE APOYO

type vistasAcomparar struct {
	recibido   comun.Vista
	referencia comun.Vista
	t          *testing.T
}

func (vs vistasAcomparar) comprobar() {
	if vs.recibido.Primario != vs.referencia.Primario {
		vs.t.Fatalf(
			"%s : PRIMARIO recibido (%s) y de referencia (%s) no coinciden",
			vs.t.Name(), vs.recibido.Primario, vs.referencia.Primario)
	}
	if vs.recibido.Copia != vs.referencia.Copia {
		vs.t.Fatalf("%s : COPIA recibido (%s) y de referencia (%s) no coinciden",
			vs.t.Name(), vs.recibido.Copia, vs.referencia.Copia)
	}
	if vs.recibido.NumVista != vs.referencia.NumVista {
		vs.t.Fatalf(
			"%s : NUM VISTA recibido (%d) y de referencia (%d) no coinciden",
			vs.t.Name(), vs.recibido.NumVista, vs.referencia.NumVista)
	}
}
