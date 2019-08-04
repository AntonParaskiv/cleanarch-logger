package LoggerInteractor

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"github.com/AntonParaskiv/cleanarch-logger/interfaces/LoggerRepositoryMock"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantI *Interactor
	}{
		{
			name: "Success",
			wantI: &Interactor{
				repositories: map[string]Repository{},
				logLevel:     LoggerDomain.LogLevelNone,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestInteractor_Fork(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name                 string
		fields               fields
		wantForkedInteractor *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				repositories: map[string]Repository{
					"repository1": LoggerRepositoryMock.New(),
					"repository2": LoggerRepositoryMock.New(),
				},
				logLevel: LoggerDomain.LogLevelDebug,
			},
			wantForkedInteractor: &Interactor{
				repositories: map[string]Repository{
					"repository1": LoggerRepositoryMock.New(),
					"repository2": LoggerRepositoryMock.New(),
				},
				logLevel: LoggerDomain.LogLevelDebug,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			gotForkedInteractor := i.Fork()
			if !reflect.DeepEqual(gotForkedInteractor, tt.wantForkedInteractor) {
				t.Errorf("Interactor.Fork() = %v, want %v", gotForkedInteractor, tt.wantForkedInteractor)
			}

			// must to point to different struct
			if gotForkedInteractor == i {
				t.Errorf("Interactor.Fork() must return pointer to new cloned struct")
			}

		})
	}
}
