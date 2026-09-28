# Milestones

## Milestone 0: Paquete de jornadas

### Objetivo

Representar en código una jornada real del supermercado, sin entrar todavía en cómo se van a repartir las tareas.

### Qué se entrega

Un paquete en Python que permita representar los trabajadores, su disponibilidad, las tareas que puede realizar cada uno, las tareas que hay que cubrir durante la jornada y las restricciones que afectan al reparto.

### Cómo se comprueba

Se usarán los casos reales descritos en las jornadas de usuario: una jornada normal y otra en la que haya recogida de palés.

En ambos casos se tiene que poder representar toda la información necesaria sobre trabajadores, tareas y restricciones.

En este milestone todavía no se realiza el reparto de tareas.

## Milestone 1: Comprobación de repartos

### Objetivo

Comprobar si el reparto de tareas de una jornada cumple las reglas del supermercado.

### Qué se entrega

Se amplía el paquete en Python del milestone anterior para que, además de representar una jornada, permita revisar si un reparto de tareas cumple las reglas.

También se añaden pruebas automáticas usando distintos casos del supermercado.

### Cómo se comprueba

Se probarán repartos correctos y otros que tengan algún problema.

El paquete deberá distinguir cuándo un reparto es válido y cuándo no lo es.

Entre los casos de prueba habrá trabajadores no disponibles, una persona con dos tareas al mismo tiempo, la caja sin cubrir o días en los que no se pueda realizar la recogida de palés.
