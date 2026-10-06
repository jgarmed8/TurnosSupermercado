# Organización de tareas en el supermercado

## Situación actual

Juan Justo es quien se encarga de preparar los turnos y organizar el trabajo del supermercado.

Para este problema se parte de que los turnos y horarios ya están hechos. A partir de ahí Juan decide qué tarea hace cada trabajador durante su horario.

Para preparar el reparto necesita saber quién está disponible, qué tareas hay que hacer, a qué hora hacen falta y cuáles puede hacer cada persona.

En el supermercado trabajan nueve personas, aunque no están las nueve trabajando a la vez y tampoco todas pueden hacer siempre las mismas tareas.

Cuando se habla de un hueco se refiere simplemente a una tarea que necesita que haya una persona haciéndola durante un horario concreto.

Hay tareas que no se pueden dejar para después. La caja tiene que estar atendida durante el horario de apertura y, cuando toca recoger palés, tiene que haber trabajadores disponibles para ir a Granada. Otras tareas, como la reposición, se pueden dejar para más tarde si falta gente.

Por eso no basta con mirar cuántos trabajadores hay en total. También importa quién está trabajando a esa hora y qué tareas puede hacer cada persona.

## Cómo se suele organizar el trabajo

### Viaje de usuario 1: preparar el reparto

Juan suele preparar el reparto durante el fin de semana o antes de los días en los que toca recoger palés.

Antes de empezar comprueba:

- Qué trabajadores están disponibles.
- Qué horario tiene cada uno.
- Qué tareas hay que hacer.
- A qué hora hay que hacerlas.
- Qué tareas puede hacer cada trabajador.
- Si ese día hay recogida de palés.

Después va repartiendo el trabajo.

Tiene que evitar que una misma persona tenga dos tareas a la vez, asegurarse de que la caja siga atendida y tener en cuenta que, cuando toca recoger palés, los trabajadores que van a Granada no están disponibles en la tienda durante esas horas.

Cuando lo necesario está cubierto puede repartir otras tareas, como la reposición.

El dispositivo desde el que se consulte la información no cambia realmente este proceso. Lo importante es poder consultar los trabajadores, sus horarios y las tareas cuando haga falta.

### Viaje de usuario 2: cambiar un reparto que ya estaba hecho

También puede pasar que Juan ya tenga el reparto preparado y después un trabajador avise de que no puede ir.

Lo primero que hace es mirar qué tareas tenía esa persona y a qué horas tenía que hacerlas.

Después comprueba quién sigue disponible y qué puede hacer cada trabajador para volver a repartir las tareas que se hayan quedado sin cubrir.

No siempre basta con sustituir una persona por otra, porque el nuevo reparto tiene que seguir cumpliendo las mismas reglas.

Un caso que ocurrió fue un martes en el que dos trabajadores ya tenían que ir a Granada a recoger palés y después otro trabajador avisó de que no podía ir.

Durante esas horas había una persona menos en la tienda de la que Juan había tenido en cuenta cuando preparó el reparto. Los trabajadores que iban a Granada tenían que seguir haciendo esa tarea y la caja también tenía que seguir atendida.

Por eso hubo que cambiar de tarea a uno de los trabajadores para que se quedara en caja. La reposición se dejó para más tarde.

Este caso muestra que un reparto que al principio sirve puede dejar de hacerlo cuando falta alguien.

## Reglas del reparto

Para hacer el reparto hay que tener en cuenta varias cosas:

- Una tarea solo se puede dar a alguien que esté disponible a esa hora.
- Una persona no puede hacer dos tareas al mismo tiempo.
- Si alguien está en Granada, durante esas horas no puede hacer tareas en la tienda.
- La caja tiene que estar atendida durante el horario de apertura.
- Cuando hay recogida de palés tienen que estar disponibles los trabajadores necesarios para hacerla. Normalmente van dos.
- La reposición se puede dejar para más tarde si falta gente.

## Problemas que se quieren resolver

### [HU001] Cubrir las tareas necesarias con los trabajadores disponibles

Cuando Juan prepara el trabajo ya conoce los horarios, quién está disponible, qué tareas hay que hacer y cuáles puede hacer cada trabajador.

El problema es decidir quién hace cada tarea para que las que no se pueden dejar para después estén cubiertas y se respeten las reglas del supermercado.

Por ejemplo, cuando hay recogida de palés hay trabajadores que durante unas horas están fuera de la tienda. Juan tiene que organizar a los que quedan para que la caja y las demás tareas necesarias sigan atendidas.

El problema está resuelto cuando todas las tareas necesarias tienen a alguien que pueda hacerlas.

Si con las personas disponibles no se puede cubrir todo lo necesario, no hay una solución para ese turno.

Para Juan esto significa poder dejar organizado el día sabiendo que lo importante está atendido.

### [HU002] Reorganizar el reparto cuando falta un trabajador

Puede pasar que Juan ya tenga el reparto preparado y después una persona diga que no puede ir.

Las tareas que tenía pueden quedarse sin cubrir, así que hay que volver a repartirlas entre los trabajadores que siguen disponibles.

El caso del martes explicado anteriormente es un ejemplo de este problema.

El problema está resuelto cuando las tareas necesarias vuelven a tener a alguien que pueda hacerlas.

Si con las personas que quedan no se puede cubrir todo lo necesario, no hay una solución para esa situación.

Para Juan esto significa poder reaccionar ante una ausencia sin dejar las tareas importantes sin atender.

## Primeras entregas del proyecto

### Milestone 0: modelo

El primer milestone parte de la HU001.

A partir de esta historia se aplicará DDD para modelar el problema. Los problemas que aparezcan durante ese proceso se plantearán como issues y servirán para guiar el modelado.

El milestone estará terminado cuando esos problemas estén resueltos y el modelo obtenido permita continuar con el siguiente milestone.

### Milestone 1: implementación verificable

Este milestone parte del modelo obtenido a partir de la HU001 en el milestone anterior y trabaja con la HU002.

A partir de esta historia se desarrollará la lógica de negocio que resulte necesaria. Los problemas que aparezcan durante este proceso se plantearán como issues.

El resultado se comprobará mediante pruebas automáticas.

El milestone estará terminado cuando los problemas planteados estén resueltos y las pruebas permitan comprobar automáticamente el comportamiento desarrollado a partir de la HU002.
