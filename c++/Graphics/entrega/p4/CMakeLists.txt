cmake_minimum_required(VERSION 3.17)
project(p4)

set(CMAKE_CXX_STANDARD 20)

add_executable(p4 main.cpp Sphere.hpp Plane.hpp geometryRGBFigures.hpp Material.hpp)

add_executable(p3 mainP3.cpp Sphere.hpp Plane.hpp geometryRGBFigures.hpp)