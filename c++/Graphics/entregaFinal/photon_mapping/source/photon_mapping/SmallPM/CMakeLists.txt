cmake_minimum_required(VERSION 3.21)
project(photon_mapping)

set(CMAKE_CXX_STANDARD 20)

if (NOT CMAKE_BUILD_TYPE)
    set(CMAKE_BUILD_TYPE Release)
endif ()

set(CMAKE_CXX_FLAGS_DEBUG "-g")
set(CMAKE_CXX_FLAGS_RELEASE "-O3")

include(CheckPIESupported)
check_pie_supported()

add_executable(smallpm src/main.cpp src/PhotonMapping.cpp src/RenderEngine.cpp)
target_include_directories(smallpm PRIVATE include smallrt/include)
target_link_libraries(smallpm PRIVATE ${CMAKE_SOURCE_DIR}/smallrt/lib/smallrt.a)
set_property(TARGET smallpm PROPERTY POSITION_INDEPENDENT_CODE FALSE)