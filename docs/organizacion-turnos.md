# Organización de turnos del supermercado

## Situación actual

Juan Justo es quien se encarga de organizar el trabajo del supermercado. Para preparar el reparto necesita saber qué trabajadores van a estar disponibles, qué tareas hay que hacer ese día y cuáles puede realizar cada uno.

Los trabajadores le avisan cuando no pueden ir, cuando cambia su disponibilidad o cuando hay alguna tarea que no pueden realizar. Con esa información Juan organiza el trabajo del día.

En el supermercado trabajan nueve personas. No todas tienen siempre la misma disponibilidad ni pueden hacer las mismas tareas, así que el reparto depende bastante de quién pueda trabajar ese día.

Además, un reparto que ya estaba hecho puede dejar de servir si después falta alguien o cambia su disponibilidad.

## Cómo se suele organizar el trabajo

### Viaje de usuario 1: preparar el reparto

Juan suele preparar el reparto durante el fin de semana o antes de los días en los que hay que recoger palés en Granada.

Antes de empezar necesita saber:

- Qué trabajadores están disponibles.
- Qué tareas hay que hacer.
- Qué tareas puede realizar cada trabajador.
- Si ese día hay recogida de palés.
- Qué restricciones hay que respetar.

Primero mira quién puede trabajar y qué tareas hay pendientes. Si ese día hay recogida de palés, también tiene que comprobar quién puede desplazarse a Granada.

Después va repartiendo las tareas. Tiene que asegurarse de que una misma persona no tenga dos tareas al mismo tiempo, que la caja quede cubierta y que, cuando corresponda, haya gente disponible para recoger los palés.

Cuando esas tareas están cubiertas puede repartir el resto del trabajo, como reponer.

El dispositivo desde el que se consulte la información no cambia realmente este proceso. Lo importante es poder consultar los datos de los trabajadores, las tareas y las restricciones del día.

### Viaje de usuario 2: cambiar un reparto que ya estaba hecho

También puede pasar que Juan ya tenga el reparto preparado y después un trabajador avise de que no puede ir o cambie su disponibilidad.

En ese caso primero tiene que mirar qué tarea tenía esa persona. Después comprueba qué trabajadores siguen disponibles y vuelve a repartir las tareas que hayan quedado sin cubrir.

No siempre basta con cambiar una persona por otra. El nuevo reparto tiene que seguir cumpliendo las mismas reglas que el anterior.

Por ejemplo, si la persona que falta estaba en caja, hay que conseguir que otra persona pueda cubrirla. Si ese día hay recogida de palés, también tiene que seguir habiendo trabajadores disponibles para ir a Granada.

## Reglas que afectan al reparto

Para que un reparto pueda darse por bueno se tienen que cumplir varias condiciones:

- Una tarea solo se puede asignar a un trabajador que esté disponible.
- Una persona no puede realizar dos tareas al mismo tiempo.
- Si un trabajador se desplaza a Granada, durante ese tiempo no puede estar realizando otra tarea en el supermercado.
- La caja tiene que quedar cubierta.
- Cuando haya recogida de palés tiene que haber personal disponible para realizarla; hacen falta dos trabajadores.

Estas reglas son importantes porque puede parecer que un reparto está completo pero en realidad no se puede llevar a cabo.

## Problemas que se quieren resolver

### [HU001] Necesito tener clara la disponibilidad y qué tareas puede realizar cada trabajador

Cuando Juan empieza a organizar el día necesita tener actualizada la disponibilidad de los trabajadores y saber qué tareas puede realizar cada uno.

Si esa información no está clara puede contar con alguien que finalmente no esté disponible o asignarle una tarea que no pueda realizar.

Este es el primer problema que se va a trabajar, porque antes de comprobar un reparto hace falta tener la información correcta.

### [HU002] No sé si el reparto que he preparado es válido

Aunque Juan tenga toda la información necesaria, al hacer el reparto puede haber errores.

Una misma persona puede acabar con dos tareas que coinciden, puede quedar una tarea necesaria sin cubrir o puede haberse contado con alguien que no estaba disponible en ese momento.

El problema es poder saber si el reparto cumple las reglas del supermercado antes de darlo por bueno.

## Primeras entregas del proyecto

### Milestone 0: modelo base para los turnos

El primer milestone parte de la HU001.

A partir de esta historia se aplicará el diseño dirigido por el dominio (DDD) para modelar el problema. Los problemas que vayan apareciendo durante ese proceso se plantearán como issues y se irán resolviendo dentro del milestone.

El resultado será un modelo que permita trabajar con la información necesaria de una jornada del supermercado. Por ahora no se comprobará si un reparto está bien o mal.

Daremos el milestone 0 por terminado cuando los issues planteados estén resueltos y el modelo permita continuar con el siguiente milestone.

### Milestone 1: paquete con validación de repartos

Este milestone utiliza lo hecho en el anterior y trabaja con la HU002.

Sobre el resultado anterior se añadirá la lógica para comprobar si un reparto cumple las reglas del supermercado. Los problemas que aparezcan durante esa parte también se irán resolviendo mediante issues.

Además se añadirán pruebas automáticas para comprobar esa lógica. El milestone estará terminado cuando esas pruebas permitan comprobar el comportamiento implementado.
