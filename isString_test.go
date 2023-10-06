package govalidator_test

import (
	"testing"

	"github.com/bontusss/govalidator"
	"github.com/stretchr/testify/require"
)

func TestIsString(t *testing.T) {
	assert := require.New(t)
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{"Hello World", true},
		{34, false},
		{2.23, false},
		{nil, false},
	}
	for _, test := range tests {
		result := govalidator.IsString(test.input)
		assert.Equal(test.expected, result)
	}
}
