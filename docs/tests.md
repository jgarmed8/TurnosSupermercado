# Herramientas para los tests

Necesitamos comprobar que Juan puede repartir las tareas entre los trabajadores respetando sus horarios, las tareas que pueden realizar y las asignaciones que ya tienen.

He buscado alternativas en la [documentación de Go](https://pkg.go.dev/testing) y en el apartado de testing de [Awesome Go](https://github.com/avelino/awesome-go#testing).

## Biblioteca de aserciones

### Criterios

- Cómo se comprueba el resultado de una asignación y el error que devuelve cuando no se puede hacer.
- Tiempo de ejecución del mismo test con cada opción.

### Opciones

- Sin biblioteca de aserciones, utilizando if, t.Errorf y errors.Is.
- [Testify](https://github.com/stretchr/testify).
- [go-cmp](https://github.com/google/go-cmp), para comparar valores.
- [Gomega](https://github.com/onsi/gomega), que utiliza matchers.

## Framework de tests

### Criterios

- Cómo se agrupan varios casos del mismo problema y se identifica cuál ha fallado.
- Tiempo de ejecución de los mismos casos.

### Opciones

- [testing](https://pkg.go.dev/testing).
- [Ginkgo](https://github.com/onsi/ginkgo).
- [GoConvey](https://github.com/smartystreets/goconvey).

## Herramienta CLI

### Criterios

- Que termine por sí sola después de ejecutar los tests.
- Que funcione con los fuentes en solo lectura.
- Tiempo de ejecución de los mismos tests.

### Opciones

- [go test](https://pkg.go.dev/cmd/go).
- [Ginkgo CLI](https://onsi.github.io/ginkgo/).
- [GoConvey CLI](https://github.com/smartystreets/goconvey).

También he considerado gotestsum, aunque utiliza go test para ejecutar las pruebas y cambia la presentación de los resultados.

## Forma de medir

Utilizo una copia del código del reparto en un directorio temporal para probar las alternativas sin añadir sus dependencias al proyecto.

Las medidas se realizan con Go 1.27.4 y Hyperfine 1.18.0 en Linux, con tres ejecuciones de calentamiento y treinta mediciones. Con go test utilizo -count=1 para ejecutar los tests sin reutilizar resultados anteriores.
