package turnos

import (
	"errors"
	"slices"
	"testing"
	"time"
)

func trabajadorDePrueba(nombre string, tramo TramoHorario, tareas ...Tarea) *Trabajador {
	permitidas := make(map[Tarea]struct{})

	for _, tarea := range tareas {
		permitidas[tarea] = struct{}{}
	}

	return &Trabajador{
		nombre:           nombre,
		disponibilidad:   []TramoHorario{tramo},
		tareasPermitidas: permitidas,
	}
}

func TestAsignar(t *testing.T) {
	dia := Dia{2026, time.November, 24}
	horario := tramoDePrueba(10, 0, 12, 0)

	ana := trabajadorDePrueba("Ana", horario, "caja", "palés")
	luis := trabajadorDePrueba("Luis", horario, "palés")
	eva := trabajadorDePrueba("Eva", horario, "caja")
	corto := trabajadorDePrueba("Pablo", tramoDePrueba(10, 0, 10, 30), "palés")

	plantilla := Plantilla{
		trabajadores: []*Trabajador{ana, luis, eva, corto},
	}

	casos := []struct {
		nombre     string
		dia        Dia
		indice     int
		trabajador *Trabajador
		esperado   error
	}{
		{"asignación válida", dia, 1, luis, nil},
		{"trabajador ocupado", dia, 1, ana, ErrTrabajadorOcupado},
		{"tarea no permitida", dia, 1, eva, ErrTareaNoPermitida},
		{"horario incompleto", dia, 1, corto, ErrSinDisponibilidad},
		{"ajeno a plantilla", dia, 1, trabajadorDePrueba("Luis", horario, "palés"), ErrTrabajadorAjeno},
		{"sin trabajador", dia, 1, nil, ErrTrabajadorAjeno},
		{"hueco cubierto", dia, 0, eva, ErrHuecoCubierto},
		{"hueco inexistente", dia, 4, luis, ErrHuecoNoExiste},
		{"índice negativo", dia, -1, luis, ErrHuecoNoExiste},
		{"día sin reparto", Dia{2026, time.November, 25}, 1, luis, ErrDiaSinReparto},
		{"horario inválido", dia, 2, luis, ErrHorarioInvalido},
		{"tarea consecutiva", dia, 3, ana, nil},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			huecos := []Hueco{
				{tarea: "caja", tramo: tramoDePrueba(10, 0, 11, 0), cubiertoPor: ana},
				{tarea: "palés", tramo: horario},
				{tarea: "palés", tramo: tramoDePrueba(10, 0, 10, 0)},
				{tarea: "caja", tramo: tramoDePrueba(11, 0, 12, 0)},
			}

			esperado := slices.Clone(huecos)

			reparto := RepartosDeTareas{
				plantilla:    &plantilla,
				huecosPorDia: map[Dia][]Hueco{dia: huecos},
			}

			err := reparto.Asignar(caso.dia, caso.indice, caso.trabajador)

			if !errors.Is(err, caso.esperado) {
				t.Fatalf("esperado %v, obtenido %v", caso.esperado, err)
			}

			if caso.esperado == nil {
				esperado[caso.indice].cubiertoPor = caso.trabajador
			}

			if !slices.Equal(reparto.huecosPorDia[dia], esperado) {
				t.Error("el reparto no ha quedado como se esperaba")
			}
		})
	}
}

func TestCandidatos(t *testing.T) {
	dia := Dia{2026, time.November, 24}
	horario := tramoDePrueba(10, 0, 12, 0)

	ana := trabajadorDePrueba("Ana", horario, "palés", "caja")
	luis := trabajadorDePrueba("Luis", horario, "palés")
	eva := trabajadorDePrueba("Eva", horario, "caja")
	corto := trabajadorDePrueba("Pablo", tramoDePrueba(10, 0, 11, 0), "palés")

	reparto := RepartosDeTareas{
		plantilla: &Plantilla{
			trabajadores: []*Trabajador{ana, luis, eva, corto},
		},
		huecosPorDia: map[Dia][]Hueco{
			dia: {
				{tarea: "caja", tramo: horario, cubiertoPor: ana},
				{tarea: "palés", tramo: horario},
				{tarea: "inventario", tramo: horario},
			},
		},
	}

	casos := []struct {
		indice    int
		esperados []*Trabajador
		err       error
	}{
		{1, []*Trabajador{luis}, nil},
		{2, nil, nil},
		{0, nil, ErrHuecoCubierto},
		{3, nil, ErrHuecoNoExiste},
	}

	for _, caso := range casos {
		candidatos, err := reparto.Candidatos(dia, caso.indice)

		if !errors.Is(err, caso.err) || !slices.Equal(candidatos, caso.esperados) {
			t.Errorf("hueco %d: candidatos o error incorrectos", caso.indice)
		}
	}
}
