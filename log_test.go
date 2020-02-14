package log

import (
	"bytes"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	type args struct {
		wc WriterConfig
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "init & info",
			args: args{
				wc: WriterConfig{
					AppName: "randomApp",
				},
			},
			wantW: "set it in the testcase below",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Init(w, tt.args.wc)
			Info("bug me not!")
			tt.wantW = "[randomApp] " + time.Now().Format("2006-01-02 15:04:05.999Z") + " [INFO] [log_test.go:32] bug me not!"
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("\n got = %v want= %v", gotW, tt.wantW)
			}
		})
	}
}
