//
// Created by diego on 29/12/18.
//

#include "LindaDriver.hpp"
#include <iostream>
#include <string>
#include <cassert>
#include <thread>
#include "stdlib.h"



//si DEBUG_MODE, generará eventos en log; si no, no hará nada
#if DEBUG_MODE
#define LOG(l,s) l->addMessage(s); //LOG(log, meNAaje): añade "meNAaje" evento en "log"
#else
#define LOG(l,s)                   //LOG(log, meNAaje): no hace nada
#endif

//constructor del LindaDriver
LindaDriver::LindaDriver(char* IP, char* puerto):socket(IP,atoi(puerto)){
    int count;
    do {

        socket_fd = socket.Connect();
        count++;
        if(socket_fd == -1)
        {
            cout << "ERROR AL CONECTAR A LINDADRIVER"<<endl;
            this_thread::sleep_for(chrono::seconds(1));
        }
    }while(socket_fd == -1 && count < 10);

};

//postnote
void LindaDriver::PN(Tupla t){
    //creamos el mensaje para que lo pueda entender LindaServer
    string message = t.to_string();
    int tam = t.size();
    string stam = to_string(tam);
    message = "PN" + message;      //message = ACCION [TUPLA]
    message = message + stam;      //message = ACCION [TUPLA] TAM_TUPLA
    //Send
    int send_bytes = socket.Send(socket_fd,message);
    if(send_bytes == -1){
        cerr << "ERROR AL ENVIAR DATOS: " << strerror(errno) << endl;
        socket.Close(socket_fd);
    }
    cout<<"Enviado a LindaServer: "<<message<<endl;

    //tupla
    string buffer;

    //Recibimos la respuesta del servidor
    int read_bytes = socket.Recv(socket_fd, buffer, 4001);
    if(read_bytes == -1)
    {
        cerr<<"Error al recibir información"<<endl;
    }
    cout<<"Recibido de LindaServer: "<<buffer<<endl;
//recibo mensaje de confirmación pero en PN no hace nada

}

//removenote
void LindaDriver::RN(Tupla& t){
    //creamos el mensaje para que lo pueda entender LindaServer
    string message = t.to_string();
    int tam = t.size();
    string stam = to_string(tam);
    message = "RN" + message;      //message = ACCION [TUPLA]
    message = message + stam;      //message = ACCION [TUPLA] TAM_TUPLA
    //send
    int send_bytes = socket.Send(socket_fd,message);
    if(send_bytes == -1){
        cerr << "ERROR AL ENVUAR DATOS: " << strerror(errno) << endl;
        socket.Close(socket_fd);
    }
    //recibir
    cout<<"Enviado a LindaServer: "<<message<<endl;
    //tupla
    string buffer;

    //Recibimos la respuesta del servidor
    int read_bytes = socket.Recv(socket_fd, buffer, 4001);
    if(read_bytes == -1)
    {
        cerr<<"Error al recibir información"<<endl;
    }
    t.from_string(buffer);                  //tupla de salida es la recibida
    cout<<"Recibido de LindaServer: "<<buffer<<endl;
}

//readnote
void LindaDriver::ReadN(Tupla& t){
    //creamos el mensaje para que lo pueda entender LindaServer
    string message = t.to_string();
    int tam = t.size();
    string stam = to_string(tam);
    message = "ReadN" + message;      //message = ACCION [TUPLA]
    message = message + stam;      //message = ACCION [TUPLA] TAM_TUPLA
    int send_bytes = socket.Send(socket_fd,message);
    if(send_bytes == -1){
        cerr << "ERROR AL ENVIAR DATOS: " << strerror(errno) << endl;
        socket.Close(socket_fd);
    }
    //recibir
    cout<<"Enviado a LindaServer: "<<message<<endl;
    //tupla
    string buffer;

    //Recibimos la respuesta del servidor
    int read_bytes = socket.Recv(socket_fd, buffer, 4001);
    if(read_bytes == -1)
    {
        cerr<<"Error al recibir información"<<endl;
    }
    t.from_string(buffer);              //tupla de salida es la enviada por el servidor linda
    cout<<"Recibido de LindaServer: "<<buffer<<endl;
}

void LindaDriver::desconectarse(){
	
	//creamos el mensaje
	string message = "DESCONECTAR CLIENTE";
	
	int send_bytes = socket.Send(socket_fd,message);
    if(send_bytes == -1){
        cerr << "ERROR AL ENVIAR DATOS: " << strerror(errno) << endl;
        socket.Close(socket_fd);
    }
    //recibir
    cout<<"Enviado a LindaServer: "<<message<<endl;
}