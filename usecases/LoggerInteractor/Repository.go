package LoggerInteractor

func (i *Interactor) AddRepository(repository Repository) *Interactor {
	i.repositories[repository.GetName()] = repository
	return i
}

func (i *Interactor) ClearRepositories() *Interactor {
	i.repositories = make(map[string]Repository)
	return i
}
