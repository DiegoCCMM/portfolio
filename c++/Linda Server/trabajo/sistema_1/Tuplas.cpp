#include "Tuplas.hpp"
#include <stdio.h>
#include <string.h>
#include <iostream>
#include <sstream>

Tupla::Tupla() {
}

Tupla::Tupla(int n) {
    for (int i = 0; i < n; i++) {
        cadena.emplace_back("");
    }
};

Tupla::Tupla(string n1) {
    cadena.emplace_back(n1);
};

Tupla::Tupla(string n1, string n2) {
    cadena.emplace_back(n1);
    cadena.emplace_back(n2);
};

Tupla::Tupla(string n1, string n2, string n3) {
    cadena.emplace_back(n1);
    cadena.emplace_back(n2);
    cadena.emplace_back(n3);
};

Tupla::Tupla(string n1, string n2, string n3, string n4) {
    cadena.emplace_back(n1);
    cadena.emplace_back(n2);
    cadena.emplace_back(n3);
    cadena.emplace_back(n4);
};

Tupla::Tupla(string n1, string n2, string n3, string n4, string n5) {
    cadena.emplace_back(n1);
    cadena.emplace_back(n2);
    cadena.emplace_back(n3);
    cadena.emplace_back(n4);
    cadena.emplace_back(n5);
};

Tupla::Tupla(string n1, string n2, string n3, string n4, string n5, string n6) {
    cadena.emplace_back(n1);
    cadena.emplace_back(n2);
    cadena.emplace_back(n3);
    cadena.emplace_back(n4);
    cadena.emplace_back(n5);
    cadena.emplace_back(n6);
};

void Tupla::set(int n, string label) {
    cadena[n - 1] = label;
};

string Tupla::get(int n) {
    return cadena[n - 1];
};

string Tupla::to_string() {
    std::stringstream ss;
    ss << "[";
    for (int i = 0; i < size(); i++) {
        ss << this->cadena[i] << ",";
    }
    string label = ss.str();
    size_t tam = label.length();
    label.erase(tam - 1);
    label = label + "]";
    return label;
};

void Tupla::from_string(string label) {
    label.erase(0, 1);
    size_t tam = label.length();
    cadena.clear();
    label.erase(tam - 1);//Ya no hay []. Hay que separarlo por comas.
    string patron = ","; //Patron a buscar.
    int posInit = 0;
    int posFound = 0;
    string splitted;
    int i = 0;
    while (posFound >= 0) {
        posFound = label.find(patron, posInit); //Encuentra la posicion de la primera coma.
        splitted = label.substr(posInit, posFound - posInit); //Se carga la la coma.
        posInit = posFound + 1;
        cadena.emplace_back(splitted); //Mete en la tupla el elemento sin coma.
        i++;
    }
};

int Tupla::size() {
    return cadena.size();
};

bool Tupla::esVariable(int n) {
    string x = this->cadena[n - 1];
    return x.length() == 2 && x.substr(0, 1) == "?" && x.substr(1) >= "A" && x.substr(1) <= "Z";
};

int Tupla::letra(int n) {
    string x = this->cadena[n - 1];
    char letra = x[1];    // x[1] = letra
    return letra - 'A';
};