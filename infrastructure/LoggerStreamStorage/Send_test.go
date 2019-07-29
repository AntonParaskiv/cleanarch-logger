package LoggerStreamStorage

import (
	"github.com/AntonParaskiv/cleanarch-logger/infrastructure/FileMock"
	"reflect"
	"testing"
)

func TestStorage_Send(t *testing.T) {
	type fields struct {
		stream File
	}
	type args struct {
		message string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantStorage *Storage
	}{
		{
			name: "Success",
			fields: fields{
				stream: FileMock.New(),
			},
			args: args{
				message: "TestMessage",
			},
			wantStorage: &Storage{
				stream: &FileMock.Mock{
					Buffer: []byte("TestMessage"),
				},
			},
		},
		{
			name: "Err",
			fields: fields{
				stream: FileMock.New().SimulateError(),
			},
			args: args{
				message: "TestMessage",
			},
			wantErr: true,
			wantStorage: &Storage{
				stream: &FileMock.Mock{},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				stream: tt.fields.stream,
			}
			if err := s.Send(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(s, tt.wantStorage) {
				t.Errorf("Storage = %v, want %v", s, tt.wantStorage)
			}

		})
	}
}
