# TurnosSupermercado

## Problema

El responsable de un supermercado tiene que preparar los turnos de los trabajadores y repartir las tareas que hay que hacer cada día.

Una vez que los turnos están hechos, tiene que decidir qué tarea hace cada persona durante su horario.

Los lunes, martes, jueves y viernes suele haber recogida de palés en Granada. Normalmente van dos trabajadores y, mientras están fuera, hay menos gente disponible en la tienda.

Si falta alguien o cambia su disponibilidad, puede ser necesario volver a repartir parte del trabajo.

## De dónde viene esto

Este problema lo conozco porque mi padre es quien se encarga de organizar el trabajo del supermercado. Yo también ayudo allí durante el verano, así que he visto cómo se hacen los turnos y qué pasa cuando falta gente.

En el supermercado trabajan nueve personas. No todas trabajan al mismo tiempo ni pueden hacer siempre las mismas tareas.

## Cómo se organiza ahora

Los turnos se llevan en un Excel guardado en el ordenador de casa del responsable.

Cuando el cuadrante está hecho, se avisa a cada trabajador por WhatsApp. Si hay algún cambio, hay que modificar el Excel y volver a avisar a las personas afectadas.

## Un caso real

Un martes faltó uno de los trabajadores y ese mismo día ya había dos personas que tenían que ir a Granada a recoger palés.

Durante esas horas había menos gente en la tienda de la que se había tenido en cuenta al preparar el reparto.

Hubo que cambiar a otro trabajador de tarea para que se quedara en caja y la reposición se dejó para más tarde.

## Datos disponibles

Para organizar el trabajo se dispone de:

- los trabajadores del supermercado;
- los turnos y horarios de cada uno;
- su disponibilidad;
- las tareas que hay que hacer;
- los días de recogida de palés;
- qué trabajadores pueden ir a Granada.

Los turnos actuales están guardados en el Excel que utiliza el responsable.

## Qué hay que tener en cuenta

La caja tiene que estar atendida durante el horario de apertura.

Cuando hay recogida de palés tienen que estar disponibles los trabajadores necesarios para realizarla. Normalmente van dos y, mientras están en Granada, no se puede contar con ellos para las tareas de la tienda.

Si falta alguien puede ser necesario cambiar el reparto que ya estaba hecho.

La reposición puede dejarse para más tarde si no hay suficiente gente para hacerlo todo.

## Qué hace falta procesar

Para hacer el reparto hay que tener en cuenta qué trabajadores están disponibles, sus horarios y las tareas que hay que hacer.

También hay que comprobar qué personas pueden hacer cada tarea y quién estará fuera de la tienda cuando haya recogida de palés.

Si cambia la disponibilidad de alguien, hay que volver a repartir las tareas que se hayan quedado sin cubrir.

## Por qué hace falta que esté en la nube

Ahora mismo el cuadrante está guardado en un Excel en el ordenador de casa del responsable, mientras que los trabajadores reciben sus turnos y los cambios por WhatsApp.

Esto hace que la planificación dependa de un único ordenador. Para consultar o modificar el cuadrante completo hay que acceder a ese equipo, aunque los avisos a los trabajadores sí puedan hacerse desde el móvil.

Tener el cuadrante accesible desde distintos dispositivos permitiría consultar y modificar la misma planificación sin depender de ese único ordenador.

## Juego de rol

[Tarjeta utilizada en el juego de rol](img/tarjeta-cliente.png)

## Documentación

- [Configuración del repositorio](docs/objetivo-0.md)

## Planificación

La planificación del objetivo 1 está en:

- [Organización de tareas en el supermercado](docs/organizacion-turnos.md)

### Historias de usuario en GitHub

- [HU001 - Cubrir las tareas necesarias con los trabajadores disponibles](https://github.com/jgarmed8/TurnosSupermercado/issues/2)
- [HU002 - Reorganizar el reparto cuando falta un trabajador](https://github.com/jgarmed8/TurnosSupermercado/issues/3)

### Milestones

- [Milestone 0: modelo](https://github.com/jgarmed8/TurnosSupermercado/milestone/1)
- [Milestone 1: implementación verificable](https://github.com/jgarmed8/TurnosSupermercado/milestone/2)
