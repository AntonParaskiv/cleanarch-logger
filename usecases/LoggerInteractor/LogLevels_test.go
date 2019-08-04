package LoggerInteractor

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"reflect"
	"testing"
)

func TestInteractor_SetLogLevelNone(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.SetLogLevelNone().SetLogLevelNone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.SetLogLevelNone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_AddLogLevelDebug(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelDebug,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.AddLogLevelDebug().AddLogLevelDebug(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.AddLogLevelDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_AddLogLevelInfo(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelInfo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.AddLogLevelInfo().AddLogLevelInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.AddLogLevelInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_AddLogLevelWarn(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelWarn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.AddLogLevelWarn().AddLogLevelWarn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.AddLogLevelWarn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_AddLogLevelError(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.AddLogLevelError().AddLogLevelError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.AddLogLevelError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_AddLogLevelFatal(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelFatal,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.AddLogLevelFatal().AddLogLevelFatal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.AddLogLevelFatal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_RemoveLogLevelDebug(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelDebug,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.RemoveLogLevelDebug().RemoveLogLevelDebug(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.RemoveLogLevelDebug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_RemoveLogLevelInfo(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelInfo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.RemoveLogLevelInfo().RemoveLogLevelInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.RemoveLogLevelInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_RemoveLogLevelWarn(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelWarn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.RemoveLogLevelWarn().RemoveLogLevelWarn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.RemoveLogLevelWarn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_RemoveLogLevelError(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.RemoveLogLevelError().RemoveLogLevelError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.RemoveLogLevelError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_RemoveLogLevelFatal(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelFatal,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.RemoveLogLevelFatal().RemoveLogLevelFatal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.RemoveLogLevelFatal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetLogLevelAll(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone +
					LoggerDomain.LogLevelDebug +
					LoggerDomain.LogLevelInfo +
					LoggerDomain.LogLevelWarn +
					LoggerDomain.LogLevelError +
					LoggerDomain.LogLevelFatal,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.SetLogLevelAll().SetLogLevelAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.SetLogLevelAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetLogLevelDebugInfoWarn(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone +
					LoggerDomain.LogLevelDebug +
					LoggerDomain.LogLevelInfo +
					LoggerDomain.LogLevelWarn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.SetLogLevelDebugInfoWarn().SetLogLevelDebugInfoWarn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.SetLogLevelDebugInfoWarn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetLogLevelInfoWarn(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone +
					LoggerDomain.LogLevelInfo +
					LoggerDomain.LogLevelWarn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.SetLogLevelInfoWarn().SetLogLevelInfoWarn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.SetLogLevelInfoWarn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetLogLevelErrorFatal(t *testing.T) {
	type fields struct {
		repositories map[string]Repository
		logLevel     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Interactor
	}{
		{
			name: "Success",
			want: &Interactor{
				logLevel: LoggerDomain.LogLevelNone +
					LoggerDomain.LogLevelError +
					LoggerDomain.LogLevelFatal,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				repositories: tt.fields.repositories,
				logLevel:     tt.fields.logLevel,
			}
			if got := i.SetLogLevelErrorFatal().SetLogLevelErrorFatal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.SetLogLevelErrorFatal() = %v, want %v", got, tt.want)
			}
		})
	}
}
