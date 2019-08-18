package FileMock

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

func TestMock_Write(t *testing.T) {
	type fields struct {
		Fdescr            uintptr
		Name              string
		Buffer            []byte
		SimulateErrorFlag bool
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
		wantM   *Mock
	}{
		{
			name: "Success",
			args: args{
				b: []byte("TestMessage"),
			},
			wantN: len([]byte("TestMessage")),
			wantM: &Mock{
				buffer: []byte("TestMessage"),
			},
		},
		{
			name: "Error",
			fields: fields{
				Buffer:            []byte{},
				SimulateErrorFlag: true,
			},
			args: args{
				b: []byte("TestMessage"),
			},
			wantN:   0,
			wantErr: true,
			wantM: &Mock{
				buffer:            []byte{},
				simulateErrorFlag: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Fdescr:            tt.fields.Fdescr,
				Name:              tt.fields.Name,
				buffer:            tt.fields.Buffer,
				simulateErrorFlag: tt.fields.SimulateErrorFlag,
			}
			gotN, err := m.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mock.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Mock.Write() = %v, want %v", gotN, tt.wantN)
			}
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("Mock = %v, want %v", m, tt.wantM)
			}
		})
	}
}

func TestMock_Fd(t *testing.T) {
	type fields struct {
		Fdescr            uintptr
		Name              string
		Buffer            []byte
		SimulateErrorFlag bool
	}
	tests := []struct {
		name   string
		fields fields
		want   uintptr
	}{
		{
			name: "Success",
			fields: fields{
				Fdescr: 111,
			},
			want: 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Fdescr:            tt.fields.Fdescr,
				Name:              tt.fields.Name,
				buffer:            tt.fields.Buffer,
				simulateErrorFlag: tt.fields.SimulateErrorFlag,
			}
			if got := m.Fd(); got != tt.want {
				t.Errorf("Mock.Fd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMock_SimulateError(t *testing.T) {
	type fields struct {
		Fdescr            uintptr
		Name              string
		Buffer            []byte
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
				Fdescr:            tt.fields.Fdescr,
				Name:              tt.fields.Name,
				buffer:            tt.fields.Buffer,
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
		Fdescr            uintptr
		Name              string
		Buffer            []byte
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				Fdescr:            tt.fields.Fdescr,
				Name:              tt.fields.Name,
				buffer:            tt.fields.Buffer,
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
