package LoggerStreamStorage

import (
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		stream *os.File
	}
	tests := []struct {
		name  string
		args  args
		wantS *Storage
	}{
		{
			name: "Success",
			args: args{
				stream: os.Stdout,
			},
			wantS: &Storage{
				stream: os.Stdout,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := New(tt.args.stream); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("New() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestNewFromFileName(t *testing.T) {
	type args struct {
		fileName string
		perm     os.FileMode
	}
	tests := []struct {
		name                string
		args                args
		wantStorageFileName string
		wantErr             bool
		wantStorage         *Storage
	}{
		{
			name: "Success",
			args: args{
				fileName: "/tmp/LoggerStreamStorage_TestNewFromFileName.log",
				perm:     0777,
			},
			wantStorageFileName: "/tmp/LoggerStreamStorage_TestNewFromFileName.log",
		},
		{
			name: "Err",
			args: args{
				fileName: "/",
				perm:     0777,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStorage, err := NewFromFileName(tt.args.fileName, tt.args.perm)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				tt.wantStorage = &Storage{
					stream: os.NewFile(gotStorage.stream.Fd(), tt.wantStorageFileName),
				}
			}
			if !reflect.DeepEqual(gotStorage, tt.wantStorage) {
				t.Errorf("NewFromFileName() = %v, want %v", gotStorage, tt.wantStorage)
			}
		})
	}
}
