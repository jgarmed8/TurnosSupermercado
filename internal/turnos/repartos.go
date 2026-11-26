package turnos

import "slices"

type RepartosDeTareas struct {
	huecosPorDia map[Dia][]Hueco
}

func (r RepartosDeTareas) ocupadoDurante(trabajador *Trabajador, tramo TramoHorario) bool {
	for _, huecos := range r.huecosPorDia {
		if slices.ContainsFunc(huecos, func(h Hueco) bool {
			return h.cubiertoPor == trabajador && h.tramo.solapa(tramo)
		}) {
			return true
		}
	}

	return false
}
