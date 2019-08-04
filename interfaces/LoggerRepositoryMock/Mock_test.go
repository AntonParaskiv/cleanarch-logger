package LoggerRepositoryMock

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantM *Mock
	}{
		{
			name:  "Success",
			wantM: new(Mock),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := New(); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("New() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestMock_SimulateError(t *testing.T) {
	type fields struct {
		name              string
		logLevel          int
		logTime           time.Time
		message           string
		simulateErrorFlag bool
	}
	tests := []struct {
		name   string
		fields fields
		wantM  *Mock
	}{
		{
			name: "Success",
			fields: fields{
				simulateErrorFlag: false,
			},
			wantM: &Mock{
				simulateErrorFlag: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Name:              tt.fields.name,
				LogLevel:          tt.fields.logLevel,
				LogTime:           tt.fields.logTime,
				Message:           tt.fields.message,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if got := m.SimulateError(); !reflect.DeepEqual(got, tt.wantM) {
				t.Errorf("Mock.SimulateError() = %v, want %v", got, tt.wantM)
			}
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("Mock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestMock_Error(t *testing.T) {
	type fields struct {
		name              string
		logLevel          int
		logTime           time.Time
		message           string
		simulateErrorFlag bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		wantM   *Mock
	}{
		{
			name: "Success",
			fields: fields{
				simulateErrorFlag: true,
			},
			wantM: &Mock{
				simulateErrorFlag: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Name:              tt.fields.name,
				LogLevel:          tt.fields.logLevel,
				LogTime:           tt.fields.logTime,
				Message:           tt.fields.message,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := m.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Mock.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("Mock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestMock_GetName(t *testing.T) {
	type fields struct {
		name              string
		logLevel          int
		logTime           time.Time
		message           string
		simulateErrorFlag bool
	}
	tests := []struct {
		name     string
		fields   fields
		wantName string
	}{
		{
			name: "Success",
			fields: fields{
				name: "TestName",
			},
			wantName: "TestName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Name:              tt.fields.name,
				LogLevel:          tt.fields.logLevel,
				LogTime:           tt.fields.logTime,
				Message:           tt.fields.message,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if gotName := m.GetName(); gotName != tt.wantName {
				t.Errorf("Mock.GetName() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
