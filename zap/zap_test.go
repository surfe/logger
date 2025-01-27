package zap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Field string
}

func Test_appendFilledFieldsOnly(t *testing.T) {
	tests := []struct {
		name   string
		fields []any
		value  any
		want   []any
	}{
		{
			name:   "nil value",
			fields: []any{},
			value:  nil,
			want:   []any{},
		},
		{
			name:   "nil pointer",
			fields: []any{},
			value:  (*string)(nil),
			want:   []any{},
		},
		{
			name:   "empty string value",
			fields: []any{},
			value:  "",
			want:   []any{},
		},
		{
			name:   "string with spaces",
			fields: []any{},
			value:  "  ",
			want:   []any{"key", "  "},
		},
		{
			name:   "string with special characters",
			fields: []any{},
			value:  "!@#$%",
			want:   []any{"key", "!@#$%"},
		},
		{
			name:   "zero int value",
			fields: []any{},
			value:  0,
			want:   []any{},
		},
		{
			name:   "zero float value",
			fields: []any{},
			value:  0.0,
			want:   []any{},
		},
		{
			name:   "negative number",
			fields: []any{},
			value:  -1,
			want:   []any{"key", -1},
		},
		{
			name:   "positive float value",
			fields: []any{},
			value:  3.14,
			want:   []any{"key", 3.14},
		},
		{
			name:   "boolean false value",
			fields: []any{},
			value:  false,
			want:   []any{"key", false},
		},
		{
			name:   "boolean true value",
			fields: []any{},
			value:  true,
			want:   []any{"key", true},
		},
		{
			name:   "empty struct",
			fields: []any{},
			value:  testStruct{},
			want:   []any{},
		},
		{
			name:   "non-empty struct",
			fields: []any{},
			value:  testStruct{Field: "value"},
			want:   []any{"key", testStruct{Field: "value"}},
		},
		{
			name:   "empty slice",
			fields: []any{},
			value:  []string{},
			want:   []any{},
		},
		{
			name:   "non-empty slice",
			fields: []any{},
			value:  []string{"item"},
			want:   []any{"key", []string{"item"}},
		},
		{
			name:   "empty map",
			fields: []any{},
			value:  map[string]string{},
			want:   []any{},
		},
		{
			name:   "non-empty map",
			fields: []any{},
			value:  map[string]string{"key": "value"},
			want:   []any{"key", map[string]string{"key": "value"}},
		},
		{
			name:   "existing fields",
			fields: []any{"existing", 1},
			value:  "value",
			want:   []any{"existing", 1, "key", "value"},
		},
		{
			name:   "boolean true with existing fields",
			fields: []any{"existing", 1},
			value:  true,
			want:   []any{"existing", 1, "key", true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendFilledFieldsOnly(&tt.fields, "key", tt.value)
			assert.Equal(t, tt.want, tt.fields)
		})
	}
}
