package LoggerInteractor

import (
	"github.com/AntonParaskiv/cleanarch-logger/interfaces/LoggerRepositoryMock"
	"reflect"
	"testing"
)

func TestInteractor_AddRepository(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		repository Repository
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				repositories: map[string]Repository{
					"Repository1": &LoggerRepositoryMock.Mock{
						Name: "Repository1",
					},
				},
			},
			args: args{
				repository: &LoggerRepositoryMock.Mock{
					Name: "Repository2",
				},
			},
			want: &Interactor{
				repositories: map[string]Repository{
					"Repository1": &LoggerRepositoryMock.Mock{
						Name: "Repository1",
					},
					"Repository2": &LoggerRepositoryMock.Mock{
						Name: "Repository2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.AddRepository(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.AddRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_ClearRepositories(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				repositories: map[string]Repository{
					"Repository1": &LoggerRepositoryMock.Mock{
						Name: "Repository1",
					},
					"Repository2": &LoggerRepositoryMock.Mock{
						Name: "Repository2",
					},
				},
			},
			want: &Interactor{
				repositories: map[string]Repository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.ClearRepositories(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.ClearRepositories() = %v, want %v", got, tt.want)
			}
		})
	}
}
