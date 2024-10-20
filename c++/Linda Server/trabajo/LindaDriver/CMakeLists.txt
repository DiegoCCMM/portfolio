cmake_minimum_required(VERSION 3.13)
project(LindaDriver)

set(CMAKE_CXX_STANDARD 11)

add_executable(LindaDriver LindaDriver.hpp LindaDriver.cpp admin.cpp)