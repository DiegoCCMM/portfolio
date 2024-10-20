#ifndef CONTROL_TUPLAS
#define CONTROL_TUPLAS

#include "lista.h"
#include "Tuplas.hpp"
#include <iostream>
#include <string>
#include <mutex>
#include <condition_variable>

using namespace std;

class ControlTuplas {
public:
	//----------------- Constructor
	ControlTuplas();

	// Añade la tupla t al arbol a
	void anyadir(Tupla t, bool& error);

	// Elimina del avl a una tupla t y actualiza sus valores
	void eliminar(Tupla& t);

	// Lee del avl a una tupla t y actualiza sus valores
	void leer(Tupla& t);

	// Cambia el estado de la variable salir a verdadero y envia una tupla vacia a los clientes
	void finalizar();

	void salgo(bool& s);

private:
	lista<Tupla> a;	// Arbol AVL donde se almacenan las tuplas
	bool salir;
	mutex mtx;
	condition_variable esperandoTupla;
};
#endif
