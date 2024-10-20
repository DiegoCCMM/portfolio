#ifndef _TUPLAS_H
#define _TUPLAS_H

#include <iostream>
#include <string>

using namespace std;

class Tupla {
    public:
        //funciones publicas;
        //Constructores (Una para cada tamaño de tupla)
        Tupla(const int n); //Por ejemplo, n = 2 -> la tupla seria ("","")
        Tupla(const string n1);
        Tupla(const string n1, string n2);
        Tupla(const string n1, string n2, string n3);
        Tupla(const string n1, string n2, string n3, string n4);
        Tupla(const string n1, string n2, string n3, string n4, string n5);
        Tupla(const string n1, string n2, string n3, string n4, string n5, string n6);
        //Deconstructor
        ~Tupla();

        //Establece "label" en el apartado n de la tupla.
        void set (const int n, const string label);

        //Devuelve el string correspondiente al apartado n de la tupla.
        string get (const int n);

        //Devuelve todas las tuplas en forma de un solo string
        string to_string ();

        //Modifica el contenido de la tupla para que quede como en label.
        //PARCIAL: label tiene que contener el mismo numero de elementos que la tupla.
        void from_string(string label);

        //Devuelve el numero de elementos de la tupla.
        int size();


    private:
        int tamanyo;
        string* cadena;
};


#endif