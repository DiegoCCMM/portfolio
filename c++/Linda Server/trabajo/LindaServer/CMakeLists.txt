cmake_minimum_required(VERSION 3.4)
project(LindaServer)

set(CMAKE_CXX_STANDARD 11)

add_executable(LindaServer ls.cpp Tuplas.cpp Tuplas.hpp ControlLS.cpp ControlLS.hpp Socket.hpp Socket.cpp ls2.cpp ls2.hpp)
add_executable(clienteCarga clienteCarga.cpp Tuplas.hpp )