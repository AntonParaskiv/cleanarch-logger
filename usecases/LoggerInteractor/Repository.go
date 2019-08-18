package LoggerInteractor

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

func (i *Interactor) sendToRepositories(message string, level int, time time.Time) {
	for _, repository := range i.repositories {
		if err := i.sendToRepository(repository, message, level, time); err != nil {
			err = errors.Errorf("%s %s", repository.GetName(), err.Error())
			fmt.Println(err)
		}
	}
}

func (i *Interactor) sendToRepository(repository Repository, message string, level int, time time.Time) (err error) {
	err = repository.Log(level, time, message)
	return
}

func (i *Interactor) AddRepository(repository Repository) *Interactor {
	i.repositories[repository.GetName()] = repository
	return i
}

func (i *Interactor) ClearRepositories() *Interactor {
	i.repositories = make(map[string]Repository)
	return i
}

func (i *Interactor) setRepositories(repositories map[string]Repository) {
	i.repositories = repositories
}

func (i *Interactor) getRepositories() (repositories map[string]Repository) {
	repositories = i.repositories
	return
}
