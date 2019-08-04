package LoggerStorageMock

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantM *Mock
	}{
		{
			name:  "Success",
			wantM: &Mock{},
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

func TestMock_Send(t *testing.T) {
	type fields struct {
		Message           string
		SimulateErrorFlag bool
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantM   *Mock
	}{
		{
			name: "Success",
			args: args{
				message: "TestMessage",
			},
			wantM: &Mock{
				Message: "TestMessage",
			},
		},
		{
			name: "Error",
			fields: fields{
				Message:           "",
				SimulateErrorFlag: true,
			},
			args: args{
				message: "",
			},
			wantErr: true,
			wantM: &Mock{
				Message:           "",
				simulateErrorFlag: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Message:           tt.fields.Message,
				simulateErrorFlag: tt.fields.SimulateErrorFlag,
			}
			if err := m.Send(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Mock.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("Mock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestMock_SimulateError(t *testing.T) {
	type fields struct {
		Message           string
		SimulateErrorFlag bool
	}
	tests := []struct {
		name   string
		fields fields
		wantM  *Mock
	}{
		{
			name: "Success",
			fields: fields{
				SimulateErrorFlag: false,
			},
			wantM: &Mock{
				simulateErrorFlag: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Message:           tt.fields.Message,
				simulateErrorFlag: tt.fields.SimulateErrorFlag,
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
		Message           string
		SimulateErrorFlag bool
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
				SimulateErrorFlag: true,
			},
			wantM: &Mock{
				simulateErrorFlag: false,
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Message:           tt.fields.Message,
				simulateErrorFlag: tt.fields.SimulateErrorFlag,
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
