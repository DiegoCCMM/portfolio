cmake_minimum_required(VERSION 3.4)
project(Pruebas)

set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_FLAGS "-I${CMAKE_SOURCE_DIR} -pthread")

add_executable(clienteprobatinas Socket.cpp cliente_probatinas.cpp controlTuplas.cpp LindaDriver.cpp Tuplas.cpp)
add_executable(clienteprueba Socket.cpp clienteprueba.cpp controlTuplas.cpp LindaDriver.cpp Tuplas.cpp)
add_executable(clienteSimpe Socket.cpp clienteSimple.cpp controlTuplas.cpp LindaDriver.cpp Tuplas.cpp)
add_executable(clienteCarga clienteCarga.cpp LindaDriver.cpp Tuplas.cpp controlTuplas.cpp Socket.cpp)
add_executable(LindaServer Tuplas.cpp ControlLS.cpp LindaServer.cpp Socket.cpp)
add_executable(servidor Socket.cpp servidor.cpp controlTuplas.cpp Tuplas.cpp)
add_executable(admin Socket.cpp Tuplas.cpp admin.cpp)