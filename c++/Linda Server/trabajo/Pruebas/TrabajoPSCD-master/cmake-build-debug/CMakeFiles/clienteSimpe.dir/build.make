# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.13

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /opt/clion/bin/cmake/linux/bin/cmake

# The command to remove a file.
RM = /opt/clion/bin/cmake/linux/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug

# Include any dependencies generated for this target.
include CMakeFiles/clienteSimpe.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/clienteSimpe.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/clienteSimpe.dir/flags.make

CMakeFiles/clienteSimpe.dir/Socket.cpp.o: CMakeFiles/clienteSimpe.dir/flags.make
CMakeFiles/clienteSimpe.dir/Socket.cpp.o: ../Socket.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/clienteSimpe.dir/Socket.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/clienteSimpe.dir/Socket.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Socket.cpp

CMakeFiles/clienteSimpe.dir/Socket.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/clienteSimpe.dir/Socket.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Socket.cpp > CMakeFiles/clienteSimpe.dir/Socket.cpp.i

CMakeFiles/clienteSimpe.dir/Socket.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/clienteSimpe.dir/Socket.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Socket.cpp -o CMakeFiles/clienteSimpe.dir/Socket.cpp.s

CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.o: CMakeFiles/clienteSimpe.dir/flags.make
CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.o: ../clienteSimple.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/clienteSimple.cpp

CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/clienteSimple.cpp > CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.i

CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/clienteSimple.cpp -o CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.s

CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.o: CMakeFiles/clienteSimpe.dir/flags.make
CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.o: ../controlTuplas.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/controlTuplas.cpp

CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/controlTuplas.cpp > CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.i

CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/controlTuplas.cpp -o CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.s

CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.o: CMakeFiles/clienteSimpe.dir/flags.make
CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.o: ../LindaDriver.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building CXX object CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/LindaDriver.cpp

CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/LindaDriver.cpp > CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.i

CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/LindaDriver.cpp -o CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.s

CMakeFiles/clienteSimpe.dir/Tuplas.cpp.o: CMakeFiles/clienteSimpe.dir/flags.make
CMakeFiles/clienteSimpe.dir/Tuplas.cpp.o: ../Tuplas.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Building CXX object CMakeFiles/clienteSimpe.dir/Tuplas.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/clienteSimpe.dir/Tuplas.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Tuplas.cpp

CMakeFiles/clienteSimpe.dir/Tuplas.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/clienteSimpe.dir/Tuplas.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Tuplas.cpp > CMakeFiles/clienteSimpe.dir/Tuplas.cpp.i

CMakeFiles/clienteSimpe.dir/Tuplas.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/clienteSimpe.dir/Tuplas.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Tuplas.cpp -o CMakeFiles/clienteSimpe.dir/Tuplas.cpp.s

# Object files for target clienteSimpe
clienteSimpe_OBJECTS = \
"CMakeFiles/clienteSimpe.dir/Socket.cpp.o" \
"CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.o" \
"CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.o" \
"CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.o" \
"CMakeFiles/clienteSimpe.dir/Tuplas.cpp.o"

# External object files for target clienteSimpe
clienteSimpe_EXTERNAL_OBJECTS =

clienteSimpe: CMakeFiles/clienteSimpe.dir/Socket.cpp.o
clienteSimpe: CMakeFiles/clienteSimpe.dir/clienteSimple.cpp.o
clienteSimpe: CMakeFiles/clienteSimpe.dir/controlTuplas.cpp.o
clienteSimpe: CMakeFiles/clienteSimpe.dir/LindaDriver.cpp.o
clienteSimpe: CMakeFiles/clienteSimpe.dir/Tuplas.cpp.o
clienteSimpe: CMakeFiles/clienteSimpe.dir/build.make
clienteSimpe: CMakeFiles/clienteSimpe.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_6) "Linking CXX executable clienteSimpe"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/clienteSimpe.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/clienteSimpe.dir/build: clienteSimpe

.PHONY : CMakeFiles/clienteSimpe.dir/build

CMakeFiles/clienteSimpe.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/clienteSimpe.dir/cmake_clean.cmake
.PHONY : CMakeFiles/clienteSimpe.dir/clean

CMakeFiles/clienteSimpe.dir/depend:
	cd /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/cmake-build-debug/CMakeFiles/clienteSimpe.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/clienteSimpe.dir/depend
