package LoggerInteractorStd

import (
	"bytes"
	"log"
	"testing"
)

func TestInteractor_log(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				message: "TestMessage",
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.log(tt.args.message)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_logf(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.logf(tt.args.format, tt.args.a...)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Debug(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				message: "TestMessage",
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Debug(tt.args.message)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Debugf(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Debugf(tt.args.format, tt.args.a...)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Info(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				message: "TestMessage",
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Info(tt.args.message)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Infof(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Infof(tt.args.format, tt.args.a...)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Warn(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				message: "TestMessage",
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Warn(tt.args.message)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Warnf(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Warnf(tt.args.format, tt.args.a...)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Error(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				message: "TestMessage",
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Error(tt.args.message)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Errorf(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Errorf(tt.args.format, tt.args.a...)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Fatal(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				message: "TestMessage",
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Fatal(tt.args.message)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}

func TestInteractor_Fatalf(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	type fields struct {
		logger *log.Logger
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBuf string
	}{
		{
			name: "Success",
			fields: fields{
				logger: log.New(buf, "", 0),
			},
			args: args{
				format: "Test%s",
				a: []interface{}{
					"Message",
				},
			},
			wantBuf: "TestMessage\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				logger: tt.fields.logger,
			}
			i.Fatalf(tt.args.format, tt.args.a...)
		})
		if buf.String() != tt.wantBuf {
			t.Errorf("Buf = %v, want %v", buf.String(), tt.wantBuf)
		}
	}
}
