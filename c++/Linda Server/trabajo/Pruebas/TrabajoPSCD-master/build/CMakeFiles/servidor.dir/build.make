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
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build

# Include any dependencies generated for this target.
include CMakeFiles/servidor.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/servidor.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/servidor.dir/flags.make

CMakeFiles/servidor.dir/Socket.cpp.o: CMakeFiles/servidor.dir/flags.make
CMakeFiles/servidor.dir/Socket.cpp.o: ../Socket.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/servidor.dir/Socket.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/servidor.dir/Socket.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Socket.cpp

CMakeFiles/servidor.dir/Socket.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/servidor.dir/Socket.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Socket.cpp > CMakeFiles/servidor.dir/Socket.cpp.i

CMakeFiles/servidor.dir/Socket.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/servidor.dir/Socket.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Socket.cpp -o CMakeFiles/servidor.dir/Socket.cpp.s

CMakeFiles/servidor.dir/servidor.cpp.o: CMakeFiles/servidor.dir/flags.make
CMakeFiles/servidor.dir/servidor.cpp.o: ../servidor.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/servidor.dir/servidor.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/servidor.dir/servidor.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/servidor.cpp

CMakeFiles/servidor.dir/servidor.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/servidor.dir/servidor.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/servidor.cpp > CMakeFiles/servidor.dir/servidor.cpp.i

CMakeFiles/servidor.dir/servidor.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/servidor.dir/servidor.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/servidor.cpp -o CMakeFiles/servidor.dir/servidor.cpp.s

CMakeFiles/servidor.dir/controlTuplas.cpp.o: CMakeFiles/servidor.dir/flags.make
CMakeFiles/servidor.dir/controlTuplas.cpp.o: ../controlTuplas.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/servidor.dir/controlTuplas.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/servidor.dir/controlTuplas.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/controlTuplas.cpp

CMakeFiles/servidor.dir/controlTuplas.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/servidor.dir/controlTuplas.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/controlTuplas.cpp > CMakeFiles/servidor.dir/controlTuplas.cpp.i

CMakeFiles/servidor.dir/controlTuplas.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/servidor.dir/controlTuplas.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/controlTuplas.cpp -o CMakeFiles/servidor.dir/controlTuplas.cpp.s

CMakeFiles/servidor.dir/Tuplas.cpp.o: CMakeFiles/servidor.dir/flags.make
CMakeFiles/servidor.dir/Tuplas.cpp.o: ../Tuplas.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building CXX object CMakeFiles/servidor.dir/Tuplas.cpp.o"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/servidor.dir/Tuplas.cpp.o -c /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Tuplas.cpp

CMakeFiles/servidor.dir/Tuplas.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/servidor.dir/Tuplas.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Tuplas.cpp > CMakeFiles/servidor.dir/Tuplas.cpp.i

CMakeFiles/servidor.dir/Tuplas.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/servidor.dir/Tuplas.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/Tuplas.cpp -o CMakeFiles/servidor.dir/Tuplas.cpp.s

# Object files for target servidor
servidor_OBJECTS = \
"CMakeFiles/servidor.dir/Socket.cpp.o" \
"CMakeFiles/servidor.dir/servidor.cpp.o" \
"CMakeFiles/servidor.dir/controlTuplas.cpp.o" \
"CMakeFiles/servidor.dir/Tuplas.cpp.o"

# External object files for target servidor
servidor_EXTERNAL_OBJECTS =

servidor: CMakeFiles/servidor.dir/Socket.cpp.o
servidor: CMakeFiles/servidor.dir/servidor.cpp.o
servidor: CMakeFiles/servidor.dir/controlTuplas.cpp.o
servidor: CMakeFiles/servidor.dir/Tuplas.cpp.o
servidor: CMakeFiles/servidor.dir/build.make
servidor: CMakeFiles/servidor.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Linking CXX executable servidor"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/servidor.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/servidor.dir/build: servidor

.PHONY : CMakeFiles/servidor.dir/build

CMakeFiles/servidor.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/servidor.dir/cmake_clean.cmake
.PHONY : CMakeFiles/servidor.dir/clean

CMakeFiles/servidor.dir/depend:
	cd /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build /home/diego/Desktop/PSCD/trabajo/Pruebas/TrabajoPSCD-master/build/CMakeFiles/servidor.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/servidor.dir/depend
