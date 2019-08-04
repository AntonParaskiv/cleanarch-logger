package LoggerInteractorStd

import (
	"log"
	"os"
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
				log.New(os.Stdout, "", log.LstdFlags),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, wantI %v", gotI, tt.wantI)
			}
		})
	}
}

func TestInteractor_SetLogger(t *testing.T) {
	type fields struct {
		logger *log.Logger
	}
	type args struct {
		logger *log.Logger
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantI  *Interactor
	}{
		{
			name: "Success",
			args: args{
				logger: log.New(os.Stdout, "TestPrefix", log.Llongfile),
			},
			wantI: &Interactor{
				logger: log.New(os.Stdout, "TestPrefix", log.Llongfile),
			},
		},
		{
			name: "SuccessNil",
			args: args{
				logger: nil,
			},
			wantI: &Interactor{
				logger: log.New(os.Stdout, "", log.LstdFlags),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			if got := i.SetLogger(tt.args.logger); !reflect.DeepEqual(got, tt.wantI) {
				t.Errorf("Interactor.SetLogger() = %v, want %v", got, tt.wantI)
			}
		})
	}
}
