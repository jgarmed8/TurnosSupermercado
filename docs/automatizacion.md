# Automatización de tareas

## Dependencias

El proyecto está hecho en Go y utiliza Go Modules para gestionar las dependencias. Ya tengo `go.mod` y, por ahora, el modelo de turnos no necesita librerías externas. Al ejecutar `go mod download` no hay módulos que descargar.

## Qué necesito automatizar

Después del objetivo 2 tengo seis archivos Go en `internal/turnos`. En el issue #23 planteé automatizar la comprobación de que los paquetes compilan y la instalación de dependencias.

## Criterios para elegir el gestor

He comparado Make, Just, Task, Mage y XC.

En cuanto a estándares, `make` está definido por POSIX. Just, Task, Mage y XC son gestores distintos y no siguen esa especificación. Go tampoco exige utilizar un gestor de tareas concreto, así que no he descartado ninguna herramienta por este motivo.

Para elegir me he fijado en:

- **Tiempo de ejecución:** cuánto tarda cada herramienta en ejecutar la misma comprobación.
- **Instalación:** si necesito instalar algo aparte para utilizarla.

## Pruebas de ejecución

Hice las pruebas en mi Ubuntu con Go 1.27.4. El archivo `go.mod` indica Go 1.26.8 como versión mínima, pero todas las mediciones las hice con la versión que tengo instalada.

Utilicé Hyperfine con tres ejecuciones de calentamiento y treinta mediciones por comando.

### Make, Just, Task y Mage

Para comparar las cuatro herramientas utilicé `go build ./...`, que comprueba si los paquetes Go compilan y detecta errores de nombres y tipos.

Los resultados fueron:

- Sin gestor (primera comparación): **145,7 ± 6,9 ms**.
- Make: **146,8 ± 4,0 ms**.
- Just: **146,3 ± 5,4 ms**.
- Task: **192,6 ± 5,6 ms**.
- Mage: **528,3 ± 9,0 ms**.

![Comparación utilizando go build](capturas/ejecucion-build.png)

### XC

Después de elegir Make probé XC v0.9.0, que permite definir tareas en un archivo Markdown. Lo comparé con Make y con `go build ./...` directamente para ver si me interesaba cambiar de gestor.

Repetí las medidas con Hyperfine, utilizando tres calentamientos y treinta ejecuciones por comando.

Los resultados fueron:

- Sin gestor (comparación con XC): **139,2 ± 3,9 ms**.
- Make: **138,9 ± 4,3 ms**.
- XC: **141,3 ± 4,3 ms**.

![Comparación de XC con Make](capturas/xc-ejecucion.png)

Estas medidas son de otra ejecución de Hyperfine, así que no las comparo directamente con las de la primera prueba.

## Instalación de las herramientas

En mi Ubuntu ya tenía Make instalado. Para probar las otras herramientas instalé Just con `apt`, Task con su instalador y Mage con `go install`.

Medí estos tiempos de instalación:

- Task: **5,124 s**.
- Mage: **6,544 s**.

![Instalación de Task](capturas/instalacion-task.png)

![Instalación de Mage](capturas/instalacion-mage.png)

No cronometré la instalación inicial de Just. Tampoco he utilizado los tiempos de Task y Mage para comparar cuál se instala más rápido, porque los instalé de formas distintas.

Además, comprobé las herramientas disponibles en `golang:1.27.4-bookworm`. Esa imagen ya traía Make, mientras que Just, Task y Mage no estaban instalados.

![Herramientas disponibles en la imagen de Go](capturas/entorno-go.png)

XC lo instalé con `go install` en mi Ubuntu y tardó **6,030 s**.

![Instalación de XC](capturas/xc-instalacion.png)

## Elección

Me quedo con Make.

En la primera comparación, Make y Just tardaron prácticamente lo mismo, mientras que Task y Mage necesitaron más tiempo. Después probé XC y tampoco encontré una diferencia clara respecto a Make.

Make ya estaba instalado tanto en mi Ubuntu como en la imagen de Go que comprobé. Para las tareas que necesito ahora, prefiero mantenerlo sin instalar otro gestor.

## Comprobación de la tarea

He añadido dos tareas al `Makefile`:

- `make install`, que ejecuta `go mod download`.
- `make check`, que ejecuta `go build ./...`.

Con los archivos actuales del proyecto, `make check` termina con código de salida `0`.

También preparé un archivo temporal con una función que debía devolver un `int`, pero devolvía un `string`. Al ejecutar `make check`, detectó el error de tipos y terminó con un código distinto de `0`.

La prueba se hizo fuera del modelo de turnos. `check` comprueba la compilación de los paquetes, sin ejecutar tests ni `go vet`.

![Comprobación de make check](capturas/check-build.png)
