package turnos

import (
	"testing"
	"time"
)

func tramoDePrueba(hi, mi, hf, mf int) TramoHorario {
	return TramoHorario{
		inicio: time.Date(2026, 11, 24, hi, mi, 0, 0, time.UTC),
		fin:    time.Date(2026, 11, 24, hf, mf, 0, 0, time.UTC),
	}
}

func TestOcupadoDurante(t *testing.T) {
	ana := &Trabajador{nombre: "Ana"}
	luis := &Trabajador{nombre: "Luis"}
	dia := Dia{2026, time.November, 24}

	reparto := RepartosDeTareas{
		huecosPorDia: map[Dia][]Hueco{
			dia: {
				{
					tramo:       tramoDePrueba(10, 0, 11, 0),
					cubiertoPor: ana,
				},
			},
		},
	}

	casos := []struct {
		nombre     string
		trabajador *Trabajador
		tramo      TramoHorario
		esperado   bool
	}{
		{"se solapa", ana, tramoDePrueba(10, 30, 12, 0), true},
		{"mismo horario", ana, tramoDePrueba(10, 0, 11, 0), true},
		{"horario contenido", ana, tramoDePrueba(10, 15, 10, 45), true},
		{"empieza al terminar", ana, tramoDePrueba(11, 0, 12, 0), false},
		{"termina al empezar", ana, tramoDePrueba(9, 0, 10, 0), false},
		{"otro trabajador", luis, tramoDePrueba(10, 0, 11, 0), false},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			obtenido := reparto.ocupadoDurante(caso.trabajador, caso.tramo)

			if obtenido != caso.esperado {
				t.Errorf("esperado %v, obtenido %v", caso.esperado, obtenido)
			}
		})
	}
}
