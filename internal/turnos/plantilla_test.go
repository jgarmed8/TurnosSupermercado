package turnos

import (
	"slices"
	"testing"
)

func TestPlantilla(t *testing.T) {
	ana := &Trabajador{nombre: "Ana"}
	otraAna := &Trabajador{nombre: "Ana"}
	plantilla := Plantilla{}

	for _, trabajador := range []*Trabajador{ana, otraAna} {
		if !plantilla.agregar(trabajador) || !plantilla.contiene(trabajador) {
			t.Fatal("no se ha añadido el trabajador")
		}
	}

	if plantilla.agregar(ana) || plantilla.agregar(nil) {
		t.Error("se ha aceptado un duplicado o una referencia nula")
	}

	if plantilla.contiene(&Trabajador{nombre: "Ana"}) || plantilla.contiene(nil) {
		t.Error("se ha encontrado un trabajador ajeno a la plantilla")
	}

	if !slices.Equal(plantilla.trabajadores, []*Trabajador{ana, otraAna}) {
		t.Error("la plantilla no conserva los dos trabajadores distintos")
	}
}
