package govalidator_test

import (
	"testing"
	"time"

	"github.com/bontusss/govalidator"
	"github.com/stretchr/testify/require"
)

func TestIsAfter(t *testing.T) {
	require := require.New(t)
	layout := time.RFC3339

	// Test case 1: Date is after comparisonDate
	date1 := "2023-10-06T12:00:00Z"
	comparisondate1 := "2023-10-05T12:00:00Z"
	result1 := govalidator.IsAfter(date1, layout, comparisondate1)
	require.True(result1)

	// Test case 2: Date is before comparisonDate
	date2 := "2023-10-06T12:00:00Z"
	comparisondate2 := "2023-10-07T12:00:00Z"
	result2 := govalidator.IsAfter(date2, layout, comparisondate2)
	require.False(result2)

	// Test case 3: Invalid date format
	date3 := "invalid-date"
	comparisondate3 := "2023-10-05T12:00:00Z"
	result3 := govalidator.IsAfter(date3, layout, comparisondate3)
	require.False(result3)

	// Test case 4: Default comparisonDate (current time)
	today := time.Now()
	tomorrow := today.AddDate(0,0,1).String()
	result4 := govalidator.IsAfter(tomorrow, layout, nil)
	require.False(result4)
}
