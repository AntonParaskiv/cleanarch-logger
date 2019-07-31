package LoggerRepositoryMock

import (
	"reflect"
	"testing"
	"time"
)

func TestMock_Log(t *testing.T) {
	type fields struct {
		name              string
		logLevel          int
		logTime           time.Time
		message           string
		simulateErrorFlag bool
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
		wantM   *Mock
	}{
		{
			name: "Success",
			args: args{
				level:   100,
				time:    time.Date(2019, 07, 30, 17, 23, 46, 0, time.UTC),
				message: "TestMessage",
			},
			wantM: &Mock{
				logLevel: 100,
				logTime:  time.Date(2019, 07, 30, 17, 23, 46, 0, time.UTC),
				message:  "TestMessage",
			},
		},
		{
			name: "Error",
			fields: fields{
				simulateErrorFlag: true,
			},
			args: args{
				level:   100,
				time:    time.Date(2019, 07, 30, 17, 23, 46, 0, time.UTC),
				message: "TestMessage",
			},
			wantErr: true,
			wantM:   &Mock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mock{
				name:              tt.fields.name,
				logLevel:          tt.fields.logLevel,
				logTime:           tt.fields.logTime,
				message:           tt.fields.message,
				simulateErrorFlag: tt.fields.simulateErrorFlag,
			}
			if err := m.Log(tt.args.level, tt.args.time, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Mock.Log() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(m, tt.wantM) {
				t.Errorf("Mock = %v, want %v", m, tt.wantM)
			}
		})
	}
}
