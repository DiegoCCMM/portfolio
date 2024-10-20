#include "Tuplas.hpp"
#include <stdio.h>
#include <string.h>
#include <iostream>

Tupla::Tupla(const int n){
    this -> tamanyo = n;
    this -> cadena = new string[n];
    for (int i = 0; i < n; i++) {
        this -> cadena[i] = "";
    }
};

Tupla::Tupla(const string n1){
    this -> tamanyo = 1;
    this -> cadena = new string[1];
    this -> cadena[0] = n1;
};

Tupla::Tupla(const string n1, string n2){
    this -> tamanyo = 2;
    this -> cadena = new string[2];
    this -> cadena[0] = n1;
    this -> cadena[1] = n2;
};

Tupla::Tupla(const string n1, string n2, string n3){
    this -> tamanyo = 3;
    this -> cadena = new string[3];
    this -> cadena[0] = n1;
    this -> cadena[1] = n2;
    this -> cadena[2] = n3;
};

Tupla::Tupla(const string n1, string n2, string n3, string n4){
    this -> tamanyo = 4;
    this -> cadena = new string[4];
    this -> cadena[0] = n1;
    this -> cadena[1] = n2;
    this -> cadena[2] = n3;
    this -> cadena[3] = n4;
};

Tupla::Tupla(const string n1, string n2, string n3, string n4, string n5){
    this -> tamanyo = 5;
    this -> cadena = new string[5];
    this -> cadena[0] = n1;
    this -> cadena[1] = n2;
    this -> cadena[2] = n3;
    this -> cadena[3] = n4;
    this -> cadena[4] = n5;
};

Tupla::Tupla(const string n1, string n2, string n3, string n4, string n5, string n6){
    this -> tamanyo = 6;
    this -> cadena = new string[6];
    this -> cadena[0] = n1;
    this -> cadena[1] = n2;
    this -> cadena[2] = n3;
    this -> cadena[3] = n4;
    this -> cadena[4] = n5;
    this -> cadena[5] = n6;
};

void Tupla::set (const int n, const string label){
    this -> cadena[n-1] = label;
};

string Tupla::get (const int n){
    return  this -> cadena[n-1];
};

string Tupla::to_string (){
    string label = "[";
    for (int i = 0; i < this->tamanyo ; i++){
        label = label + this->cadena[i]+",";
    }
    size_t tam = label.length();
    label.erase(tam-1);
    label = label +"]";
    return label;
};

void Tupla::from_string(string label){
    label.erase(0,1);
    size_t tam = label.length();
    label.erase(tam-1);//Ya no hay []. Hay que separarlo por comas.
    string patron = ","; //Patron a buscar.
    int posInit = 0;
    int posFound = 0;
    string splitted;
    int i = 0;
    while (posFound >= 0){
        posFound = label.find(patron, posInit); //Encuentra la posicion de la primera coma.
        splitted = label.substr(posInit, posFound-posInit); //Se carga la la coma.
        posInit = posFound +1;
        this -> cadena[i] = splitted; //Mete en la tupla el elemento sin coma.
        i++;
    }
};

int Tupla::size(){
    return this->tamanyo;
};

Tupla::~Tupla(){
    delete(cadena);
}