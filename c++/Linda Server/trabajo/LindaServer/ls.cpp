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

const string SERVER_ADDRESS1 = "0.0.0.0";  //el que se ocupe del tamaño 1-3
const string SERVER_ADDRESS2 = "0.0.0.0";  //el que se ocupe del tamaño 4-5
const string SERVER_ADDRESS3 = "0.0.0.0";  //el que se ocupe del tamaño 6
const int SERVER_PORT = 5000;      //el puerto que se conectan los clientes (ls se comporta como servidor)
const int MEM_PORT1 = 5001;      //el puerto en el que recibe de los servidores (ls se comporta como cliente)
const int MEM_PORT2 = 5002;
const int MEM_PORT3 = 5003;
ControlLS cLS;

//soc = client_port, soc2 = server_port
void servCliente(Socket& soc, int client_fd) {
    cLS.entraCliente();
    // Buffer para recibir el mensaje
    int length = 100;
    char buffer[length];
    Socket soc2(SERVER_ADDRESS1, MEM_PORT1);
    int socket_fd;
    bool out = false; // Inicialmente no salir del bucle
    while(!out) {
        // Recibimos el mensaje del cliente
        int rcv_bytes = soc.Recv(client_fd,buffer,length);
        if (rcv_bytes == -1) {
            string mensError(strerror(errno));
            cerr << "Error al recibir datos: " + mensError + "\n";
            // Cerramos los sockets
            soc.Close(client_fd);
            return;
        }
        //antes de procesar el mensaje, igual teniamos que salir y estaba bloqueado
        cLS.salida(out);
        //si hay que salir, me salgo, sino proceso el mensaje
        if(!out) {
            cout << "Mensaje recibido: " << buffer << endl;
            //si el mensaje recibido es el de finalización de servidor o el de sacar a todos los clientes, lo procesamos
            if(strncmp(buffer, "FIN_SERVER",10) == 0)
            {
                cout << "\nESPERANDO A QUE FINALICEN LOS CLIENTES" << endl;
                cout << "SE CIERRA LA ENTRADA" << endl;

                    cout<<"\n Se ha cerrado el servidor, esperando a los clientes"<<endl;
                    cLS.finalizacion();
                    //desbloquear del accept mediante un hijo o proceso finalizador
                    Socket socket("0.0.0.0", SERVER_PORT);
                    int socket_fd = socket.Connect();
                    socket.Close(socket_fd);
                    out = true;
                    string mens="YA HABEMO ACABAO";
                    int fd;
                    soc.Send(fd,mens);

            }else{//si el mensaje recibido no es el de finalización, entonces procesamos el mensaje de LD
                string sbuffer;
                sbuffer.assign(buffer, buffer + length);
                //formato del mensaje recibido : ACCION [TUPLA] TAM_TUPLA
                string sn = sbuffer.substr(sbuffer.find_first_of("]"),
                                           sbuffer.find_first_of(","));   //n tamaño de la tupla
                sn.pop_back();      //elimina ","
                sn.erase(sn.begin());     //elimina "]"
                int n = stoi(sn);
                string inst = sbuffer.substr(0, sbuffer.find_first_of("["));   //instruccion de la tupla
                Tupla t(n);
                string info = sbuffer.substr(sbuffer.find_first_of("["),
                                             sbuffer.find_first_of("]"));   //info tupla en si
                t.from_string(info);    //t es la tupla enviado de tipo tupla
                int tam = t.size();     //tam su tamaño
                if (tam != n)            //comprobación de que esta bien
                {
                    cerr << "no son iguales la tupla mandada y la parte del mensaje que nos dice su tamaño" << endl;
                }

                if (tam == 1 || tam == 2 || tam == 3) {
                    Socket soc2(SERVER_ADDRESS1, MEM_PORT1);//rellenar con ip y puerto correspondientes al s1
                    cout << "Va al server 1" << endl;
                } else if (tam == 4 || tam == 5) {
                    Socket soc2(SERVER_ADDRESS2, MEM_PORT2);//rellenar con ip y puerto correspondientes al s2
                    cout << "Va al server 2" << endl;
                } else if (tam == 6) {
                    Socket soc2(SERVER_ADDRESS3, MEM_PORT3);//rellenar con ip y puerto correspondientes al s3
                    cout << "Va al server 3" << endl;
                }

                // Conectamos con el servidor. Probamos varias conexiones
                const int MAX_ATTEMPS = 10;
                int count = 0;
                do {
                    // Conexión con el servidor
                    socket_fd = soc2.Connect();
                    count++;
                    cout << "PostConnect" << endl;
                    // Si error --> esperamos 1 segundo para reconectar
                    if (socket_fd == -1) {
                        cout << "error al conectar" << endl;
                        this_thread::sleep_for(chrono::seconds(1));
                    }
                } while ((socket_fd == -1) && count < MAX_ATTEMPS);

                // Chequeamos si se ha realizado la conexión
                if (socket_fd == -1) {
                    cerr << "error en el connect" << endl;
                }

                // Enviamos mensaje sabiendo que está correcto
                int send_bytes = soc2.Send(socket_fd, buffer);
                if (send_bytes == -1) {
                    cout << "error al enviar a servidores" << endl;
                    string mensError(strerror(errno));
                    cerr << "Error al enviar datos: " + mensError + "\n";
                    // Cerramos los sockets
                    soc2.Close(socket_fd);
                    out = true;
                }
                cout << "Mensaje enviado a servidor x: " << buffer << endl;
                rcv_bytes = soc2.Recv(socket_fd, buffer, length);
                if (rcv_bytes == -1) {
                    string mensError(strerror(errno));
                    cerr << "Error al recibir datos: " + mensError + "\n";
                    soc2.Close(socket_fd);
                    out = true;
                }
                cout << "Mensaje recibido del servidor x: " << buffer << endl;
                if (inst == "PN") {
                    //envio el mensaje de confirmacion tal cual a Linda Driver
                    int send_bytes = soc.Send(client_fd, buffer);
                    if (send_bytes == -1) {
                        string mensError(strerror(errno));
                        cerr << "Error al enviar datos: " + mensError + "\n";
                        // Cerramos los sockets
                        soc.Close(socket_fd);
                        out = true;
                    }
                } else if (inst == "RN") {
                    //recibe tupla, tiene que reenviarla al cliente
                    int send_bytes = soc.Send(client_fd, buffer);
                    if (send_bytes == -1) {
                        string mensError(strerror(errno));
                        cerr << "Error al enviar datos: " + mensError + "\n";
                        // Cerramos los sockets
                        soc.Close(socket_fd);
                        out = true;
                    }
                } else if (inst == "ReadN") {
                    //reciba la tupla la envie al cliente
                    int send_bytes = soc.Send(client_fd, buffer);
                    if (send_bytes == -1) {
                        string mensError(strerror(errno));
                        cerr << "Error al enviar datos: " + mensError + "\n";
                        // Cerramos los sockets
                        soc.Close(socket_fd);
                        out = true;
                    }
                }

                //actualiza el valor de las tuplas y si hay 0 y se tiene que cerrar el servidor echa al cliente
                cLS.actualizaTuplas(inst, out);
                //si linda esta vacio de tuplas y hay que acabar, entonces me salgo, aviso via monitor que hay que salir y despierto a los demás clientes
                if (out) {
                    Socket salida("0.0.0.0", SERVER_PORT);
                    salida.Connect();
                    soc.Accept();
                    int send_bytes = salida.Send(socket_fd, "ey");
                    //todo igual necesitamos que todos los clientes al salir vayan desbloqueando ya que 1 Send desbloquea 1.
                    //todo es decir que se vayan desbloqueando "recursivamente"
                }
            }
        }
    soc2.Close(socket_fd);
    }cLS.saleCliente();
    soc.Close(client_fd);
}

int main(){
    signal(SIGPIPE,SIG_IGN);
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
    cout<<"ya escuxho"<<endl;

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
        if(!salgo) {//si tengo que salir, no creo un nuevo cliente
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
    cLS.entraCliente();
    // Despedida
    cout << "Bye bye" << endl;

}