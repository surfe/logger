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
		key    string
		value  any
		want   []any
	}{
		// Nil and Zero cases
		{
			name:   "nil value",
			fields: []any{},
			key:    "key",
			value:  nil,
			want:   []any{},
		},
		{
			name:   "nil pointer",
			fields: []any{},
			key:    "key",
			value:  (*string)(nil),
			want:   []any{},
		},

		// String cases
		{
			name:   "zero string value",
			fields: []any{},
			key:    "key",
			value:  "",
			want:   []any{},
		},
		{
			name:   "string with spaces",
			fields: []any{},
			key:    "key",
			value:  "  ",
			want:   []any{"key", "  "},
		},
		{
			name:   "special characters",
			fields: []any{},
			key:    "key",
			value:  "!@#$%",
			want:   []any{"key", "!@#$%"},
		},

		// Numeric cases
		{
			name:   "zero int value",
			fields: []any{},
			key:    "key",
			value:  0,
			want:   []any{},
		},
		{
			name:   "zero float value",
			fields: []any{},
			key:    "key",
			value:  0.0,
			want:   []any{},
		},
		{
			name:   "negative number",
			fields: []any{},
			key:    "key",
			value:  -1,
			want:   []any{"key", -1},
		},
		{
			name:   "positive float value",
			fields: []any{},
			key:    "key",
			value:  3.14,
			want:   []any{"key", 3.14},
		},

		// Boolean cases
		{
			name:   "boolean false value",
			fields: []any{},
			key:    "key",
			value:  false,
			want:   []any{"key", false},
		},
		{
			name:   "boolean true value",
			fields: []any{},
			key:    "key",
			value:  true,
			want:   []any{"key", true},
		},

		// Complex types
		{
			name:   "empty struct",
			fields: []any{},
			key:    "key",
			value:  testStruct{},
			want:   []any{},
		},
		{
			name:   "non-empty struct",
			fields: []any{},
			key:    "key",
			value:  testStruct{Field: "value"},
			want:   []any{"key", testStruct{Field: "value"}},
		},
		{
			name:   "empty slice",
			fields: []any{},
			key:    "key",
			value:  []string{},
			want:   []any{},
		},
		{
			name:   "non-empty slice",
			fields: []any{},
			key:    "key",
			value:  []string{"item"},
			want:   []any{"key", []string{"item"}},
		},
		{
			name:   "empty map",
			fields: []any{},
			key:    "key",
			value:  map[string]string{},
			want:   []any{},
		},
		{
			name:   "non-empty map",
			fields: []any{},
			key:    "key",
			value:  map[string]string{"key": "value"},
			want:   []any{"key", map[string]string{"key": "value"}},
		},

		// Edge cases
		{
			name:   "empty key",
			fields: []any{},
			key:    "",
			value:  "value",
			want:   []any{"", "value"},
		},
		{
			name:   "existing fields",
			fields: []any{"existing", 1},
			key:    "key",
			value:  "value",
			want:   []any{"existing", 1, "key", "value"},
		},
		{
			name:   "nil key",
			fields: []any{},
			key:    "nil",
			value:  "value",
			want:   []any{"nil", "value"},
		},
		{
			name:   "boolean true with existing fields",
			fields: []any{"existing", 1},
			key:    "key",
			value:  true,
			want:   []any{"existing", 1, "key", true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendFilledFieldsOnly(&tt.fields, tt.key, tt.value)
			assert.Equal(t, tt.want, tt.fields)
		})
	}
}
