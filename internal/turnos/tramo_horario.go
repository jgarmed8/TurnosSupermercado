package turnos

import "time"

type TramoHorario struct {
	inicio time.Time
	fin    time.Time
}

func (t TramoHorario) valido() bool {
	return t.inicio.Before(t.fin)
}
