package turnos

import (
	"errors"
	"slices"
)

var (
	ErrDiaSinReparto     = errors.New("no hay reparto para ese día")
	ErrHuecoNoExiste     = errors.New("el hueco no existe")
	ErrHuecoCubierto     = errors.New("el hueco ya está cubierto")
	ErrHorarioInvalido   = errors.New("el horario del hueco no es válido")
	ErrTrabajadorAjeno   = errors.New("el trabajador no pertenece a la plantilla")
	ErrSinDisponibilidad = errors.New("el trabajador no está disponible durante todo el hueco")
	ErrTareaNoPermitida  = errors.New("el trabajador no puede realizar esa tarea")
	ErrTrabajadorOcupado = errors.New("el trabajador ya tiene otra tarea en ese horario")
)

type RepartosDeTareas struct {
	plantilla    *Plantilla
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

func (r RepartosDeTareas) huecoPendiente(dia Dia, indice int) (*Hueco, error) {
	huecos, existe := r.huecosPorDia[dia]

	if !existe {
		return nil, ErrDiaSinReparto
	}

	if indice < 0 || indice >= len(huecos) {
		return nil, ErrHuecoNoExiste
	}

	hueco := &huecos[indice]

	if hueco.cubiertoPor != nil {
		return nil, ErrHuecoCubierto
	}

	if !hueco.tramo.valido() {
		return nil, ErrHorarioInvalido
	}

	return hueco, nil
}

func (r RepartosDeTareas) comprobarTrabajador(hueco Hueco, trabajador *Trabajador) error {
	if !trabajador.puedeRealizar(hueco.tarea) {
		return ErrTareaNoPermitida
	}

	if !trabajador.disponibleDurante(hueco.tramo) {
		return ErrSinDisponibilidad
	}

	if r.ocupadoDurante(trabajador, hueco.tramo) {
		return ErrTrabajadorOcupado
	}

	return nil
}

func (r *RepartosDeTareas) Asignar(dia Dia, indice int, trabajador *Trabajador) error {
	hueco, err := r.huecoPendiente(dia, indice)

	if err != nil {
		return err
	}

	if !r.plantilla.contiene(trabajador) {
		return ErrTrabajadorAjeno
	}

	if err := r.comprobarTrabajador(*hueco, trabajador); err != nil {
		return err
	}

	hueco.cubiertoPor = trabajador

	return nil
}

func (r RepartosDeTareas) Candidatos(dia Dia, indice int) ([]*Trabajador, error) {
	hueco, err := r.huecoPendiente(dia, indice)

	if err != nil {
		return nil, err
	}

	var candidatos []*Trabajador

	if r.plantilla == nil {
		return candidatos, nil
	}

	for _, trabajador := range r.plantilla.trabajadores {
		if r.comprobarTrabajador(*hueco, trabajador) == nil {
			candidatos = append(candidatos, trabajador)
		}
	}

	return candidatos, nil
}

func (r RepartosDeTareas) PendientesObligatorios(dia Dia) ([]int, error) {
	huecos, existe := r.huecosPorDia[dia]

	if !existe {
		return nil, ErrDiaSinReparto
	}

	var pendientes []int

	for indice, hueco := range huecos {
		if !hueco.aplazable && hueco.cubiertoPor == nil {
			pendientes = append(pendientes, indice)
		}
	}

	return pendientes, nil
}
