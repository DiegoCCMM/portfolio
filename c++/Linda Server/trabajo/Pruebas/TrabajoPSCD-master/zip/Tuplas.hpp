#ifndef _TUPLAS_H
#define _TUPLAS_H

#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Tupla {
    public:
        //funciones publicas;
        //Constructores (Una para cada tamaño de tupla)
		Tupla();
        Tupla(int n); //Por ejemplo, n = 2 -> la tupla seria ("","")
        Tupla(string n1);
        Tupla(string n1, string n2);
        Tupla(string n1, string n2, string n3);
        Tupla(string n1, string n2, string n3, string n4);
        Tupla(string n1, string n2, string n3, string n4, string n5);
        Tupla(string n1, string n2, string n3, string n4, string n5, string n6);

        //Establece "label" en el apartado n de la tupla.
        void set (int n, string label);

        //Devuelve el string correspondiente al apartado n de la tupla.
        string get (int n);

        //Devuelve todas las tuplas en forma de un solo string
        string to_string ();

        //Modifica el contenido de la tupla para que quede como en label.
        //PARCIAL: label tiene que contener el mismo numero de elementos que la tupla.
        void from_string(string label);

        //Devuelve el numero de elementos de la tupla.
        int size();
		
		// Devuelve true si el elemento de la tupla indicado es una variable
		bool esVariable (int n);
		
		// Devuelve la posicion de una letra en el abecedario
		int letra(int n);

    private:
        int tamanyo;
        std::vector<string> cadena;
};


#endif