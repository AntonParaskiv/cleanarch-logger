package LoggerRepositoryJson

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"github.com/AntonParaskiv/cleanarch-logger/infrastructure/LoggerStorageMock"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"testing"
	"time"
)

func TestRepository_Log(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	type args struct {
		level   int
		time    time.Time
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantR   *Repository
	}{
		{
			name: "SuccessSimple",
			fields: fields{
				storage:    LoggerStorageMock.New(),
				logLevel:   LoggerDomain.LogLevelAllBits,
				prefix:     "",
				timeFormat: "",
			},
			args: args{
				level:   LoggerDomain.LogLevelNone,
				time:    time.Now(),
				message: "TestMessage",
			},
			wantR: &Repository{
				storage: &LoggerStorageMock.Mock{
					Message: `{"message":"TestMessage"}` + "\n",
				},
				logLevel: LoggerDomain.LogLevelAllBits,
			},
		},
		{
			name: "SuccessWithGarbage",
			fields: fields{
				storage:    LoggerStorageMock.New(),
				logLevel:   LoggerDomain.LogLevelAllBits,
				prefix:     "",
				timeFormat: "",
			},
			args: args{
				level:   LoggerDomain.LogLevelNone,
				time:    time.Now(),
				message: " \t\r\nTestMessage \t\r\n",
			},
			wantR: &Repository{
				storage: &LoggerStorageMock.Mock{
					Message: `{"message":"TestMessage"}` + "\n",
				},
				logLevel: LoggerDomain.LogLevelAllBits,
			},
		},
		{
			name: "SuccessWithPrefix",
			fields: fields{
				storage:    LoggerStorageMock.New(),
				logLevel:   LoggerDomain.LogLevelAllBits,
				prefix:     "WithPrefix",
				timeFormat: "",
			},
			args: args{
				level:   LoggerDomain.LogLevelNone,
				time:    time.Now(),
				message: "TestMessage",
			},
			wantR: &Repository{
				storage: &LoggerStorageMock.Mock{
					Message: `{"message":"WithPrefix TestMessage"}` + "\n",
				},
				logLevel: LoggerDomain.LogLevelAllBits,
				prefix:   "WithPrefix",
			},
		},
		{
			name: "SuccessWithLogLevelTitle",
			fields: fields{
				storage:    LoggerStorageMock.New(),
				logLevel:   LoggerDomain.LogLevelAllBits,
				prefix:     "",
				timeFormat: "",
			},
			args: args{
				level:   LoggerDomain.LogLevelInfo,
				time:    time.Now(),
				message: "TestMessage",
			},
			wantR: &Repository{
				storage: &LoggerStorageMock.Mock{
					Message: `{"level":"` + logLevelTitle[LoggerDomain.LogLevelInfo] + `","message":"TestMessage"}` + "\n",
				},
				logLevel: LoggerDomain.LogLevelAllBits,
			},
		},
		{
			name: "SuccessWithTime",
			fields: fields{
				storage:    LoggerStorageMock.New(),
				logLevel:   LoggerDomain.LogLevelAllBits,
				prefix:     "",
				timeFormat: time.RFC850,
			},
			args: args{
				time:    time.Date(2019, 07, 30, 17, 23, 46, 0, time.UTC),
				message: "TestMessage",
			},
			wantR: &Repository{
				storage: &LoggerStorageMock.Mock{
					Message: `{"timestamp":"Tuesday, 30-Jul-19 17:23:46 UTC","message":"TestMessage"}` + "\n",
				},
				logLevel:   LoggerDomain.LogLevelAllBits,
				timeFormat: time.RFC850,
			},
		},
		{
			name: "SuccessWithAll",
			fields: fields{
				storage:    LoggerStorageMock.New(),
				logLevel:   LoggerDomain.LogLevelAllBits,
				prefix:     "WithPrefix",
				timeFormat: time.RFC850,
			},
			args: args{
				level:   LoggerDomain.LogLevelInfo,
				time:    time.Date(2019, 07, 30, 17, 23, 46, 0, time.UTC),
				message: " \t\r\nTestMessage \t\r\n",
			},
			wantR: &Repository{
				storage: &LoggerStorageMock.Mock{
					Message: `{"level":"` + logLevelTitle[LoggerDomain.LogLevelInfo] + `","timestamp":"Tuesday, 30-Jul-19 17:23:46 UTC","message":"WithPrefix TestMessage"}` + "\n",
				},
				logLevel:   LoggerDomain.LogLevelAllBits,
				prefix:     "WithPrefix",
				timeFormat: time.RFC850,
			},
		},
		{
			name: "SuccessNoLogByLogLevel",
			fields: fields{
				storage:    LoggerStorageMock.New(),
				logLevel:   LoggerDomain.LogLevelError,
				prefix:     "",
				timeFormat: "",
			},
			args: args{
				level:   LoggerDomain.LogLevelInfo,
				time:    time.Now(),
				message: "TestMessage",
			},
			wantR: &Repository{
				storage: &LoggerStorageMock.Mock{
					Message: "",
				},
				logLevel: LoggerDomain.LogLevelError,
			},
		},
		{
			name: "Err",
			fields: fields{
				storage:    LoggerStorageMock.New().SimulateError(),
				logLevel:   LoggerDomain.LogLevelAllBits,
				prefix:     "",
				timeFormat: "",
			},
			args: args{
				level:   LoggerDomain.LogLevelInfo,
				time:    time.Now(),
				message: "TestMessage",
			},
			wantR: &Repository{
				storage: &LoggerStorageMock.Mock{
					Message: "",
				},
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			wantErr: true,
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
			if err := r.Log(tt.args.level, tt.args.time, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Log() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.wantR) {
				t.Errorf("Repository = %v, want %v", r, tt.wantR)
				spew.Dump(r)
			}
		})
	}
}
