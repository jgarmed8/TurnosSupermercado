package turnos

type TipoTarea string

const (
	Caja          TipoTarea = "caja"
	RecogidaPales TipoTarea = "recogida_pales"
	Reposicion    TipoTarea = "reposicion"
)

type Tarea struct {
	tipo      TipoTarea
	intervalo Intervalo
	aplazable bool
}
