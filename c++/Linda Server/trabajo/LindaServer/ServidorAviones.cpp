//******************************************************************
// File:   ServidorMulticliente.cpp
// Author: PSCD-Unizar
// Date:   Noviembre 2015
// Coms:   Ejemplo de servidor multicliente con comunicación síncrona mediante sockets
//*****************************************************************

#include "Socket/Socket.hpp"
#include <iostream>
#include <thread>
#include <string>
#include "ControlAviones.hpp"
#include "ControlAviones.cpp"	//OJO
#include <cstring> //manejo de cadenas tipo C
#include <csignal> //manejo de señales tipo C


using namespace std;
ControlAviones cA(16);
int puerto;
//-------------------------------------------------------------
void signal_handler(int signal)
{		bool ignorar;
		cA.acabo(ignorar);
		//si ya me habian finalizado, tengo que ignorar todos estos pasos
		if(!ignorar){
			cout<<"\n Se ha cerrado el servidor, esperando a los clientes"<<endl;
			//desbloquear del accept mediante un hijo o proceso finalizador
			Socket socket("0.0.0.0", puerto);
			int socket_fd = socket.Connect();
			socket.Close(socket_fd);
			//mientras tanto el padre (servidor) altera su monitor
			cA.finalizacion();
		}
}
//-------------------------------------------------------------
void servCliente(Socket& soc, int client_fd, ControlAviones& cA) {
	cA.entraCliente();	//contabiliza la entrada de un nuevo cliente
	char MENS_FIN[]="END OF SERVICE";
	// Buffer para recibir el mensaje
	int length = 100;
	char buffer[length];
	//*****************************************************	
	//declaración avión
	const int N = 16;
	bool asientos[N];
	for(int i = 0; i < N; i++)
	{
		asientos[i] = true;
	}

	bool out = false; // Inicialmente no salir del bucle
	//*****************************************************
	while(!out) {
		// Recibimos el mensaje del cliente
		int rcv_bytes = soc.Recv(client_fd,buffer,length);
		if (rcv_bytes == -1) {
			string mensError(strerror(errno));
    		cerr << "Error al recibir datos: " + mensError + "\n";
			// Cerramos los sockets
			soc.Close(client_fd);
			out = true;
		}

		cout << "Mensaje recibido: " << buffer << endl;
		string sbuffer=(string)buffer;
		string message;
		// Si recibimos "END OF SERVICE" --> Fin de la comunicación
		if (0 == strcmp(buffer, MENS_FIN)) {
			out = true; // Salir del bucle
		}else{string str2 = sbuffer.substr(0, 13);
		 	if(str2 == "INICIO COMPRA")
		 	{//caso de bienvenida
		 	message = "BIENVENIDO AL SERVICIO";
		 	}else {//caso de no bienvenida y no final
		 		bool reserva = true; 
		 		bool lleno = false; 
				cA.adivinaAsiento(sbuffer, reserva, lleno);
				if(reserva)
				{	 
					message = "RESERVADA:";
					message.append(sbuffer); //reservada + fila/columna 
				}else if(lleno)
				{//si esta lleno, final
					message = "VUELO COMPLETO";
					out = true;
				}else if(!reserva && !lleno)
				{//se ha buscado uno ocupado
					cA.mostrar(message);
				}	
			}
		}	
			int send_bytes = soc.Send(client_fd, message);
			if(send_bytes == -1) {
				string mensError(strerror(errno));
    			cerr << "Error al enviar datos: " + mensError + "\n";
				// Cerramos los sockets
				soc.Close(client_fd);
				out = true;
			}
		
	}	cA.saleCliente(false);
	
	soc.Close(client_fd);
}
//-------------------------------------------------------------
int main(int argc, char * argv[]) {	
	const int N = 100;
	// Dirección y número donde escucha el proceso servidor
	string SERVER_ADDRESS = "0.0.0.0";
    int SERVER_PORT = atoi(argv[1]);
    int numProcess = 0;		//Numero de procesos en activo
    int client_fd;
	
	puerto = SERVER_PORT;
	
	// Creación del socket con el que se llevará a cabo
	// la comunicación con el servidor.
	Socket socket(SERVER_PORT);

	// Bind 
	int socket_fd = socket.Bind();
	if (socket_fd == -1) {
		string mensError(strerror(errno));
    	cerr << "Error en el bind: " + mensError + "\n";
    	exit(1);
	}

	// Listen
    int max_connections = N;
	int error_code = socket.Listen(max_connections);
	if (error_code == -1) {
		string mensError(strerror(errno));
    	cerr << "Error en el listen: " + mensError + "\n";
		// Cerramos el socket
		socket.Close(socket_fd);
	}
	//bucle servidor************************************************
	int i = 1;
	bool salgo = false;
	while (!salgo) {
		signal(SIGPIPE, SIG_IGN);		//ignora los ctrl+c de los clientes
		signal(SIGINT, signal_handler);	//si le llega ctrl+c al servidor la recoge e interpreta.
		signal(SIGTSTP, SIG_IGN);
		
		// Accept
		client_fd = socket.Accept();
		// Durante el bloqueo del accept me han finalizado?
		cA.acabo(salgo);
		
		if(client_fd == -1) {
			string mensError(strerror(errno));
    		cerr << "Error en el accept: " + mensError + "\n";
			// Cerramos el socket
			socket.Close(socket_fd);
		} 
		
		if(!salgo)
		{
//si tiene que salir, no añade threads y simplemente sale del while.
			cout << "Lanzo thread nuevo cliente " + to_string(i) + "\n";
			thread t = thread(&servCliente, ref(socket), client_fd, ref(cA));
			t.detach();
			cout << "Nuevo cliente " + to_string(i) + " aceptado" + "\n";
			cA.nClientes(numProcess);
			i++;
		}
	}
	
    // Cerramos el socket del servidor
    error_code = socket.Close(socket_fd);
    if (error_code == -1) {
		string mensError(strerror(errno));
    	cerr << "Error cerrando el socket del servidor: " + mensError + "\n";
	}
	
	// Despedida
	cout << "Bye bye" << endl;

    return error_code;
}
//-------------------------------------------------------------
