//*****************************************************************
// File:   clienteCarga.cpp
// Author: PSCD-Unizar
// Date:   enero 2018
// Coms:   TP6 curso 18-19. Cliente Linda de carga
//         invocar mediante 
//            clienteCarga numTuplas IP_LindaServer Port_LindaServer ...
//*****************************************************************
#include <iostream>
#include <cstdlib>
#include <string>

#include "LindaDriver.hpp"
#include "Tuplas.hpp"

using namespace std;
//-----------------------------------------------------
void errorInvocacion() {
    cerr << "Invocar como:" << endl;
    cerr << "   clienteCarga <NT> <IP_LS> <Port_LS> ..." << endl;
    cerr << "      <NT>: num de tuplas" << endl;
    cerr << "      <IP_LS>: IP del servidor Linda" << endl;
    cerr << "      <Port_LS>: puerto del servidor Linda" << endl;
}
//-----------------------------------------------------
int main(int argc, char* argv[]) {
    const int MAX_LENGTH = 6; //máx longitud tuplas
    //faltan argumentos en la invocación?
    if (argc < 4 ) {
        errorInvocacion();
        return 1;
    }
    int NT = atoi(argv[1]); //num de tuplas a tratar
    // un driver con los datos pasados en la invocación
    LindaDriver LD(argv[2], argv[3]);
    // La conexión con el servidor Linda ya está establecido
    //-----------------------------------------------------
    //MAX_LENGTH tuplas para ir rellenando y enviando a Linda
    //Se crean con cadenas vacías
    Tupla *tuplas[MAX_LENGTH];
    for (int j=0; j<MAX_LENGTH; j++) {
        tuplas[j] = new Tupla(j+1);
    }
    cout << "Tuplas creadas" << endl;
    //-----------------------------------------------------
    Tupla *patrones[MAX_LENGTH]; //van a ser constantes
    patrones[0] = new Tupla("?A");
    patrones[1] = new Tupla("?A","?B");
    patrones[2] = new Tupla("?A","?B","?C");
    patrones[3] = new Tupla("?A","?B","?C","?D");
    patrones[4] = new Tupla("?A","?B","?C","?D","?E");
    patrones[5] = new Tupla("?A","?B","?C","?D","?E","?F");
    cout << "Patrones creados" << endl;
    //-----------------------------------------------------
    // Insertar 6*NT tuplas
    for (int i=1; i<=NT; i++) {
        string vT[MAX_LENGTH];
        vT[0] = to_string(i);
        for (int j=1; j<MAX_LENGTH; j++) {
            vT[j] = vT[j-1] + "," + to_string(i+j);
        }

        for (int j=0; j<MAX_LENGTH; j++) {
            tuplas[j]->from_string('[' + vT[j] + ']');
            // PosNote
            LD.PN(*tuplas[j]);
        }
    }
    cout << "Tuplas insertadas" << endl;
    // Retirar 6*NT tuplas
    for (int i=1; i<=NT; i++) {
        for (int j=0; j<MAX_LENGTH; j++) {
            // RemoveNote
            *tuplas[j] = LD.RN(*patrones[j]); //misma longitud
        }
        const int GRUPO = 50;
        // mostrar primera y última tuplas de cada grupo de GRUPO tuplas rescatadas
        if ((i-1)%GRUPO == 0) {
            cout << "    " << tuplas[0]->to_string() << endl;
            cout << "    " << tuplas[MAX_LENGTH-1]->to_string() << endl;
        }
    }
    cout << "Tuplas retiradas" << endl;
    //-----------------------------------------------------
    for (int i=0; i<MAX_LENGTH; i++) {
    	delete patrones[i];
    	delete tuplas[i];
    }
    //-----------------------------------------------------

    return 0;
}