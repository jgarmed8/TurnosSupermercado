package turnos

import (
	"slices"
	"testing"
	"time"
)

func TestDisponibilidadCompleta(t *testing.T) {

	hueco := tramoDePrueba(10, 0, 11, 0)

	casos := []struct {
		nombre         string
		disponibilidad []TramoHorario
		esperado       bool
	}{
		{
			"todo el horario",
			[]TramoHorario{tramoDePrueba(9, 0, 12, 0)},
			true,
		},
		{
			"horario exacto",
			[]TramoHorario{tramoDePrueba(10, 0, 11, 0)},
			true,
		},
		{
			"tramos repetidos",
			[]TramoHorario{
				tramoDePrueba(10, 0, 11, 0),
				tramoDePrueba(10, 0, 11, 0),
			},
			true,
		},
		{
			"tramos consecutivos",
			[]TramoHorario{
				tramoDePrueba(10, 0, 10, 30),
				tramoDePrueba(10, 30, 11, 0),
			},
			true,
		},
		{
			"tramos solapados",
			[]TramoHorario{
				tramoDePrueba(10, 20, 11, 0),
				tramoDePrueba(10, 0, 10, 35),
			},
			true,
		},
		{
			"hay un hueco sin disponibilidad",
			[]TramoHorario{
				tramoDePrueba(10, 0, 10, 30),
				tramoDePrueba(10, 45, 11, 0),
			},
			false,
		},
		{
			"solo cubre una parte",
			[]TramoHorario{tramoDePrueba(10, 0, 10, 30)},
			false,
		},
		{
			"sin disponibilidad",
			nil,
			false,
		},
		{
			"disponible otro dia",
			[]TramoHorario{
				{
					inicio: time.Date(2026, 11, 25, 10, 0, 0, 0, time.UTC),
					fin:    time.Date(2026, 11, 25, 11, 0, 0, 0, time.UTC),
				},
			},
			false,
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			trabajador := Trabajador{
				disponibilidad: caso.disponibilidad,
			}

			original := append(
				[]TramoHorario(nil),
				trabajador.disponibilidad...,
			)

			obtenido := trabajador.disponibleDurante(hueco)

			if obtenido != caso.esperado {
				t.Errorf(
					"esperado %v, obtenido %v",
					caso.esperado,
					obtenido,
				)
			}

			if !slices.Equal(trabajador.disponibilidad, original) {
				t.Error("se ha modificado la disponibilidad del trabajador")
			}
		})
	}
}

func TestDisponibilidadHuecoInvalido(t *testing.T) {
	inicio := time.Date(2026, 11, 24, 10, 0, 0, 0, time.UTC)

	trabajador := Trabajador{
		disponibilidad: []TramoHorario{
			{
				inicio: inicio,
				fin:    inicio.Add(2 * time.Hour),
			},
		},
	}

	casos := []struct {
		nombre string
		hueco  TramoHorario
	}{
		{
			"sin duracion",
			TramoHorario{
				inicio: inicio,
				fin:    inicio,
			},
		},
		{
			"fin anterior al inicio",
			TramoHorario{
				inicio: inicio.Add(time.Hour),
				fin:    inicio,
			},
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			if trabajador.disponibleDurante(caso.hueco) {
				t.Error("un hueco con horario incorrecto no puede cubrirse")
			}
		})
	}
}

func TestTareaPermitida(t *testing.T) {
	trabajador := &Trabajador{
		tareasPermitidas: map[Tarea]struct{}{
			"caja":       {},
			"inventario": {},
		},
	}

	casos := []struct {
		nombre     string
		trabajador *Trabajador
		tarea      Tarea
		esperado   bool
	}{
		{
			"puede atender caja",
			trabajador,
			"caja",
			true,
		},
		{
			"puede hacer inventario",
			trabajador,
			"inventario",
			true,
		},
		{
			"no puede recoger pales",
			trabajador,
			"recogida de palés",
			false,
		},
		{
			"sin tareas permitidas",
			&Trabajador{},
			"caja",
			false,
		},
		{
			"sin trabajador",
			nil,
			"caja",
			false,
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			obtenido := caso.trabajador.puedeRealizar(caso.tarea)

			if obtenido != caso.esperado {
				t.Errorf(
					"esperado %v, obtenido %v",
					caso.esperado,
					obtenido,
				)
			}
		})
	}
}
