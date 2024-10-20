package integracion_test

import (
	"fmt"
	"net"
	"servistas/v2/internal/cltssh"
	"servistas/v2/internal/gvcomun"
	"servistas/v2/internal/msgsys"
	"testing"
	"time"
)

const (
	//hosts
	MAQUINA1 = "155.210.154.200"
	MAQUINA2 = "155.210.154.201"
	MAQUINA3 = "155.210.154.202"
	MAQUINA4 = "155.210.154.203"

	//puertos
	PUERTO0 = "23232"
	PUERTO1 = "23232"
	PUERTO2 = "23231"
	PUERTO3 = "23239"
	PUERTO4 = "23235"

	//nodos
	NODOTEST     = MAQUINA1 + ":" + PUERTO0
	NODOGV       = MAQUINA2 + ":" + PUERTO1
	NODOCLIENTE1 = MAQUINA3 + ":" + PUERTO2
	NODOCLIENTE2 = MAQUINA4 + ":" + PUERTO3
	NODOCLIENTE3 = MAQUINA1 + ":" + PUERTO4

	// PATH de los (plural) ejecutables de modulo golang de servicio de vistas
	PATH = "/home/a738712/SSDD/p4/v2/cmd"

	// fichero main de ejecutables relativos a PATH previo
	EXECGV      = "cmdsrvvts/main.go " + NODOGV // un parámetro
	EXECTESTCLT = "testcltvts/main.go "         // 2 parámetros en llamadas ssh

	// comandos completo a ejecutar en máquinas remota con ssh
	// cd /home/tmp/servistas/v2/cmd/; go run cmdsrvvts/main.go 127.0.0.1:29001
	SRVVTSCMD = "cd " + PATH + " ; /usr/local/go/bin/go run " + EXECGV
	// cd /home/tmp/servistas/v2/cmd/;
	// go run testcltvts/main.go 127.0.0.1:29003 127.0.0.1:29001 127.0.0.1:29000
	CLTVTSCMD = "cd " + PATH + " ; /usr/local/go/bin/go run " + EXECTESTCLT

	// Ubicar, en esta constante, el PATH completo a vuestra clave privada local
	// emparejada con la clave pública en authorized_keys de máquinas remotas

	PRIVKEYFILE = "/home/a738712/.ssh/id_rsa"
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

	// Test4: Primario cae, ahora copia es primaro
	t.Run("P=1;C=0:T4",
		func(t *testing.T) { ts.CopiaAPrimarioTest4(t) })
	// Test5: Primario vuelve, ahora es copia
	t.Run("P=1;C=1(alternos):T5",
		func(t *testing.T) { ts.PrimarioACopiaTest5(t) })
	t.Run("P=1;C=3:T6",
		func(t *testing.T) { ts.EsperaACopiaTest6(t) })
	t.Run("P=1;C=3;el primario no está:T7",
		func(t *testing.T) { ts.PrimarioAEsperaTest7(t) })
	t.Run("P=1;C=3;VistaValidaVacia",
		func(t *testing.T) { ts.VistaValidaNoValidadTest8(t) })
	t.Run("GV down",
		func(t *testing.T) { ts.ErrorCritico9(t) })
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
	gvcomun.RegistrarTiposMensajesGV()

	ts = testServer{
		MsgSys:    msgsys.MakeMsgSys(me),
		cmdOutput: make(chan string, 1000),
	}

	return ts
}

func (ts *testServer) stop() {
	ts.CloseMessageSystem()
	close(ts.cmdOutput)

	// Visulaizar salidas obtenidos de los ssh ejecutados por servidor de tests
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
	time.Sleep(10000 * time.Millisecond)
}

func (ts *testServer) stopDistributedProcesses() {

	// Parar procesos distribuidos con ssh
	// una opción :
	//ts.Send(NODOGV, gvcomun.MsgFin{})	TODO SIEMPRE QUE SE CORRA TEST9
	ts.Send(NODOCLIENTE1, gvcomun.MsgFin{})
	ts.Send(NODOCLIENTE2, gvcomun.MsgFin{})
	ts.Send(NODOCLIENTE3, gvcomun.MsgFin{})

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
		gvcomun.MsgPeticionPrimario{Remitente: NODOCLIENTE1},
		gvcomun.ANSWERWAITTIME*time.Millisecond)

	if !ok {
		fmt.Printf("TIMEOUT SENDRECEIVE NINGUN PRIMERO TEST: %s", NODOCLIENTE1)
		t.Fatalf(
			"Ha saltado timeout esperando respuesta de Gestor de Vistas %#v",
			t.Name())
	}

	if p != gvcomun.MsgPrimario(msgsys.HOSTINDEFINIDO) {
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
		referencia: gvcomun.Vista{Primario: NODOCLIENTE1,
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
		gvcomun.MsgLatido{-1, NODOCLIENTE1},
		gvcomun.ANSWERWAITTIME*time.Millisecond,
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
		referencia: gvcomun.Vista{Primario: NODOCLIENTE1,
			Copia:    NODOCLIENTE2,
			NumVista: 2},
	}

	// Comprobar vista tentativa recibida
	vac.comprobar()

	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) CopiaAPrimarioTest4(t *testing.T) {
	//t.Skip("SKIPPED PrimerCopiaTest3")

	//espera 4 (+ 1 por si acaso) latidos
	for i := 0; i < 5; i++ {
		ts.enviaLatidoVivo(t, NODOCLIENTE2)
		time.Sleep(gvcomun.INTERVALOLATIDOS * time.Millisecond)
	}

	//pide vista tentativa
	a := ts.vistaTentativa(t, NODOCLIENTE2)

	vac := vistasAcomparar{t: t,
		recibido: a,
		referencia: gvcomun.Vista{Primario: NODOCLIENTE2,
			Copia:    msgsys.HOSTINDEFINIDO,
			NumVista: 3},
	}

	// Comprobar vista tentativa recibida
	vac.comprobar()

	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) PrimarioACopiaTest5(t *testing.T) {
	//t.Skip("SKIPPED PrimerCopiaTest3")

	//latido del que es primario
	_, ok := ts.SendReceive(NODOCLIENTE2,
		gvcomun.MsgLatido{-1, NODOCLIENTE2},
		gvcomun.ANSWERWAITTIME*time.Millisecond,
	)
	if !ok {
		t.Fatal("Salta timeout esperando respuesta a latido -1 de cliente")
	}

	a := ts.clienteLatido0(t, NODOCLIENTE1)

	vac := vistasAcomparar{t: t,
		recibido: a,
		referencia: gvcomun.Vista{Primario: NODOCLIENTE2,
			Copia:    NODOCLIENTE1,
			NumVista: 4},
	}

	// Comprobar vista tentativa recibida
	vac.comprobar()

	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) EsperaACopiaTest6(t *testing.T) {
	//t.Skip("SKIPPED PrimerCopiaTest3")

	//se apunta del copia
	a := ts.clienteLatido0(t, NODOCLIENTE3)

	for i := 0; i < 5; i++ {
		ts.enviaLatidoVivo(t, NODOCLIENTE1)
		ts.enviaLatidoVivo(t, NODOCLIENTE3)
		time.Sleep(gvcomun.INTERVALOLATIDOS * time.Millisecond)
	}

	//pide vista tentativa
	a = ts.vistaTentativa(t, NODOCLIENTE1)

	vac := vistasAcomparar{t: t,
		recibido: a,
		referencia: gvcomun.Vista{Primario: NODOCLIENTE1,
			Copia:    NODOCLIENTE3,
			NumVista: 5},
	}

	// Comprobar vista tentativa recibida
	vac.comprobar()

	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) PrimarioAEsperaTest7(t *testing.T) {

	//notificamos vida de todos
	ts.enviaLatidoVivo(t, NODOCLIENTE1)
	ts.enviaLatidoVivo(t, NODOCLIENTE3)
	//se apunta del copia
	a := ts.clienteLatido0(t, NODOCLIENTE2)

	vac := vistasAcomparar{t: t,
		recibido: a,
		referencia: gvcomun.Vista{Primario: NODOCLIENTE1,
			Copia:    NODOCLIENTE3,
			NumVista: 5},
	}

	// Comprobar vista tentativa recibida
	vac.comprobar()

	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) VistaValidaNoValidadTest8(t *testing.T) {

	for i := 0; i < 4; i++ {
		//notificamos vida de todos
		ts.enviaLatidoVivo(t, NODOCLIENTE1)
		ts.enviaLatidoVivo(t, NODOCLIENTE3)
		ts.enviaLatidoVivo(t, NODOCLIENTE2)
	}

	//mirar valida
	p, ok := ts.SendReceive(NODOCLIENTE1,
		gvcomun.MsgPeticionVistaValida{NODOCLIENTE1},
		gvcomun.ANSWERWAITTIME*time.Millisecond)
	if !ok {
		t.Fatal("Salta timeout esperando respuesta a latido -1 de cliente")
	}

	var a gvcomun.Vista

	switch x := p.(type) {
	case gvcomun.MsgVistaValida:
		a = x.Vista // salida correcta
	}
	//La vista valida es distinta a la vista tentativa, pues está vacía
	vac := vistasAcomparar{t: t,
		recibido: a,
		referencia: gvcomun.Vista{Primario: msgsys.HOSTINDEFINIDO,
			Copia:    msgsys.HOSTINDEFINIDO,
			NumVista: 0},
	}

	// Comprobar vista tentativa recibida
	vac.comprobar()

	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) ErrorCritico9(t *testing.T) {

	//notificamos vida de todos
	ts.enviaLatidoVivo(t, NODOCLIENTE1)
	ts.enviaLatidoVivo(t, NODOCLIENTE3)
	ts.enviaLatidoVivo(t, NODOCLIENTE2)
	//validamos tentativa
	_, ok := ts.SendReceive(NODOCLIENTE1,
		gvcomun.MsgLatido{5, NODOCLIENTE1},
		gvcomun.ANSWERWAITTIME*time.Millisecond)
	if !ok {
		t.Fatal("Salta timeout esperando respuesta a latido -1 de cliente")
	}

	for i := 0; i < 8; i++ {
		//esperamos muerte de todos
		time.Sleep(gvcomun.INTERVALOLATIDOS*time.Millisecond)
	}

	_, err := net.Dial("tcp", string(NODOGV))
	if err == nil{
		t.Fatalf("se esperaba GV caido y está vivo")
	}


	fmt.Println(".............", t.Name(), "Superado")
}

func (ts *testServer) vistaTentativa(t *testing.T, nodo msgsys.HostPuerto) gvcomun.Vista {
	p, ok := ts.SendReceive(nodo,
		gvcomun.MsgLatido{NumVista: -1, Remitente: nodo},
		gvcomun.ANSWERWAITTIME*time.Millisecond)
	if !ok {
		t.Fatal("Salta timeout esperando respuesta a latido -1 de cliente")
	}
	//transformamos mensaje a vista tentativa

	var a gvcomun.Vista

	switch x := p.(type) {
	case gvcomun.MsgVistaTentativa:
		a = x.Vista // salida correcta
	}
	return a
}

func (ts *testServer) enviaLatidoVivo(t *testing.T, nodo msgsys.HostPuerto) bool {
	_, ok := ts.SendReceive(nodo,
		gvcomun.MsgLatido{-1, nodo},
		gvcomun.ANSWERWAITTIME*time.Millisecond,
	)
	if !ok {
		t.Fatal("Salta timeout esperando respuesta a latido -1 de cliente")
	}
	return ok
}

func (ts *testServer) clienteLatido0(t *testing.T,
	nodoCliente msgsys.HostPuerto) gvcomun.Vista {

	// solo nos interesa la vista tentativa devuelta por latido a Gestor Vistas
	m, ok := ts.SendReceive(nodoCliente,
		gvcomun.MsgLatido{Remitente: nodoCliente},
		gvcomun.ANSWERWAITTIME*time.Millisecond,
	)
	if !ok {
		t.Fatal("Salta timeout esperando latido 0 de cliente")
	}

	switch x := m.(type) {
	case gvcomun.MsgVistaTentativa:
		return x.Vista // salida correcta
	default:
		t.Fatalf(t.Name(),
			"Mensaje recibido INCORRECTO en primerPrimario: %#v", x)
	}

	// no debería llegar a ejecutarse, pero se pone por error compilacion
	return gvcomun.Vista{}
}

// --------------------------------------------------------------------------
// FUNCIONES DE APOYO

type vistasAcomparar struct {
	recibido   gvcomun.Vista
	referencia gvcomun.Vista
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
