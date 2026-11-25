# Herramientas para los tests

En el Milestone 1 vamos a comprobar la lógica del reparto de tareas.

Los turnos ya están hechos. Lo que necesitamos es comprobar qué trabajadores pueden cubrir cada hueco teniendo en cuenta su disponibilidad, las tareas que pueden realizar y las asignaciones que ya tienen.

También hay que probar los casos en los que una asignación no se puede hacer. Por ejemplo, cuando un trabajador solo está disponible durante parte del horario o cuando ya está haciendo otra tarea.

## Búsqueda de alternativas

Antes de elegir las herramientas voy a comparar por separado las bibliotecas de aserciones, los frameworks de tests y las herramientas que permiten ejecutarlos.

La comparación se hará utilizando casos del propio proyecto. Primero estableceré los criterios y después comprobaré las alternativas para decidir cuáles utilizar.

## Biblioteca de aserciones

### Criterios

Necesitamos comprobar que el resultado obtenido coincide con el esperado.

También tendremos que comprobar qué errores devuelve el reparto cuando una asignación no es válida.

Para comparar las alternativas utilizaré los mismos casos, viendo cómo se escriben las comprobaciones y si resulta fácil entender qué ha fallado.

### Alternativas

- **testing:** permite comprobar los resultados utilizando `t.Errorf` y `t.Fatal`.
- **Testify:** incluye funciones de aserción para comparar resultados y comprobar errores.
- **go-cmp:** permite comparar valores y mostrar las diferencias entre ellos.
- **Gomega:** utiliza matchers para expresar las comprobaciones.

La elección se hará después de comprobar cómo funciona cada alternativa con los tests del proyecto.

## Framework de tests

### Criterios

Vamos a probar distintos horarios, trabajadores y situaciones del reparto.

Necesitamos poder separar los casos para saber cuál falla sin tener que revisar todos los demás.

También comprobaré cómo se organizan los mismos tests con cada herramienta y cuánto tardan en ejecutarse.

### Alternativas

- **testing:** permite organizar los casos mediante `t.Run`.
- **GoConvey:** permite agrupar distintas comprobaciones.
- **Ginkgo:** organiza las pruebas mediante especificaciones.

La elección dependerá de cómo se puedan expresar y ejecutar nuestros tests.

## Herramienta para ejecutar los tests

### Criterios

Necesitamos ejecutar todos los tests desde la raíz del proyecto y saber si alguno falla.

También comprobaré que las herramientas no necesiten escribir dentro de la carpeta de los fuentes, porque más adelante el proyecto tendrá que poder probarse con el código en solo lectura.

Compararé los tiempos de ejecución utilizando los mismos tests y evitando que se reutilicen resultados anteriores de la caché.

### Alternativas

- **go test:** permite compilar y ejecutar los tests de Go.
- **Ginkgo CLI:** permite ejecutar las pruebas organizadas con Ginkgo.

También he tenido en cuenta **gotestsum**, pero utiliza `go test` para ejecutar las pruebas y cambia la forma de presentar los resultados. Por eso no lo consideraré otro ejecutor independiente.

## Pruebas y elección

Para realizar las comparaciones utilizaré los tests de disponibilidad y tareas permitidas del paquete `internal/turnos`.

Las medidas se harán manteniendo las mismas condiciones para las distintas alternativas.

Con `go test` utilizaré `-count=1` para evitar que se reutilicen los resultados de la caché.

Todavía no he elegido las herramientas definitivas. Primero realizaré las pruebas y después añadiré aquí los resultados y los motivos de la elección.

## F.I.R.S.T.

Los tests utilizan fechas fijas y no necesitan consultar servicios externos.

Cada caso comprueba automáticamente si obtiene el resultado esperado.

También ejecutaré las pruebas varias veces y en distinto orden para comprobar que no dependen unas de otras.

Los tests de disponibilidad se escribieron antes de implementar esa comprobación. Haré lo mismo con las tareas permitidas y los siguientes problemas del reparto.

Mediré los tiempos antes de elegir las herramientas para comprobar que los tests se ejecutan rápidamente.
