package turnos

import "sort"

type Trabajador struct {
	nombre           string
	disponibilidad   []TramoHorario
	tareasPermitidas map[Tarea]struct{}
}

func (t *Trabajador) disponibleDurante(hueco TramoHorario) bool {
	if t == nil || !hueco.valido() {
		return false
	}

	tramos := append([]TramoHorario(nil), t.disponibilidad...)

	sort.Slice(tramos, func(i, j int) bool {
		return tramos[i].inicio.Before(tramos[j].inicio)
	})

	cubiertoHasta := hueco.inicio

	for _, tramo := range tramos {
		if !tramo.valido() || !tramo.fin.After(cubiertoHasta) {
			continue
		}

		if tramo.inicio.After(cubiertoHasta) {
			return false
		}

		cubiertoHasta = tramo.fin

		if !cubiertoHasta.Before(hueco.fin) {
			return true
		}
	}

	return false
}

func (t *Trabajador) puedeRealizar(tarea Tarea) bool {
	if t == nil {
		return false
	}

	_, permitida := t.tareasPermitidas[tarea]
	return permitida
}
