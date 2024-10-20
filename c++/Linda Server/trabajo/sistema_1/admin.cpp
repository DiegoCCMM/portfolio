#include <iostream>
#include "Socket.hpp"
#include "Tuplas.hpp"
#include <string>
#include <thread>

//funcion de administrador de finalización
int main(int argc, char* argv[]){
    //enviar
    int socket_fd;
    //pasar primero la IP y luego el puerto
    Socket socket(argv[1], atoi(argv[2]));
    int count;
    do {
        //Connect
        socket_fd = socket.Connect();
        count++;
        if(socket_fd == -1)
        {
            cout << "ERROR AL ENVIAR A LINDADRIVER"<<endl;
            this_thread::sleep_for(chrono::seconds(1));
        }
    }while(socket_fd == -1 && count < 10);
    string message = "FIN_SERVER";
    //Send
    int send_bytes = socket.Send(socket_fd,message);
    if(send_bytes == -1){
        cerr << "ERROR AL ENVIAR DATOS: " << strerror(errno) << endl;
        socket.Close(socket_fd);
    }
    cout<<"Enviado a LindaServer: "<<message<<endl;

    //respuesta LindaServer
    string buffer;

    //Recibimos la respuesta del LS
    int read_bytes = socket.Recv(socket_fd, buffer, 4001);
    if(read_bytes == -1)
    {
        cerr<<"Error al recibir información"<<endl;
    }
    cout<<buffer<<endl;

}