//
// Created by diego on 29/12/18.
//

#ifndef LINDADRIVER
#define LINDADRIVER

#include <iostream>
#include <string>
#include <mutex>
#include <condition_variable>
#include "Socket.hpp"
#include "Tuplas.hpp"

using namespace std;

class LindaDriver {
public:
    //constructor del LindaDriver
    LindaDriver(char* IP, char* puerto);

    //postnote
    void PN(Tupla t);

    //removenote
    Tupla RN(Tupla t);

    //readnote
    Tupla ReadN(Tupla t);


private:
    Socket socket;
    mutex mtx;
};
#endif //LINDADRIVER_LINDADRIVER_HPP
