#include "controlTuplas.hpp"
#include <cassert>
#include <iostream>

//-----------------------------------------------------
ControlTuplas::ControlTuplas() {
	vacio(a);	// Creamos el arbol vacio donde se guardarán las tuplas
	salir = false;
};

void ControlTuplas::anyadir(Tupla t){
	unique_lock<mutex> lck(mtx);
	insertar(a,t.to_string(),t);
	esperandoTupla.notify_all();
	// Se ha insertado la tupla
};

void ControlTuplas::eliminar(Tupla& t){
	unique_lock<mutex> lck(mtx);
	while(!esta(a,t) && !salir){
		// No se ha encontrado la tupla
		esperandoTupla.wait(lck); // Se bloquea
	}
	if(salir){
		Tupla x(t.size()); // Creamos una tupla completamente vacia
		t = x;	// Que sera la que enviemos al cliente
	}
	else{
		// Le pasamos la tupla t con sus valores actualizados
		borrar(a,t.to_string());
		// Se ha borrado correctamente la tupla
	}
};

void ControlTuplas::leer(Tupla& t){
	unique_lock<mutex> lck(mtx);
	while(!esta(a,t) && !salir){
		esperandoTupla.wait(lck);
	}
	if(salir){
		Tupla x(t.size()); // Creamos una tupla completamente vacia
		t = x;	// Que sera la que enviemos al cliente
	}
	// else: Los valores de t se han modificado
};

void ControlTuplas::finalizar(){
	unique_lock<mutex> lck(mtx);
	salir = true;
	esperandoTupla.notify_all();
}
void ControlTuplas::salgo(bool& s){
	unique_lock<mutex> lck(mtx);
	s = salir;
}
