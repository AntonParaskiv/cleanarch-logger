package LoggerRepositoryJson

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"github.com/AntonParaskiv/cleanarch-logger/infrastructure/LoggerStorageMock"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		storage Storage
	}
	tests := []struct {
		name  string
		args  args
		wantR *Repository
	}{
		{
			name: "Success",
			args: args{
				storage: LoggerStorageMock.New(),
			},
			wantR: &Repository{
				storage:    LoggerStorageMock.New(),
				logLevel:   LoggerDomain.LogLevelNone,
				timeFormat: time.Stamp,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := New(tt.args.storage); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("New() = %v, wantR %v", gotR, tt.wantR)
			}
		})
	}
}

func TestRepository_SetName(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantR  *Repository
	}{
		{
			name: "Success",
			args: args{
				name: "TestName",
			},
			wantR: &Repository{
				name: "TestName",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				name:       tt.fields.name,
				storage:    tt.fields.storage,
				logLevel:   tt.fields.logLevel,
				prefix:     tt.fields.prefix,
				timeFormat: tt.fields.timeFormat,
			}
			if got := r.SetName(tt.args.name); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.SetName() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_SetPrefix(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantR  *Repository
	}{
		{
			name: "Success",
			args: args{
				prefix: "TestPrefix",
			},
			wantR: &Repository{
				prefix: "TestPrefix",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				name:       tt.fields.name,
				storage:    tt.fields.storage,
				logLevel:   tt.fields.logLevel,
				prefix:     tt.fields.prefix,
				timeFormat: tt.fields.timeFormat,
			}
			if got := r.SetPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.SetPrefix() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_SetTimeFormat(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	type args struct {
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantR  *Repository
	}{
		{
			name: "Success",
			args: args{
				timeFormat: time.Stamp,
			},
			wantR: &Repository{
				timeFormat: time.Stamp,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				name:       tt.fields.name,
				storage:    tt.fields.storage,
				logLevel:   tt.fields.logLevel,
				prefix:     tt.fields.prefix,
				timeFormat: tt.fields.timeFormat,
			}
			if got := r.SetTimeFormat(tt.args.timeFormat); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.SetTimeFormat() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_GetName(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
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
			r := &Repository{
				name:       tt.fields.name,
				storage:    tt.fields.storage,
				logLevel:   tt.fields.logLevel,
				prefix:     tt.fields.prefix,
				timeFormat: tt.fields.timeFormat,
			}
			if gotName := r.GetName(); gotName != tt.wantName {
				t.Errorf("Repository.GetName() = %v, wantR %v", gotName, tt.wantName)
			}
		})
	}
}
