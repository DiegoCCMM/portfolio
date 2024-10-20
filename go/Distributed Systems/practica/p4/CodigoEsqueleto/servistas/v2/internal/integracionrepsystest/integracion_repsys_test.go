package integracionrepsystest

import (
	"fmt"
	"log"
	"math/rand"
	"servistas/v2/internal/gvcomun"
	"servistas/v2/internal/msgsys"
	"servistas/v2/pkg/cltalm"
	"strconv"
	"testing"
	"time"
)

const (
	//hosts 127.0.0.1
	MAQUINA1 = "155.210.154.200"
	MAQUINA2 = "155.210.154.201"
	MAQUINA3 = "155.210.154.202"
	MAQUINA4 = "155.210.154.202"		//155.210.154.201

	//puertos
	PUERTOTEST = "29000"
	PUERTOGV   = "29001"
	PUERTOCL1  = "29222"
	PUERTOCL2  = "29223"
	PUERTOCL3  = "29224"
	PUERTOSA1  = "29305"	//puertos SA
	PUERTOSA2  = "29296"
	PUERTOSA3  = "29297"

	//nodos test, gv y clientes en la misma maquina
	NODOTEST = MAQUINA1 + ":" + PUERTOTEST
	NODOGV   = MAQUINA1 + ":" + PUERTOGV
	NODOCL1  = MAQUINA1 + ":" + PUERTOCL1
	NODOCL2  = MAQUINA1 + ":" + PUERTOCL2
	NODOCL3  = MAQUINA1 + ":" + PUERTOCL3
	NODOSA1  = MAQUINA2 + ":" + PUERTOSA1
	NODOSA2  = MAQUINA3 + ":" + PUERTOSA2
	NODOSA3  = MAQUINA4 + ":" + PUERTOSA3

	// PATH de los ejecutables de modulo golang de servicio de vistas
	PATH = "/home/a738712/SSDD/p5/v2/cmd "	//~/Desktop/info/4-1/SSDD/p4/CodigoEsqueleto/servistas/v2/cmd

	// fichero main de ejecutables relativos a PATH previo
	EXECGV     = "cmdsrvvts/main.go " + NODOGV // un parámetro
	EXECSRVALM = "cmdsrvalm/main.go "          // 2 parámetros en llamadas ssh

	// comandos completo a ejecutar en máquinas remota con ssh
	// cd /home/tmp/servistas/v2/cmd/; go run cmdsrvvts/main.go 127.0.0.1:29001
	SRVVTSCMD = "cd " + PATH + "; /usr/local/go/bin/go run " + EXECGV
	// cd /home/tmp/servistas/v2/cmd/;
	// go run testcltvts/main.go 127.0.0.1:29003 127.0.0.1:29001 127.0.0.1:29000
	SRVALMCMD = "cd " + PATH + "; /usr/local/go/bin/go run " + EXECSRVALM

	// Ubicar, en esta constante, el PATH completo a vuestra clave privada local
	// emparejada con la clave pública en authorized_keys de máquinas remotas

	PRIVKEYFILE = "/home/a738712/.ssh/id_rsa"	///home/diego/.ssh/id_rsa
)

// TEST primer rango
func TestPrimerasPruebas(t *testing.T) { // (m *testing.M) {
	// <setup code>
	// Crear servidor de test y
	//procesos en maquinas remotas : servidor de vistas, servidores
	ts := startTestServer(NODOTEST)

	// Run test sequence

	// Test1 : No debería haber ningun primario, si SV no ha recibido aún latidos
	t.Run("T1:SoloArranqueYParada",
		func(t *testing.T) { ts.soloArranqueYparadaTest1(t) })

	// Test2: tenemos el primer primario correcto
	t.Run("T2:AlgunasEscriturasYFalloCopia",
		func(t *testing.T) { ts.AlgunasEscriturasYFalloCopiaTest2(t) })

	// Test3: Primer nodo copia
	t.Run("T3:EscriturasConcurrentesDeClientesYFalloPrimario",
		func(t *testing.T) { ts.writeConcurVariosClientCaidaPrimarioTest3(t) })

	// tear down code
	// eliminar procesos en máquinas remotas
	ts.stop()
}

// ---------------------------------------------------------------------
// Servidor de test

type testServer struct {
	msgsys.MsgSys
	// Canal de resultados de ejecución de comadnos ssh remotos
	cmdOutput chan string
}

// El servidor de tests se comportará como un cliente de almacenammiento
// desde u
func startTestServer(me msgsys.HostPuerto) (ts testServer) {
	// Registrar tipos de mensaje de cliente de almacenamiento
	// El servidor de tests se comportará, tambien, como cliente de almacenam.
	gvcomun.RegistrarMensajesCLTSA()

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
var NODOSA1V string
// start  gestor de vistas; mapa de replicas y maquinas donde ubicarlos;
// y lista clientes (host:puerto)
func (ts *testServer) startDistributedProcesses(clientes []string,
	replicasMaquinas map[string]string) (nodoscliente []cltalm.ClienteAlm) {

	// Poner en marcha servidor/gestor de vistas
	/*cltssh.ExecMutipleHosts(SRVVTSCMD,
		[]string{MAQUINA1}, ts.cmdOutput, PRIVKEYFILE)

	// esperar que gestor de vistas este en marcha
	time.Sleep(2000 * time.Millisecond)


		cltssh.ExecMutipleHosts(SRVALMCMD+NODOSA1+" "+NODOGV,
			[]string{MAQUINA1}, ts.cmdOutput, PRIVKEYFILE)
	time.Sleep(4000 * time.Millisecond)
	cltssh.ExecMutipleHosts(SRVALMCMD+NODOSA2+" "+NODOGV,
			[]string{MAQUINA2}, ts.cmdOutput, PRIVKEYFILE)
	time.Sleep(4000 * time.Millisecond)
	cltssh.ExecMutipleHosts(SRVALMCMD+NODOSA3+" "+NODOGV,
			[]string{MAQUINA3}, ts.cmdOutput, PRIVKEYFILE)
	time.Sleep(4000 * time.Millisecond)
*/
		// dar tiempo para se establezacn primario, copia y nodos en espera
		// en la secuencia suministrada


	// Crear los msgsys para clientes locales con un puerto especifico cada uno
	//TODO utilizar para lanzar los clientes sólos
	for _, clientHost := range clientes {
		nodoscliente = append(nodoscliente,
			cltalm.Start(msgsys.HostPuerto(clientHost),
				msgsys.HostPuerto(NODOGV)))
	}

	return nodoscliente
}

//
func (ts *testServer) stopDistributedProcesses(nodosCliente []cltalm.ClienteAlm,
	replicas []string) {

	// Parar procesos distribuidos con ssh
	// una opción :
	for _, cliente := range nodosCliente {
		cliente.Stop()
	}

	time.Sleep(10 * time.Millisecond)
/*
	for _, nodoSA := range replicas {
		ts.Send(msgsys.HostPuerto(nodoSA), gvcomun.MsgFin{})
	}
*/
	// esperar parada se servidores remotos el tiempo suficiente
	// para volcar salida de ejecuciones ssh en cmdOutput
	time.Sleep(100 * time.Millisecond)

	//ts.Send(msgsys.HostPuerto(NODOGV), gvcomun.MsgFin{})
}

// --------------------------------------------------------------------------
// FUNCIONES DE SUBTESTS

var nodosCliente []cltalm.ClienteAlm

// No debería haber primario
func (ts *testServer) soloArranqueYparadaTest1(t *testing.T) {
	//t.Skip("SKIPPED soloArranqueYparadaTest1")

	fmt.Println(t.Name(), ".....................")

	// Poner en marcha gestor de vistas y réplicas de almacenamiento en remoto
	nodosCliente = ts.startDistributedProcesses([]string{NODOCL1},
		map[string]string{NODOSA1: MAQUINA2, NODOSA2: MAQUINA3, NODOSA3: MAQUINA4})
	time.Sleep(4000 * time.Millisecond)

	// Parar réplicas alamcenamiento en remoto
	ts.stopDistributedProcesses(nodosCliente, []string{NODOSA1,NODOSA2,NODOSA3})

	fmt.Println(".............", t.Name(), "Superado")
}

// No debería haber primario
func (ts *testServer) AlgunasEscriturasYFalloCopiaTest2(t *testing.T) {
	t.Skip("SKIPPED algunasEscriturasVariasReplicasTest2")

	fmt.Println(t.Name(), ".....................")

	// Poner en marcha un clientes local, gestor de vistas en local
	// y 3 réplicas de almacenamiento en remoto
	clientes := []string{NODOCL1}
	replicasMaquinas :=
		map[string]string{NODOSA1: MAQUINA2, NODOSA2: MAQUINA3, NODOSA3: MAQUINA4}
	nodosCliente := ts.startDistributedProcesses(clientes, replicasMaquinas)


	fmt.Println("el primero")
	// Comprobar escrituras en configuración completa inicial (primario, copia)
	ComprobarWritesConfCompleta(nodosCliente[0], NODOSA2, NODOSA3, t)

	fmt.Println("el segundo")
	// Comprobar una escritura trás fallo de nodo copia (NODOSA2 parado)
	// con reemplazo de nodo en espera
	comprobarWriteTrasFalloCopia(nodosCliente[0], NODOSA1, NODOSA3, t)
	fmt.Println("el tercero")
	// Parar réplicas alamcenamiento en remoto (Copia ya había sido detenida)
	ts.stopDistributedProcesses(nodosCliente, nil)

	fmt.Println(".............", t.Name(), "Superado")
}

// escritura concurrente con 3 clientes, 3 replicas y caida de copia
func (ts *testServer) writeConcurVariosClientCaidaPrimarioTest3(t *testing.T) {
	//t.Skip("SKIPPED writeConcurVariosClientCaidaPrimarioTest3")

	fmt.Println(t.Name(), ".....................")

	// Poner en marcha gestor de vistas y réplicas de almacenamiento en remoto
	// Y clientes de almacenamiento en local escuchando diferentes puertos
	clientes := []string{NODOCL1, NODOCL2, NODOCL3}
	replicasMaquinas :=
		map[string]string{NODOSA1: MAQUINA2,
			NODOSA2: MAQUINA3, NODOSA3: MAQUINA4}
	nodosCliente := ts.startDistributedProcesses(clientes, replicasMaquinas)

	nodosCliente[0].GetPrimario()
	nodosCliente[1].GetPrimario()
	nodosCliente[2].GetPrimario()
	nodosCliente[0].ObtenerPrimario()
	// Escribimos con 3 clientes concurrentes,
	// datos aleatorios en claves "0" y "1"
	escrituraConcurrente(nodosCliente)

	time.Sleep(200 * time.Millisecond) // Wait to finish last writes

	// obtener valor de las claves "0" y "1" con el primer primario funcionando
	valor0Primario, _ := nodosCliente[0].Read("0")
	valor1Primario, _ := nodosCliente[2].Read("1")
	fmt.Println("primeros read")
	// Forzar parada del primario en curso
	nodosCliente[0].Send(nodosCliente[0].GetPrimario(), gvcomun.MsgFin{})

	// Esperar detección de fallo y reconfiguración
	time.Sleep(700 * time.Millisecond)
	nodosCliente[0].GetPrimario()
	nodosCliente[1].GetPrimario()
	nodosCliente[2].GetPrimario()
	nodosCliente[0].ObtenerPrimario()
	// Obtener valor de clave "0" y "1" con 2º primario en marcha(copia anterior)
	valor0Copia, _ := nodosCliente[2].Read("0")
	valor1Copia, _ := nodosCliente[1].Read("1")

	// Comprobar valores obtenidos con primario inicial y
	// con Copia (2º primario)
	switch {
	case valor0Primario != valor0Copia:
		t.Fatalf("%s: Clave '0': Valores: con Primario: (%s),  con Copia: (%s)",
			t.Name(), valor0Primario, valor0Copia)

	case valor1Primario != valor1Copia:
		t.Fatalf("%s: Clave '1': Valores: con Primario: (%s),  con Copia: (%s)",
			t.Name(), valor1Primario, valor1Copia)
	}

	// Parar réplicas almacenamiento en remoto (primer primario ya estaba parado)
	ts.stopDistributedProcesses(nodosCliente,
		[]string{NODOSA2, NODOSA3})

	fmt.Println(".............", t.Name(), "Superado")
}


// escritura concurrente con 3 clientes, 3 replicas y caida de copia
/*
func (ts *testServer) writeConcurVariosClientCaidaPrimarioYCopiaTest4(t *testing.T) {
	//t.Skip("SKIPPED writeConcurVariosClientCaidaPrimarioYCopiaTest4")

	fmt.Println(t.Name(), ".....................")

	// Poner en marcha gestor de vistas y réplicas de almacenamiento en remoto
	// Y clientes de almacenamiento en local escuchando diferentes puertos
	clientes := []string{NODOCL1, NODOCL2, NODOCL3}
	replicasMaquinas :=
		map[string]string{NODOSA1: MAQUINA2,
			NODOSA2: MAQUINA3, NODOSA3: MAQUINA4}
	nodosCliente := ts.startDistributedProcesses(clientes, replicasMaquinas)

	// Escribimos con 3 clientes concurrentes,
	// datos aleatorios en claves "0" y "1"
	escrituraConcurrente(nodosCliente)

	time.Sleep(200 * time.Millisecond) // Wait to finish last writes

	// obtener valor de las claves "0" y "1" con el primer primario funcionando
	valor0Primario, _ := nodosCliente[0].Read("0")
	valor1Primario, _ := nodosCliente[2].Read("1")

	// Forzar parada del primario en curso
	nodosCliente[0].Send(nodosCliente[0].GetPrimario(), gvcomun.MsgFin{})

	// Esperar detección de fallo y reconfiguración
	time.Sleep(700 * time.Millisecond)

	// Obtener valor de clave "0" y "1" con 2º primario en marcha(copia anterior)
	valor0Copia, _ := nodosCliente[2].Read("0")
	valor1Copia, _ := nodosCliente[1].Read("1")

	// Comprobar valores obtenidos con primario inicial y
	// con Copia (2º primario)
	switch {
	case valor0Primario != valor0Copia:
		t.Fatalf("%s: Clave '0': Valores: con Primario: (%s),  con Copia: (%s)",
			t.Name(), valor0Primario, valor0Copia)

	case valor1Primario != valor1Copia:
		t.Fatalf("%s: Clave '1': Valores: con Primario: (%s),  con Copia: (%s)",
			t.Name(), valor1Primario, valor1Copia)
	}

	// Parar réplicas almacenamiento en remoto (primer primario ya estaba parado)
	ts.stopDistributedProcesses(nodosCliente,
		[]string{NODOSA2, NODOSA3})

	fmt.Println(".............", t.Name(), "Superado")
}
*/

// --------------------------------------------------------------------------
// FUNCIONES DE APOYO

// Comprobar escrituras en configuración completa inicial (primario, copia)
func ComprobarWritesConfCompleta(nodoCliente cltalm.ClienteAlm,
	primario, copia string,
	t *testing.T) {

	fmt.Println(t.Name(),
		"Comprobar escrituras con primario, copia y nodo en espera.........")

	// comprobar vista valida correcta
	mensajeRecibido, ok := nodoCliente.SendReceive(
		NODOGV,
		gvcomun.MsgPeticionVistaValida{Remitente: nodoCliente.Me()},
		gvcomun.ANSWERWAITTIME*time.Millisecond)
	if !ok {
		t.Fatalf("No consigue recibir vista valida en subtest %s", t.Name())
	}

	 log.Printf("Mensaje recibido en ComprobarWritesConfCompleta : %#v",
	mensajeRecibido)

	log.Printf("La vista ha sido validada")
	nodoCliente.GetPrimario()
	// Comprobar algunas escrituras iniciales en configuración completa inicial
	nodoCliente.Write("a", "aa")
	time.Sleep(500*time.Millisecond)
	nodoCliente.Write("b", "bb")
	time.Sleep(500*time.Millisecond)
	nodoCliente.Write("c", "cc")
	time.Sleep(500*time.Millisecond)



	res, _ := nodoCliente.Read("a")

	if(res != "aa"){
		fmt.Println(res)
	}else{
		fmt.Println("TODO VA BIEN" )
		fmt.Println(res)
	}

	fmt.Println("......... Completado")
}

// Comprobar una escritura trás fallo de nodo copia con reemplazo en espera
func comprobarWriteTrasFalloCopia(nodoCliente cltalm.ClienteAlm,
	primario, copia string,
	t *testing.T) {

	fmt.Println(t.Name(),
		"Comprobar escritura despues de fallo de nodo Copia ...........")

	// Detener nodo réplica copia
	mensajeRecibido, ok := nodoCliente.SendReceive(
		NODOGV,
		gvcomun.MsgPeticionVistaValida{Remitente: nodoCliente.Me()},
		gvcomun.ANSWERWAITTIME*time.Millisecond)
	if !ok {
		t.Fatalf("No consigue recibir vista valida en subtest %s", t.Name())
	}

	fmt.Println(mensajeRecibido.(gvcomun.MsgVistaValida).Vista, mensajeRecibido.(gvcomun.MsgVistaValida).Vista.Copia, "HOLA")

	nodoCliente.Send(mensajeRecibido.(gvcomun.MsgVistaValida).Vista.Copia,
		gvcomun.MsgFin{})


	fmt.Println(mensajeRecibido.(gvcomun.MsgVistaValida).Vista, mensajeRecibido.(gvcomun.MsgVistaValida).Vista.Copia, "HOLA2 pasado fin")


	// Esperar tiempo necesario a la reconfiguración de réplicas
	time.Sleep(1000 * time.Millisecond)

	//log.Println("PREVIO COnsulta vista valida despues de parar SA2 !!!!!")
	// Comprobar nueva configuración primario y copia
	mensajeRecibido, ok = nodoCliente.SendReceive(
		NODOGV,
		gvcomun.MsgPeticionVistaValida{Remitente: nodoCliente.Me()},
		gvcomun.ANSWERWAITTIME*time.Millisecond)
	if !ok {
		t.Fatalf("No consigue recibir vista valida en subtest %s", t.Name())
	}

	fmt.Println(mensajeRecibido.(gvcomun.MsgVistaValida).Vista, mensajeRecibido.(gvcomun.MsgVistaValida).Vista.Copia, "HOLA3 pasado fin")

	/*	comprobarVistaValida(
		mensajeRecibido.(gvcomun.MsgVistaValida).Vista,
		gvcomun.Vista{Primario: msgsys.HostPuerto(primario),
			Copia: msgsys.HostPuerto(copia)},
		t, "Desde comprobarWriteTrasFalloCopia")*/
	nodoCliente.GetPrimario()
	// comprobar una escritura
	nodoCliente.Write("a", "aaaaaaaaa")
	//comprobarClaveValor(nodoCliente, "a", "aaaaaaaaa", t)

	res, _ := nodoCliente.Read("a")

	if(res != "aaaaaaaaa"){
		fmt.Println(res)
	}else{
		fmt.Println("TODO VA BIEN" )
		fmt.Println(res)
	}

	fmt.Println("......... Completado")
}

func comprobarVistaValida(recibido, esperado gvcomun.Vista, t *testing.T,
	who string) {

	switch {
	case recibido.Primario != esperado.Primario:
		t.Fatalf(
			"%s : %s : PRIMARIO recibido (%s) y de referencia (%s) no coinciden",
			t.Name(), who, recibido.Primario, esperado.Primario)

	case recibido.Copia != esperado.Copia:
		t.Fatalf("%s: %s : COPIA recibido (%s) y de referencia (%s) no coinciden",
			t.Name(), who, recibido.Copia, esperado.Copia)
	}
}

func comprobarClaveValor(clienteAlm cltalm.ClienteAlm,
	clave, valorEsperado string,
	t *testing.T) {

	valorEnAlmacen, ok := clienteAlm.Read(clave)
	if !ok {
		t.Fatalf(
			"Timeout Read de cliente, función comprobarClaveValor, subtest %s",
			t.Name())
	}

	if valorEnAlmacen != valorEsperado {
		t.Fatalf(
			"%s : Valor en almacen (%s) y valor esperado (%s) no coinciden",
			t.Name(), valorEnAlmacen, valorEsperado)
	}
}

func escrituraConcurrente(clientes []cltalm.ClienteAlm) {
	done := make(chan struct{})
	rand.Seed(time.Now().UnixNano()) // Generar semilla para valores aleatorios
	clientes[0].GetPrimario()
	clientes[1].GetPrimario()
	clientes[2].GetPrimario()
	go escrituraContinua(clientes[0], rand.Intn(300), done) // primer cliente

	go escrituraContinua(clientes[1], rand.Intn(300), done) // segundo cliente

	go escrituraContinua(clientes[2], rand.Intn(300), done) // tercer cliente

	time.Sleep(2 * time.Second)

	close(done) // los 3 clientes concurrentes deben terminar sus escrituras
}

// escritura continua con valores aleatorios en claves "0" y "1"
func escrituraContinua(cliente cltalm.ClienteAlm,
	valorAleatorio int, done chan struct{}) {

	clave := strconv.Itoa(valorAleatorio % 2) // solo claves "0" o "1"
	valor := strconv.Itoa(valorAleatorio)
	cliente.GetPrimario()
bucle:
	for {
		select {
		case <-done:
			fmt.Println("Acabar escrituraContinua")
			break bucle
		default:
			time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
			cliente.Write(clave, valor)
		}
	}
}
