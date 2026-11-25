package turnos

import (
	"testing"
	"time"
)

func TestDisponibilidadCompleta(t *testing.T) {
	hora := func(h, m int) time.Time {
		return time.Date(2026, 11, 24, h, m, 0, 0, time.UTC)
	}

	tramo := func(hi, mi, hf, mf int) TramoHorario {
		return TramoHorario{
			inicio: hora(hi, mi),
			fin:    hora(hf, mf),
		}
	}

	hueco := tramo(10, 0, 11, 0)

	casos := []struct {
		nombre         string
		disponibilidad []TramoHorario
		esperado       bool
	}{
		{
			"todo el horario",
			[]TramoHorario{tramo(9, 0, 12, 0)},
			true,
		},
		{
			"tramos consecutivos",
			[]TramoHorario{
				tramo(10, 0, 10, 30),
				tramo(10, 30, 11, 0),
			},
			true,
		},
		{
			"tramos solapados",
			[]TramoHorario{
				tramo(10, 20, 11, 0),
				tramo(10, 0, 10, 35),
			},
			true,
		},
		{
			"hay un hueco sin disponibilidad",
			[]TramoHorario{
				tramo(10, 0, 10, 30),
				tramo(10, 45, 11, 0),
			},
			false,
		},
		{
			"solo cubre una parte",
			[]TramoHorario{tramo(10, 0, 10, 30)},
			false,
		},
		{
			"sin disponibilidad",
			nil,
			false,
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			trabajador := Trabajador{
				disponibilidad: caso.disponibilidad,
			}

			obtenido := trabajador.disponibleDurante(hueco)

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
