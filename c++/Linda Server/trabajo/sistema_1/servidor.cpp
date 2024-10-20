#include <iostream>
#include <chrono>
#include <thread>
#include <signal.h>
#include "Socket.hpp"
#include "controlTuplas.hpp"

using namespace std;

const int MESSAGE_SIZE = 4001; //mensajes de no más 4000 caracteres
int SERVER_PORT;
bool out; // booleano que controla el cierre del servidor y clientes

//-----------------------------------------------------
void servCliente(Socket& soc, int cliente_fd, ControlTuplas& cT){
	char buf[MESSAGE_SIZE];
	string buffer;
	cout<<"Nuevo cliente"<<endl;
	while(!out) {
		// Recibimos el mensaje del Servidor Linda(cliente)
		// El formato del mensaje será: ACCION [TUPLA] NºELEMENTOS
		int rcv_bytes = soc.Recv(cliente_fd, buf, MESSAGE_SIZE);
		if (rcv_bytes == -1) {
			cerr << "Error al recibir datos: " << strerror(errno) << endl;
			// Cerramos los sockets
			soc.Close(cliente_fd);
		}

		if (strncmp(buf, "", 5) == 0) {
			//no se trata
		} else {

			buffer = string(buf);

			if (buffer == "TERMINACION") {
				cT.finalizar();
				out = true;
				//desbloquear del accept mediante un hijo o proceso finalizador
				Socket socket("0.0.0.0", SERVER_PORT);
				int socket_fd = socket.Connect();
				socket.Close(socket_fd);
			}
			else {
				string accion = buffer.substr(0,buffer.find_first_of("["));                // Guarda la acción a realizar
				buffer.erase(0, buffer.find_first_of("["));

				int nElementos = stoi(buffer.substr(buffer.find_first_of("]") + 1));        // Nº de elementos de la tupla

				buffer.erase(buffer.find_first_of("]") + 1);
				string informacion = buffer.substr(0);                                // Guarda la información de la tupla


				// Creamos la tupla recibida
				Tupla trecibida(nElementos);
				trecibida.from_string(informacion);
				bool error;
				int send_bytes;

				// Postnote
				if (accion == "PN") {
					bool error = false;
					cout << "Tupla recibida: " << trecibida.to_string() << endl;
					cT.anyadir(trecibida,error);
					if(!error){
						cout << "Tupla añadida: " << trecibida.to_string() << endl;
						// Se envia mensaje de que la operacion ha sido correcta
						send_bytes = soc.Send(cliente_fd, "PN correcto");
					}
					else{
						cout << "Tupla no añadida: " << trecibida.to_string() << endl;
						// Se envia mensaje de que la operacion no ha sido correcta
						send_bytes = soc.Send(cliente_fd, "PN incorrecto, no se pueden insertar patrones");
					}
				}
					// Removenote
				else if (accion == "RN") {
					cout << "Tupla recibida: " << trecibida.to_string() << endl;
					cT.eliminar(trecibida);
					cout << "Tupla eliminada: " << trecibida.to_string() << endl;
					// Seguidamente se envia la tupla
					send_bytes = soc.Send(cliente_fd, trecibida.to_string());
					if (send_bytes == -1) {
						cerr << "Error al enviar datos: " << strerror(errno) << endl;
						// Cerramos el socket
						soc.Close(cliente_fd);
						exit(1);
					}
				}
					// Readnote
				else if (accion == "ReadN") {
					cout << "Tupla recibida: " << trecibida.to_string() << endl;
					cT.leer(trecibida);
					cout << "Tupla leida: " << trecibida.to_string() << endl;
					// Seguidamente se envia la tupla
					send_bytes = soc.Send(cliente_fd, trecibida.to_string());
					if (send_bytes == -1) {
						cerr << "Error al enviar datos: " << strerror(errno) << endl;
						// Cerramos el socket
						soc.Close(cliente_fd);
						exit(1);
					}
				}
			}
			// Si no existe la opcion-> enviar mensaje de opcion no existente(?)
		}
	}
	soc.Close(cliente_fd);
	cout<<"Fin cliente"<<endl;
}
//-----------------------------------------------------
int main(int argc, char* argv[]) {

	signal(SIGPIPE, SIG_IGN);
	//No se ha llamado al programa con los parametros adecuados
	if(argc != 2){
		cout << "Uso: servidor <IP del servidor> <puerto del servidor>" << endl;
		return 1;
	}

	// Creamos el monitor que controla las tuplas en memoria
	ControlTuplas cT;

	// Puerto donde escucha el proceso servidor
	SERVER_PORT = atoi(argv[1]);
	int max_connections = 10;

	// Creación del socket con el que se llevará a cabo
	// la comunicación con el servidor de Linda.
	Socket socket(SERVER_PORT);

	// Bind
	int socket_fd =socket.Bind();
	if (socket_fd == -1) {
		string mensError(strerror(errno));
		cerr << "Error en el bind: " + mensError + "\n";
		exit(1);
	}

	cout << "*Servidor conectado*" << endl;

	// Listen
	int error_code = socket.Listen(max_connections);
	if (error_code == -1) {
		string mensError(strerror(errno));
		cerr << "Error en el listen: " + mensError + "\n";
		// Cerramos el socket
		socket.Close(socket_fd);
		exit(1);
	}

	out = false;
	while(!out){
		// Accept
		int cliente_fd = socket.Accept();

		if(cliente_fd == -1) {
			string mensError(strerror(errno));
			cerr << "Error en el accept: " + mensError + "\n";
			// Cerramos el socket
			socket.Close(socket_fd);
			exit(1);
		}

		if(!out){
			thread t(&servCliente, ref(socket), cliente_fd, ref(cT));
			t.detach();
		}
	}

	error_code = socket.Close(socket_fd);
	if (error_code == -1) {
		string mensError(strerror(errno));
		cerr << "Error cerrando el socket del servidor: " + mensError + "\n";
	}

	cout<<"Bye bye"<<endl;
}