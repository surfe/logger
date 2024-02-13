package simple

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/surfe/logger/logi"
)

func TestLogger_With(t *testing.T) {
	tests := []struct {
		name     string
		toAppend string
		args     []any
		want     logi.Logger
	}{
		{
			name:     "empty log",
			toAppend: "",
			args:     []any{},
			want: &Logger{
				toAppend: "",
			},
		},
		{
			name:     "Log with some keys",
			toAppend: "",
			args:     []any{"payload", struct{ key int }{key: 5}, "external", true},
			want: &Logger{
				toAppend: "payload: {5}; external: true; ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Logger{
				toAppend: tt.toAppend,
			}
			got := w.With(context.TODO(), tt.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}
