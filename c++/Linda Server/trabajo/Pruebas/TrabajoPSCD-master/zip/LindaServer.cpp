#include <iostream>
#include <thread>
#include <pthread.h>
#include <chrono>
#include <signal.h>
#include <cstring>
#include <string>
#include "Socket.hpp"
#include "Tuplas.hpp"
#include "ControlLS.hpp"

using namespace std;

string SERVER_ADDRESS1 = "0.0.0.0";  //el que se ocupe del tamaño 1-3
string SERVER_ADDRESS2 = "0.0.0.0";  //el que se ocupe del tamaño 4-5
string SERVER_ADDRESS3 = "0.0.0.0";  //el que se ocupe del tamaño 6
const int SERVER_PORT = 5000;      //el puerto que se conectan los clientes (ls se comporta como servidor)
const int MEM_PORT1 = 5001;      //el puerto en el que recibe de los servidores (ls se comporta como cliente)
const int MEM_PORT2 = 5002;
const int MEM_PORT3 = 5003;
ControlLS cLS;

//soc = client_port, soc2 = server_port
void servCliente(Socket &soc, int client_fd) {
    signal(SIGPIPE, SIG_IGN);
    cLS.entraCliente();
    // Buffer para recibir el mensaje
    int length = 100;
    char buffer[length];
    Socket soc2(SERVER_ADDRESS1, MEM_PORT1);        //los tres distintos servidores
    Socket soc3(SERVER_ADDRESS2, MEM_PORT2);
    Socket soc4(SERVER_ADDRESS3, MEM_PORT3);
    int socket_fd1;
    int socket_fd2;
    int socket_fd3;
    bool out = false; // Inicialmente no salir del bucle

    // Conectamos con el servidor 1. Probamos varias conexiones
    const int MAX_ATTEMPS = 10;
    int count = 0;
    do {
        // Conexión con el servidor 1
        socket_fd1 = soc2.Connect();
        count++;
        // Si error --> esperamos 1 segundo para reconectar
        if (socket_fd1 == -1) {
            cout << "error al conectar" << endl;
            this_thread::sleep_for(chrono::seconds(1));
        }
    } while ((socket_fd1 == -1) && count < MAX_ATTEMPS);
    // Conectamos con el servidor2. Probamos varias conexiones
    do {  //todo abrir estos connects, en pruebas solo un servidor
        // Conexión con el servidor 2
        socket_fd2 = soc3.Connect();
        count++;
        // Si error --> esperamos 1 segundo para reconectar
        if (socket_fd2 == -1) {
            cout << "error al conectar" << endl;
            this_thread::sleep_for(chrono::seconds(1));
        }
    } while ((socket_fd2 == -1) && count < MAX_ATTEMPS);
    // Conectamos con el servidor. Probamos varias conexiones
    do {      //todo abrir estos connects, en pruebas solo un servidor
    // Conexión con el servidor 3
      socket_fd3 = soc4.Connect();
      count++;
    // Si error --> esperamos 1 segundo para reconectar
        if (socket_fd3 == -1) {
            cout << "error al conectar" << endl;
            this_thread::sleep_for(chrono::seconds(1));
        }
    } while ((socket_fd3 == -1) && count < MAX_ATTEMPS);

    while (!out) {
        // Recibimos el mensaje del cliente
        int rcv_bytes = soc.Recv(client_fd, buffer, length);
        if (rcv_bytes == -1) {
            string mensError(strerror(errno));
            cerr << "Error al recibir datos: " + mensError + "\n";
            // Cerramos los sockets
            soc.Close(client_fd);
            out = true;
        }else if(rcv_bytes != 0) {

            //antes de procesar el mensaje, igual teniamos que salir y estaba bloqueado
            cLS.salida(out);
            if (strncmp(buffer, "DESCONECTAR CLIENTE", 19) == 0) {
                out = true;
                soc.Close(client_fd);
                soc2.Close(socket_fd1);
                soc3.Close(socket_fd2);
                soc4.Close(socket_fd3);
            }
            //si hay que salir, me salgo, sino proceso el mensaje
            if (!out) {
                cout << "-----------------------------------------" << endl;
                cout << "Mensaje recibido: " << buffer << endl;
                //si el mensaje recibido es el de finalización de servidor o el de sacar a todos los clientes, lo procesamos
                if (strncmp(buffer, "FIN_SERVER", 10) == 0) {
                    cout << "\nESPERANDO A QUE FINALICEN LOS CLIENTES" << endl;
                    cout << "SE CIERRA LA ENTRADA" << endl;

                    cout << "\n Se ha cerrado el servidor, esperando a los clientes" << endl;
                    string buf = "TERMINACION";

                    int send_bytes = soc2.Send(socket_fd1, buf);
                    if (send_bytes == -1) {
                        cout << "error al enviar a servidores" << endl;
                        string mensError(strerror(errno));
                        cerr << "Error al enviar datos: " + mensError + "\n";
                        // Cerramos los sockets
                        soc2.Close(socket_fd1);
                        out = true;
                    }

                    int send_bytes1 = soc3.Send(socket_fd2, buf);      //todo caso de tener que cerrar tercer servidor
                    if (send_bytes1 == -1) {
                        cout << "error al enviar a servidores" << endl;
                        string mensError(strerror(errno));
                        cerr << "Error al enviar datos: " + mensError + "\n";
                        // Cerramos los sockets
                        soc3.Close(socket_fd2);
                        out = true;
                    }

                    int send_bytes2 = soc4.Send(socket_fd3, buf);      //todo caso de tener que cerrar cuarto servidor
                    if (send_bytes2 == -1) {
                        cout << "error al enviar a servidores" << endl;
                        string mensError(strerror(errno));
                        cerr << "Error al enviar datos: " + mensError + "\n";
                        // Cerramos los sockets
                        soc4.Close(socket_fd3);
                        out = true;
                    }

                    cLS.finalizacion();
                    //desbloquear del accept mediante un hijo o proceso finalizador
                    Socket socket("0.0.0.0", SERVER_PORT);
                    int soket_fd = socket.Connect();
                    socket.Close(soket_fd);
                    out = true;
                } else {//si el mensaje recibido no es el de finalización, entonces procesamos el mensaje de LD
                    string sbuffer = buffer;
                    //formato del mensaje recibido : ACCION [TUPLA] TAM_TUPLA
                    string sn = sbuffer.substr(sbuffer.find_last_of("]"), sbuffer.length());   //n tamaño de la tupla
                    sn.erase(sn.begin());     //elimina "]"
                    int n = stoi(sn);
                    string inst = sbuffer.substr(0, sbuffer.find_first_of("["));   //instruccion de la tupla
                    Tupla t(n);
                    string info = sbuffer.substr(sbuffer.find_first_of("["),
                                                 sbuffer.find_first_of("]"));   //info tupla en si
                    info.pop_back();
                    t.from_string(info);    //t es la tupla enviado de tipo tupla
                    int tam = t.size();     //tam su tamaño
                    if (tam != n)            //comprobación de que esta bien
                    {
                        cerr << "no son iguales la tupla mandada y la parte del mensaje que nos dice su tamaño" << endl;
                    }
                    if (tam == 1 || tam == 2 || tam == 3) {
                        // Enviamos mensaje sabiendo que está correcto
                        int send_bytes = soc2.Send(socket_fd1, buffer);
                        if (send_bytes == -1) {
                            cout << "error al enviar a servidores" << endl;
                            string mensError(strerror(errno));
                            cerr << "Error al enviar datos: " + mensError + "\n";
                            // Cerramos los sockets
                            soc2.Close(socket_fd1);
                            out = true;
                        }
                        cout << "Mensaje enviado a servidor 1 : " << buffer << endl;
                        rcv_bytes = soc2.Recv(socket_fd1, buffer, length);
                        if (rcv_bytes == -1) {
                            string mensError(strerror(errno));
                            cerr << "Error al recibir datos: " + mensError + "\n";
                            soc2.Close(socket_fd1);
                            out = true;
                        }
                    } else if (tam == 4 || tam == 5) {
                        // Enviamos mensaje sabiendo que está correcto
                        int send_bytes = soc3.Send(socket_fd2, buffer);
                        if (send_bytes == -1) {
                            cout << "error al enviar a servidores" << endl;
                            string mensError(strerror(errno));
                            cerr << "Error al enviar datos: " + mensError + "\n";
                            // Cerramos los sockets
                            soc3.Close(socket_fd2);
                            out = true;
                        }
                        cout << "Mensaje enviado a servidor 2 : " << buffer << endl;
                        rcv_bytes = soc3.Recv(socket_fd2, buffer, length);
                        if (rcv_bytes == -1) {
                            string mensError(strerror(errno));
                            cerr << "Error al recibir datos: " + mensError + "\n";
                            soc3.Close(socket_fd2);
                            out = true;
                        }
                    } else if (tam == 6) {
                        // Enviamos mensaje sabiendo que está correcto
                        int send_bytes = soc4.Send(socket_fd3, buffer);
                        if (send_bytes == -1) {
                            cout << "error al enviar a servidores" << endl;
                            string mensError(strerror(errno));
                            cerr << "Error al enviar datos: " + mensError + "\n";
                            // Cerramos los sockets
                            soc4.Close(socket_fd3);
                            out = true;
                        }
                        cout << "Mensaje enviado a servidor 3: " << buffer << endl;
                        rcv_bytes = soc4.Recv(socket_fd3, buffer, length);
                        if (rcv_bytes == -1) {
                            string mensError(strerror(errno));
                            cerr << "Error al recibir datos: " + mensError + "\n";
                            soc4.Close(socket_fd3);
                            out = true;
                        }
                    }
                    cout << "Mensaje recibido del servidor x: " << buffer << endl;

                    //reciba la tupla la envie al cliente
                    int send_bytes = soc.Send(client_fd, buffer);
                    if (send_bytes == -1) {
                        string mensError(strerror(errno));
                        cerr << "Error al enviar datos: " + mensError + "\n";
                        // Cerramos los sockets
                        soc.Close(client_fd);
                        out = true;
                    }


                    //actualiza el valor de las tuplas y si hay 0 y se tiene que cerrar el servidor echa al cliente
                    cLS.actualizaTuplas(inst, out);
                    //si linda esta vacio de tuplas y hay que acabar, entonces me salgo, aviso via monitor que hay que salir y despierto a los demás clientes

                }
            }
        }
    }
    cLS.saleCliente();
    soc.Close(client_fd);
    soc2.Close(socket_fd1);
    soc3.Close(socket_fd2);
    soc4.Close(socket_fd3);
}

int main(int argc,char* argv[]) {
    signal(SIGPIPE, SIG_IGN);
    SERVER_ADDRESS1 = argv[1];
    SERVER_ADDRESS2 = argv[2];
    SERVER_ADDRESS3 = argv[3];
    // Dirección y número donde escucha el proceso servidor
    int max_connections = 10;
    // Creación del socket con el que se llevará a cabo
    // la comunicación con el servidor.
    Socket socket1(SERVER_PORT);



    // Bind
    int socket_fd = socket1.Bind();
    if (socket_fd == -1) {
        string mensError(strerror(errno));
        cerr << "Error en el bind: " + mensError + "\n";
        exit(1);
    }

    // Listen
    int error_code = socket1.Listen(max_connections);
    if (error_code == -1) {
        string mensError(strerror(errno));
        cerr << "Error en el listen: " + mensError + "\n";
        // Cerramos el socket
        socket1.Close(socket_fd);
    }
    //bucle servidor
    int i = 1;
    bool salgo = false;
    cout << "ya escucho" << endl;

    while (!salgo) {
        // Accept
        int client_fd = socket1.Accept();
        //Durante el bloqueo del accept me han finalizado?
        cLS.acabo(salgo);

        if (client_fd == -1) {
            string mensError(strerror(errno));
            cerr << "Error en el accept: " + mensError + "\n";
            socket1.Close(socket_fd);
            exit(1);
        }
        if (!salgo) {//si tengo que salir, no creo un nuevo cliente
            cout << "Lanzo thread nuevo cliente " + to_string(i) + "\n";
            thread t(&servCliente, ref(socket1), client_fd);
            t.detach();
            cout << "Nuevo cliente " + to_string(i) + " aceptado" + "\n";
            i++;
        }
    }
    error_code = socket1.Close(socket_fd);
    if (error_code == -1) {
        string mensError(strerror(errno));
        cerr << "Error cerrando el socket del servidor: " + mensError + "\n";
    }
    //Comprobamos que todos los clientes han acabado y que no hay tuplas en los servidores antes de salir
    //cLS.entraCliente();
    // Despedida
    cout << "Bye bye" << endl;

}