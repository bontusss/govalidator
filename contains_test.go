package govalidator_test

import (
	"testing"

	"github.com/bontusss/govalidator"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	tests := []struct {
		str            string
		elem           string
		options        govalidator.ContainsOptions
		expectedResult bool
	}{
		// Test case 1: Case-sensitive element present once
		{
			str:            "This is a sample string to test contains function",
			elem:           "sample",
			expectedResult: true,
		},
		// Test case 2: Case-sensitive element present more than once
		{
			str:            "sample sample sample",
			elem:           "sample",
			options:        govalidator.ContainsOptions{IgnoreCase: false, MinOccurences: 2},
			expectedResult: true,
		},
		// Test case 3: Case-sensitive element not present
		{
			str:            "This is a test",
			elem:           "missing",
			expectedResult: false,
		},
		// Test case 4: Case-insensitive element present once
		{
			str:            "This is a sample string to test contains function",
			elem:           "sAmple",
			options:        govalidator.ContainsOptions{IgnoreCase: true},
			expectedResult: true,
		},
		
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			require := require.New(t)
			result := govalidator.Contains(test.str, test.elem, test.options)

			require.Equal(test.expectedResult, result)
		})
	}
}
