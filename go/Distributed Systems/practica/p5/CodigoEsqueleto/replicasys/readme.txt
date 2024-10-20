>>> Contenido del módulo "replicasys" :
        * "g.mod": fichero de configuración básica de modulo
        * "cmd": directorio  ejecutables de replica almacenamiento y gestor vistas
        * "internal": paquetes soporte para puesta en marcha de servicio Almace.
            - "integracionrepsystest": paquete tests integracion servicio réplicas
            - "srvalm": paquete funcionalidad réplica/servidor almacenamiento
            - "integraciongvtest": paquete tests de integración servicio de vistas
            - "srvvts": paquete para código de funcionamiento de gestor vistas
            - "msgsys": paquete código de mensajería
            - "cltssh": paquete cliente ssh con autentificación clave pública
        * "pkg" : librerías cliente de servicios almacenamiento y vistas
        * "vendor": código completo local de modulos externos necesarios (ssh..)

* Esta version puede ser utilizada como ejemplo, consulta, guía o referencia
    - Para implementar la práctica 5:
        - Reutilizando vuestros ficheros implementados en práctica 4:
            - servidor_vistas.go y integracion_test.go (renombrado)
    - Refactoriza y extiende el modulo v2 del gestor de vistas de practica 4:
        - Se renombra modulo a "replicasys"
        - Se renombra paquete "gvcomun" a "comun"
        - Se modifica ligeramente paquete "msgsys":
            - Función "Send" y "SendReceive" para programar fallos de conexión
        - Se añaden 4 paquetes nuevos :
            - "integracionrepsystest", para tests de integración de almacenamiento
            - "Main" de ejecución de servidor almacenamiento
            - "srvalm", funcionalidad de servidor/replica de almacenamiento
            - "cltalm", funcionalidad de cliente de almacenamiento 
    - Se ha incluido directorio "vendor" para este modulo
        - Con todo los modulos y paquetes necesarios para paquete externo "ssh"
    - Incluye tests de integración de servicio replicas para las 3 primeras pruebas
    
* Para una versión funcional de la práctica 5 solo requeriría
        - modificar el PATH completo del directorio del modulo "replicasys"
            - en fichero "cltssh_test.go", línea 13
            - y fichero "integración_repsys_test.go", constante "PATH"
        - modificar PATH completo del fichero de clave privada para ssh
            - ubicado en fichero "integración _test.go", constante "PRIVKEYFILE"
            - y en fichero "cltssh_test.go", línea 15
        - y completar código : 
            - en el paquete "srvalm"
            - de pruebas que faltan en el paquete "integracion_repsys_test",
                - Están disponibles las 3 primeras pruebas completadas

* Pruebas iniciales en terminales separadas para cada servidor (vistas o almace.)
    - Comentando el arranque automático por ssh de servidores en código de tests
    - Se puede seguir más fácil errores de compilación y trazas de depuracion

* Probar todos los tests del modulo con : go test -v
    - Su ejecución es concurrente, luego puede haber errores por interferencias
        - Si se reejecuta conserva en cache ejecuciones correctas de tests
            - Luego menos interferencias
            - y tiempos de puesta en marcha y comunicación van mejor
            - Borrado de cache : go clean -cache -modcache -i -r
    - Clientes de almacenamiento se ponen en marcha en local, no por ssh
        - Cada uno escucha en puerto diferentes
        - Ejecución concurrente de clientes con gorutinas

            
* Los test de integración utilizan el mecanismo de "Subtests"
    - Documentación en sección "Subtests" de :
        - https://golang.org/pkg/testing/
        
* El modo depuración de tests se obtiene con :
        - o "go test -v"
        - o utilizando la opción "debug test" en vscode o editor similiar
    - con ello, además del depurador, visualiza salidas de "fmt" y "log"
    - SI EJECUCION ERRONEA de tests os pueden quedar PROCESOS SSH SIN TERMINAR
        - ELIMINARLOS con "pkill main"
            - En cada máquina distribuida si ejecución en distribuido
        
* Los tiempos de espera para :
        - arranque de servidores remotos con ssh en diferentes tests
        - utilización de funciones "SendReceive" y "ReceiveTimed"
    - Dependen de máquina o entorno distribuido de ejecución
        - Están ajustados para 2ª ejecución distribuida en el laboratorio 1.02
            - Primera ejecución es erronea porque busca ficheros codigo
                - en servidor remoto NFS (planta baja)
            - Segunda ejecución los obtiene de cache local del sistema de ficheros
            
* Se utiliza en el código el concepto de "embedded structs"
        - https://golang.org/ref/spec#Struct_types
        - https://stackoverflow.com/questions/34079466/embedded-struct
    - Para ampliar de forma directa la funcionalidad :
        - de los servidores almacenammiento, gestor de vistas, nodo de test
            - En los tipos de datos :
                - "srvalm.ServAlm
                - "srvvts.ServVistas",
                - "cltalm.ClienteAlm"
                - "integracionrepsystes.testServer"
            - en lo que respecta al sistema de mensajería para todos ellos
                - empotrando el tipo, campos y métodos de "msgsys. MsgSys"
                    - métodos
                        - Me,Send,SendReceive,InternalSend,Receive,ReceiveTimed