//*****************************************************************
// File:   ControlAviones.hpp
// Author: PSCD-Unizar
// Date:   noviembre 2018
// Coms:   
//*****************************************************************

#ifndef CONTROL_LS
#define CONTROL_LS

#include <iostream>
#include <string>
#include <mutex>
#include <condition_variable>

using namespace std;

class ControlLS {
public:
    //----------------- Constructores
    ControlLS();

    //entra un cliente en el servidor, luego se suma a la cantidad de clientes que hay
    void entraCliente();
    //sale un cliente del servidor, luego se resta cuantos hay
    void saleCliente();
    //despues de una instruccion, actualiza el valor de las tuplas en la pizarra
    void actualizaTuplas(string inst, bool& out);
    //cambia el valor de la variable de finalizacion a verdadero
    void finalizacion();
    //devuelve por s el valor de salir
    void salida(bool& s);
    //devuelve por f el valor de fin
    void acabo(bool& f);

	
private:
    int nProcess;		//num de clientes
    int nTuplas;        //num de tuplas
    bool fin;			//hay que cerrar la entrada?
    bool salir;         //hay que sacar a los clientes?
    condition_variable_any mainEspera;	//bloquea al main
    mutex mtx;                    //los dos primeros atributos se entenderán al estudiar los monitores
};
#endif
