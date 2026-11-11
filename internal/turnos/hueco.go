package turnos

type Tarea string

type Hueco struct {
	tarea       Tarea
	tramo       TramoHorario
	aplazable   bool
	cubiertoPor *Trabajador
}
