package LoggerRepository

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"reflect"
	"testing"
)

func TestRepository_SetLogLevelNone(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone,
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
			if got := r.SetLogLevelNone().SetLogLevelNone(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.SetLogLevelNone() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_AddLogLevelDebug(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelDebug,
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

			if got := r.AddLogLevelDebug().AddLogLevelDebug(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.AddLogLevelDebug() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_AddLogLevelInfo(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelInfo,
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
			if got := r.AddLogLevelInfo().AddLogLevelInfo(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.AddLogLevelInfo() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_AddLogLevelWarn(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelWarn,
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
			if got := r.AddLogLevelWarn().AddLogLevelWarn(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.AddLogLevelWarn() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_AddLogLevelError(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelError,
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
			if got := r.AddLogLevelError().AddLogLevelError(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.AddLogLevelError() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_AddLogLevelFatal(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelNone,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone + LoggerDomain.LogLevelFatal,
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
			if got := r.AddLogLevelFatal().AddLogLevelFatal(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.AddLogLevelFatal() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_RemoveLogLevelDebug(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelDebug,
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
			if got := r.RemoveLogLevelDebug().RemoveLogLevelDebug(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.RemoveLogLevelDebug() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_RemoveLogLevelInfo(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelInfo,
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
			if got := r.RemoveLogLevelInfo().RemoveLogLevelInfo(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.RemoveLogLevelInfo() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_RemoveLogLevelWarn(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelWarn,
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
			if got := r.RemoveLogLevelWarn().RemoveLogLevelWarn(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.RemoveLogLevelWarn() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_RemoveLogLevelError(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelError,
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
			if got := r.RemoveLogLevelError().RemoveLogLevelError(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.RemoveLogLevelError() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_RemoveLogLevelFatal(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			fields: fields{
				logLevel: LoggerDomain.LogLevelAllBits,
			},
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelAllBits - LoggerDomain.LogLevelFatal,
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
			if got := r.RemoveLogLevelFatal().RemoveLogLevelFatal(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.RemoveLogLevelFatal() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_SetLogLevelAll(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			wantR: &Repository{
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
			i := &Repository{
				name:       tt.fields.name,
				storage:    tt.fields.storage,
				logLevel:   tt.fields.logLevel,
				prefix:     tt.fields.prefix,
				timeFormat: tt.fields.timeFormat,
			}
			if got := i.SetLogLevelAll().SetLogLevelAll(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.SetLogLevelAll() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_SetLogLevelDebugInfoWarn(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone +
					LoggerDomain.LogLevelDebug +
					LoggerDomain.LogLevelInfo +
					LoggerDomain.LogLevelWarn,
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
			if got := r.SetLogLevelDebugInfoWarn().SetLogLevelDebugInfoWarn(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.SetLogLevelDebugInfoWarn() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_SetLogLevelInfoWarn(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone +
					LoggerDomain.LogLevelInfo +
					LoggerDomain.LogLevelWarn,
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
			if got := r.SetLogLevelInfoWarn().SetLogLevelInfoWarn(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.SetLogLevelInfoWarn() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}

func TestRepository_SetLogLevelErrorFatal(t *testing.T) {
	type fields struct {
		name       string
		storage    Storage
		logLevel   int
		prefix     string
		timeFormat string
	}
	tests := []struct {
		name   string
		fields fields
		wantR  *Repository
	}{
		{
			name: "Success",
			wantR: &Repository{
				logLevel: LoggerDomain.LogLevelNone +
					LoggerDomain.LogLevelError +
					LoggerDomain.LogLevelFatal,
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
			if got := r.SetLogLevelErrorFatal().SetLogLevelErrorFatal(); !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("Repository.SetLogLevelErrorFatal() = %v, wantR %v", got, tt.wantR)
			}
		})
	}
}
