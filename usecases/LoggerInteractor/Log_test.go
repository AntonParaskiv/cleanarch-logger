package LoggerInteractor

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"github.com/AntonParaskiv/cleanarch-logger/interfaces/LoggerRepositoryMock"
	"github.com/pkg/errors"
	"testing"
	"time"
)

func compareRepositoryMock(mocks map[string]*LoggerRepositoryMock.Mock, wantMocks map[string]*LoggerRepositoryMock.Mock, timeStart time.Time, timeStop time.Time) (err error) {
	if len(mocks) != len(wantMocks) {
		err = errors.Errorf("len repositories = %v, want %v", len(mocks), len(wantMocks))
		return
	}

	for key, wantMock := range wantMocks {

		mock, ok := mocks[key]
		if !ok {
			err = errors.Errorf("wantMock key = %v, not exist in repositories", key)
			return
		}

		if mock.LogLevel != wantMock.LogLevel {
			err = errors.Errorf("LogLevel = %v, want %v", mock.LogLevel, wantMock.LogLevel)
			return
		}
		if mock.Message != wantMock.Message {
			err = errors.Errorf("Message = %v, want %v", mock.Message, wantMock.Message)
			return
		}
		if mock.LogTime != (time.Time{}) {
			if mock.LogTime.Before(timeStart) || mock.LogTime.After(timeStop) {
				err = errors.Errorf("LogTime = %v is not match interval, want %v - %v", mock.LogTime.String(), timeStart.String(), timeStop.String())
				return
			}
		}
	}

	return
}

func TestInteractor_log(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		level   int
		message string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "SuccessNoLogByLogLevel",
			fields: fields{
				logLevel: LoggerDomain.LogLevelFatal,
			},
			args: args{
				level:   LoggerDomain.LogLevelDebug,
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
		},
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				level:   LoggerDomain.LogLevelAllBits,
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelAllBits,
					Message:  "TestMessage",
				},
			},
		},
		{
			name: "Error",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				level:   LoggerDomain.LogLevelAllBits,
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New().SimulateError(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
		},
		{
			name: "SuccessMultipleRepositories",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				level:   LoggerDomain.LogLevelAllBits,
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
				"repository2": LoggerRepositoryMock.New().SimulateError(),
				"repository3": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelAllBits,
					Message:  "TestMessage",
				},
				"repository2": LoggerRepositoryMock.New(),
				"repository3": {
					LogLevel: LoggerDomain.LogLevelAllBits,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.log(tt.args.level, tt.args.message)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_logf(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		level  int
		format string
		a      []interface{}
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				level:  LoggerDomain.LogLevelAllBits,
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelAllBits,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.logf(tt.args.level, tt.args.format, tt.args.a...)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Debug(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelDebug,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Debug(tt.args.message)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Debugf(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelDebug,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Debugf(tt.args.format, tt.args.a...)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Info(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelInfo,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Info(tt.args.message)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Infof(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelInfo,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Infof(tt.args.format, tt.args.a...)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Warn(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelWarn,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Warn(tt.args.message)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Warnf(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelWarn,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Warnf(tt.args.format, tt.args.a...)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Error(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelError,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Error(tt.args.message)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Errorf(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelError,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Errorf(tt.args.format, tt.args.a...)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Fatal(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				message: "TestMessage",
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelFatal,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Fatal(tt.args.message)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestInteractor_Fatalf(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		repositoryMocks     map[string]*LoggerRepositoryMock.Mock
		wantRepositoryMocks map[string]*LoggerRepositoryMock.Mock
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			repositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": LoggerRepositoryMock.New(),
			},
			wantRepositoryMocks: map[string]*LoggerRepositoryMock.Mock{
				"repository1": {
					LogLevel: LoggerDomain.LogLevelFatal,
					Message:  "TestMessage",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: map[string]Repository{},
				logLevel:     tt.fields.logLevel,
			}

			for key, mock := range tt.repositoryMocks {
				i.repositories[key] = mock
			}

			timeStart := time.Now()
			i.Fatalf(tt.args.format, tt.args.a...)
			timeStop := time.Now()

			if err := compareRepositoryMock(tt.repositoryMocks, tt.wantRepositoryMocks, timeStart, timeStop); err != nil {
				t.Error(err.Error())
			}
		})
	}
}
