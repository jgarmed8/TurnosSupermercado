package turnos

type Trabajador struct { 
	nombre           string 
	disponibilidad   []TramoHorario
	tareasPermitidas map[Tarea]struct{}
}
