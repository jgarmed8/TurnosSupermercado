# Automatización de tareas

## Dependencias

El proyecto está hecho en Go y utiliza Go Modules para gestionar las dependencias. Ya tengo `go.mod` y, por ahora, el modelo de turnos no necesita librerías externas. Al ejecutar `go mod download` no hay módulos que descargar.

## Qué necesito automatizar

Después del objetivo 2 tengo seis archivos Go en `internal/turnos`. En el issue #23 planteé automatizar la comprobación de la sintaxis de los archivos Go y la instalación de dependencias.

## Criterios para elegir el gestor

He comparado Make, Just, Task, Mage y XC.

En cuanto a estándares, `make` está definido por POSIX. Just, Task, Mage y XC son gestores distintos y no siguen esa especificación. Go tampoco exige utilizar un gestor de tareas concreto, así que no he descartado ninguna herramienta por este motivo.

Para elegir me he fijado en:

- **Tiempo de ejecución:** cuánto tarda cada herramienta en ejecutar la misma comprobación.
- **Instalación:** si necesito instalar algo aparte para utilizarla.

## Pruebas de ejecución

Hice las pruebas en mi Ubuntu con Go 1.27.4. El archivo `go.mod` indica Go 1.26.8 como versión mínima, pero todas las mediciones las hice con la versión que tengo instalada.

Utilicé Hyperfine con tres ejecuciones de calentamiento y treinta mediciones por comando.

Para comparar las cinco herramientas utilicé `gofmt -e -l .`, que comprueba la sintaxis de los archivos Go sin modificarlos ni compilar el proyecto.

Los resultados fueron:

- Sin gestor: **109,4 ± 5,0 ms**.
- Make: **114,2 ± 4,7 ms**.
- Just: **116,6 ± 5,5 ms**.
- Task: **157,0 ± 5,9 ms**.
- Mage: **490,5 ± 12,5 ms**.
- XC: **112,2 ± 4,4 ms**.

![Comparación de los gestores de tareas](capturas/ejecucion-sintaxis.png)

## Instalación de las herramientas

En mi Ubuntu ya tenía Make instalado. Instalé Just con `apt` y Task con su instalador. Mage v1.17.2 lo descargué como binario desde GitHub Releases.

Medí estos tiempos de instalación:

- Task: **5,124 s**.
- Mage: **1,08 s**.

![Instalación de Task](capturas/instalacion-task.png)

![Instalación de Mage](capturas/instalacion-mage.png)

No cronometré la instalación inicial de Just. Tampoco he utilizado los tiempos de Task y Mage para comparar cuál se instala más rápido, porque los instalé de formas distintas.

Además, comprobé las herramientas disponibles en `golang:1.27.4-bookworm`. Esa imagen ya traía Make, mientras que Just, Task y Mage no estaban instalados.

![Herramientas disponibles en la imagen de Go](capturas/entorno-go.png)

XC v0.9.0 lo descargué como binario para Linux desde GitHub Releases, sin compilarlo, y tardó **1,41 s**.

![Instalación de XC](capturas/xc-instalacion.png)

## Elección

Me quedo con Make.

Make, Just y XC tardaron alrededor de 115 ms, muy cerca de los 112 ms de ejecutar `gofmt` directamente. Task y Mage tardaron bastante más.

Make ya estaba instalado en mi Ubuntu y en la imagen de Go que comprobé, así que prefiero mantenerlo sin instalar otro gestor para las tareas actuales.

## Comprobación de la tarea

He añadido dos tareas al `Makefile`:

- `make install`, que ejecuta `go mod download`.
- `make check`, que utiliza `gofmt -e -l .` para comprobar la sintaxis de los archivos Go.

Con los archivos actuales del proyecto, `make check` termina con código de salida `0`.

También he probado un archivo temporal con un error de sintaxis. La tarea detectó el error y terminó con un código distinto de `0`.

Después eliminé el archivo temporal y volví a comprobar que la tarea terminaba correctamente.

La tarea solo comprueba la sintaxis, por lo que no detecta errores de tipos.
