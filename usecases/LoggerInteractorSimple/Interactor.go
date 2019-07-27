package LoggerInteractorSimple

type LoggerInteractor struct {
}

func New() (i *LoggerInteractor) {
	i = new(LoggerInteractor)
	return
}
