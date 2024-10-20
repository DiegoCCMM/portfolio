#!/bin/bash
#*****************************************************************
# Funcionamiento: 
#          1) Dar permisos de ejecución al script: "chmod u+x script_LindaServer.bash"
#          2) Ejecutar mediante "script_LindaServer.bash puerto"
#*****************************************************************

MAKE_FILE=Makefile_LindaServer

#----------------------------------------------------------------
if [[ "$#" -ne 3 ]]
then
	echo "Numero de parámetros incorrecto"
	exit
fi

make -f $MAKE_FILE &>/dev/null
resMake=$?
if [[ "$resMake" == 0 ]]
then    #ha funcionado bien
	echo "**************************************"
	echo "* Se ha compilado todo correctamente *"
	echo "**************************************"
	./LindaServer "$1" "$2" "$3"
else
	exit
fi

exit 0
