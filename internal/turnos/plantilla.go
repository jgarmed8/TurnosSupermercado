package turnos

import "slices"

type Plantilla struct {
	trabajadores []*Trabajador
}

func (p *Plantilla) contiene(trabajador *Trabajador) bool {
	return p != nil && trabajador != nil && slices.Contains(p.trabajadores, trabajador)
}

func (p *Plantilla) agregar(trabajador *Trabajador) bool {
	if trabajador == nil || p.contiene(trabajador) {
		return false
	}

	p.trabajadores = append(p.trabajadores, trabajador)
	return true
}
