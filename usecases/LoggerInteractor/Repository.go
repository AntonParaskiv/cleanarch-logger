package LoggerInteractor

func (i *LoggerInteractor) AddRepository(repository LoggerRepository) *LoggerInteractor {
	i.repositories[repository.GetName()] = repository
	return i
}

func (i *LoggerInteractor) ClearRepositories() *LoggerInteractor {
	i.repositories = make(map[string]LoggerRepository)
	return i
}
