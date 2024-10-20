//*****************************************************************
// File:   ControlLS.cpp
// Author: PSCD-Unizar
// Date:   noviembre 2018
// Coms:   Completar para la práctica 4
//*****************************************************************

#include "ControlLS.hpp"
#include <iostream>
#include <cassert>

//si DEBUG_MODE, generará eventos en log; si no, no hará nada
#if DEBUG_MODE 
    #define LOG(l,s) l->addMessage(s); //LOG(log, meNAaje): añade "meNAaje" evento en "log"
#else
    #define LOG(l,s)                   //LOG(log, meNAaje): no hace nada
#endif

//-----------------------------------------------------
ControlLS::ControlLS() {
	nProcess = 0;
	nTuplas = 0;
	fin = false;
	salir = false;
};


	//entra un cliente en el servidor, luego se suma a la cantidad de clientes que hay
    void ControlLS::entraCliente(){
		unique_lock<mutex> lck(mtx);
		
		if(fin && nProcess != 0)
		{//tiene que acabar pero tiene que esperar
			mainEspera.wait(lck);	
		}else if(fin && nProcess == 0)
		{//tiene que acabar y no hay a nadie al que esperar
		}else{
			nProcess = nProcess + 1;
		}
	}
	void ControlLS::actualizaTuplas(string inst, bool& out){
    	unique_lock<mutex> lck (mtx);

    	if(inst == "RN")
		{//despues de hacer RN ha borrado una tupla
    		nTuplas = nTuplas - 1;
		}else if(inst == "PN")
		{//despues de hacer PN ha añadido una tupla
			nTuplas = nTuplas + 1;
		}
    	if(nTuplas == 0 && fin)
        {
    	    out = true;
    	    salir = true;
        }
    }
    //sale un cliente del servidor, luego se resta cuantos hay
    void ControlLS::saleCliente(){
		unique_lock<mutex> lck(mtx);
		nProcess = nProcess - 1;
		if(nProcess == 0 && nTuplas == 0 && fin)
		{//bloqueado por entraCliente
			mainEspera.notify_one();
		}else if (fin){
			cout<<"Aun quedan por salir "<<nProcess<<" clientes y sigue habiendo "<<
			nTuplas<<" en la pizarra"<<endl;
		}
		
	}
	//cambia el estado del monitor a finalizacion
	void ControlLS::finalizacion() {
		unique_lock<mutex> lck(mtx);
		fin = true;
	}
	//devuelve por s el valor que tiene la variable del monitor salir
	void ControlLS::salida(bool& s){
		unique_lock<mutex> lck(mtx);
		s = salir;
    }
	//devuelve por f el valor que tiene la variable del monitor fin
	void ControlLS::acabo(bool& f){
		unique_lock<mutex> lck(mtx);
		f = fin;
	}
