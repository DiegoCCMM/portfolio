#ifndef _LISTA_H_
#define _LISTA_H_

#include <string>
#include <iostream>
#include <stdio.h>

using namespace std;

/* Modulo de implementacion de listas encadenadas en memoria dinamica organizada
  en pares (clave, tupla) con la repeticon admitida*/
template <typename Elemento> struct lista;

//Devuelve la lista vacia
template <typename Elemento> void vacio (lista<Elemento>& a);

//Devuelve true si a es la lista vacía, falso en caso contrario.
template <typename Elemento> bool esVacio (lista<Elemento>& a);

//Devuelve la lista a resultante de añadir el par (c,v) a la lista a. 
template <typename Elemento> void insertar (lista<Elemento>& a, string clave, Elemento& e);

//Dada una clave c, devuelve la lista a resultante de eliminar de la lista a el nodo que contenia dicha clave.
template <typename Elemento> void borrar (lista<Elemento>& a, const string clave);

// Devuelve true si el elemento e existe en memoria
template <typename Elemento> bool esta (lista<Elemento>& a, Elemento& e);

template <typename Elemento> bool esVariable (const string x);
template <typename Elemento> string card (lista<Elemento>& a);

template <typename Elemento> struct lista{
    friend void vacio<Elemento> (lista<Elemento>& a);
    friend bool esVacio<Elemento> (lista<Elemento>& a);
    friend void insertar<Elemento> (lista<Elemento>& a, string clave, Elemento& e);
    friend void borrar<Elemento> (lista<Elemento>& a, const string clave);
	friend bool esta<Elemento> (lista<Elemento>& a, Elemento& e);
    //friend string card<Elemento> (lista<Elemento>& a);
    friend bool esVariable<Elemento> (const string x);
    private:
        struct Nodo {
            string clave;
            Elemento dato;
            Nodo* siguiente;
            Nodo* anterior;
        };
        Nodo* primero;
        Nodo* ultimo;
        int cardinal;
};

template <typename Elemento> bool esVariable (const string x){
    return x.length() == 2 && x.substr(0,1) == "?" && x.substr(1) >= "A" && x.substr(1) <= "Z";
}

template <typename Elemento> void vacio (lista<Elemento>& a){
    a.primero = nullptr;
    a.ultimo = nullptr;
    a.cardinal = 0;
}

template <typename Elemento> bool esVacio (lista<Elemento>& a){
    return a.cardinal==0;
}

template <typename Elemento> void insertar (lista<Elemento>& a, string clave, Elemento& e){
    typename lista<Elemento>::Nodo* aux = new typename lista<Elemento>::Nodo;
    aux -> clave = clave;
    aux -> dato = e;
    aux -> siguiente = a.primero;
    aux -> anterior = nullptr;
    if (a.primero == nullptr){
        a.ultimo = aux;
    }
    else{
        a.primero -> anterior = aux;
    }
    a.primero = aux;
    a.cardinal = a.cardinal + 1;
}

template <typename Elemento> void borrar (lista<Elemento>& a, const string clave){
    typename lista<Elemento>::Nodo* aux = nullptr;
    aux = a.primero;
    bool encontrado = false;
    while ((aux != nullptr) && (!encontrado)){
        if(aux -> clave == clave){
            encontrado = true;
        }
        else{
            aux = aux -> siguiente;
        }
    }
    if (encontrado){
        if (a.primero == aux && a.ultimo == aux){
            a.primero = nullptr;
            a.ultimo = nullptr;
        }
        else if(a.primero == aux){
            a.primero = aux ->siguiente;
            aux -> siguiente -> anterior = nullptr;
        }
        else if (a.ultimo == aux){
            a.ultimo = aux -> anterior;
            aux -> anterior -> siguiente = nullptr;
        }
        else{
            aux -> anterior -> siguiente = aux -> siguiente;
            aux -> siguiente -> anterior = aux -> anterior;
        }
        delete aux;
        a.cardinal --;
    }
}

template <typename Elemento> bool esta (lista<Elemento>& a, Elemento& e){
    bool encontrado = false;
    if (a.cardinal == 0){
        encontrado = false;
    }
    else{
        typename lista<Elemento>::Nodo* aux = nullptr;
        aux = a.primero;
		string abc[26];
        while (aux != nullptr && !encontrado){
			for(int i=0; i<26; i++){
				abc[i] = "";
			}
			
            bool esigual = true; // En principio consideramos que las tuplas son iguales
		    for(int j=1; (j<=e.size()) && esigual; j++){
			    // Comprueba si todos los elementos de ambas tuplas son iguales
			    if(aux -> dato.get(j) != e.get(j) && !esVariable<Elemento>(e.get(j))){
				    esigual = false;	// las tuplas no son iguales
			    }
				else if(e.esVariable(j)){
					if(abc[e.letra(j)] == ""){
						abc[e.letra(j)] = aux -> dato.get(j);
					}
					else{
						esigual = (abc[e.letra(j)] == aux -> dato.get(j));
					}
				}
            }
            
            // Si la tupla no se ha encontrado en memoria
            if(!esigual){
                aux = aux -> siguiente;
            }
            else {
                e = aux -> dato;
                encontrado = true;
            }
        }
	}
    return encontrado;
}

#endif