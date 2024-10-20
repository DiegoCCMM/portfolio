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
#include <iomanip>
#include <time.h>
#include <stdio.h>
#include "Socket.hpp"
#include "Tuplas.hpp"
#include "LindaDriver.hpp"

using namespace std;

const int MESSAGE_SIZE = 4001; //mensajes de no más 4000 caracteres
int fin=0;

void pruebas(char* lip,char* lport){
	double m_escribir = 0.0, m_leer=0.0, m_eliminar=0.0;
	double t_escribir, t_leer, t_eliminar;
	LindaDriver LD(lip,lport);
	for(int i=0; i< 100; i++){
		srand(time(NULL));

		//Crear una tupla
		Tupla T1("b","a"); // servidor 1
		Tupla T2("a","a","b","b","b","a"); // servidor 3
		Tupla T3("b","a","b","a"); // servidor 2
		cout<<T1.to_string()<<endl;
		cout<<T2.to_string()<<endl;
		cout<<T3.to_string()<<endl;

		if(i%3==0){	// si i es divisible entre 3
			//----------
			clock_t start, end;
			start = clock();
			//enviarla
			LD.PN(T1);
			end = clock();
			t_escribir = ((end - start) / double(_SC_CLK_TCK))/1000.0;

			//----------
			start, end;
			start = clock();
			//eliminarla
			LD.ReadN(T1);
			end = clock();
			t_eliminar = ((end - start) / double(_SC_CLK_TCK))/1000.0;

			//----------
			start, end;
			start = clock();
			//leerla
			LD.RN(T1);
			end = clock();
			t_leer = ((end - start) / double(_SC_CLK_TCK))/1000.0;
		}
		else if(i%5){
			//----------
			clock_t start, end;
			start = clock();
			//enviarla
			LD.PN(T2);
			end = clock();
			t_escribir = ((end - start) / double(_SC_CLK_TCK))/1000.0;

			//----------
			start, end;
			start = clock();
			//eliminarla
			LD.ReadN(T2);
			end = clock();
			t_eliminar = ((end - start) / double(_SC_CLK_TCK))/1000.0;

			//----------
			start, end;
			start = clock();
			//leerla
			LD.RN(T2);
			end = clock();
			t_leer = ((end - start) / double(_SC_CLK_TCK))/1000.0;
		}
		else{
			//----------
			clock_t start, end;
			start = clock();
			//enviarla
			LD.PN(T3);
			end = clock();
			t_escribir = ((end - start) / double(_SC_CLK_TCK))/1000.0;

			//----------
			start, end;
			start = clock();
			//eliminarla
			LD.ReadN(T3);
			end = clock();
			t_eliminar = ((end - start) / double(_SC_CLK_TCK))/1000.0;
			
			//----------
			start, end;
			start = clock();
			//leerla
			LD.RN(T3);
			end = clock();
			t_leer = ((end - start) / double(_SC_CLK_TCK))/1000.0;


		}

		cout << "INSERTAR: " << t_escribir << endl;
		cout << "LEER: " << t_leer << endl;
		cout << "ELIMINAR" << t_eliminar << endl;

		m_escribir += t_escribir;
		m_leer += t_leer;
		m_eliminar += t_eliminar;
	}
	// se han terminado de realizar las 100 pruebas
	cout << "*********Tiempo de ejecución medio**********" << endl;
	cout << "INSERTAR: " << fixed << setprecision(3) << m_escribir/100.0 << " segundos" <<endl;
	cout << "LEER: " << fixed << setprecision(3) << m_leer/100.0 << " segundos" <<endl;
	cout << "ELIMINAR" << fixed << setprecision(3) << m_eliminar/100.0 << " segundos" <<endl;
	fin=1;
}

int main(int argc,char* argv[]) {
	if(argc < 3){
    	cout << "argumentos incorrectos";
    }
	else{
		int K=100;
		// Dirección y número donde escucha el proceso servidor
		char* LINDA_ADDRESS = argv[1];
		char* LINDA_PORT = argv[2];
		thread P[K];
		for(int i=0; i<K; i++){
			P[i]=thread(&pruebas,LINDA_ADDRESS,LINDA_PORT);
			P[i].detach();
		}
	}
	while(fin==0){
	}
	return(0);
}
