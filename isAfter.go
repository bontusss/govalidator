package govalidator

import "time"


// This functions takes a date string (eg. "2023-10-06T12:00:00Z").
// 
// A layout (eg. time.RFC3339)
// 
// and an optional options parameter. It checks if the date is after the comparison date.
// 
// if options == nil, comparisonDate = current time
// 
//	 date := "2023-10-06T12:00:00Z"
//	 options := map[string]interface{}{"comparisonDate": "2023-10-05T12:00:00Z"}
// 	 result := govalidator.IsAfter(date, time.RFC3339, options)
//	 println(result)
func IsAfter(date, layout string, options interface{}) bool {
	var comparisonDate string

	// Check if options is a string, if so, set it as the comparisonDate
	if str, ok := options.(string); ok {
		comparisonDate = str
	} else if opts, ok := options.(map[string]interface{}); ok {
		// check if options is a map and if comparisonDate key exists
		if cmpDate, exists := opts["comparisonDate"].(string); exists {
			comparisonDate = cmpDate
		}
	}

	// if comparisonDate is still empty, use the current time as a string
	if comparisonDate == "" {
		comparisonDate = time.Now().String()
	}

	// Parse the dates as time.Time objects
	comparison, err := time.Parse(layout, comparisonDate)
	if err != nil {
		return false
	}
	
	original, err := time.Parse(layout, date)
	if err != nil {
		return false
	}

	return original.After(comparison)
}
