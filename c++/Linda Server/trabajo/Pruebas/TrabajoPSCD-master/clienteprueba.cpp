//*****************************************************************
// File:   Cliente_p5.cpp
// Author: Sergio García Esteban 755844
// Date:   DICIEMBRE 2018
// Coms:   Práctica 5
//
//*****************************************************************

#include <iostream>
#include <chrono>
#include <thread>
#include "Socket.hpp"
#include "Tuplas.hpp"
#include "LindaDriver.hpp"

using namespace std;

const int MESSAGE_SIZE = 4001; //mensajes de no más 4000 caracteres

int main(int argc,char* argv[]) {
	if(argc < 3){
    	cout << "argumentos incorrectos";
    }
	else{
		
		// Dirección y número donde escucha el proceso servidor
		char* LINDA_ADDRESS = argv[1];
		char* LINDA_PORT = argv[2];

		//Inicializar LindaDriver
		LindaDriver LD(LINDA_ADDRESS,LINDA_PORT);//aqui mirar parametros lindadriver
		while(1){
			cout << "Accion a realizar (PN,RN,ReadN)(1,2,3): ";
			int accion;
			cin >> accion;
			cout << "Numero de elementos que desea introducir en la tupla: ";
			int n;
			cin >> n;
			string texto;
			cout << "Escriba la tupla en formato -> [elemento,elemento,...]: ";
			cin >> texto;
			//Crear una tupla
			Tupla T(n);
			T.from_string(texto);
			//enviarla
			if(accion == 1){
				LD.PN(T);
			}
			else if(accion == 2){
				LD.RN(T);
			}
			else{
				LD.ReadN(T);
			}
		}

	}
}
