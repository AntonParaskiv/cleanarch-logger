package LoggerInteractorSimple

type Interactor struct {
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}
