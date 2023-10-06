package govalidator

import "strings"

type ContainsOptions struct {
	// case sensitivity
	IgnoreCase    bool
	
	// Minimum number of occurences
	MinOccurences int
}

// Check if a string (str) contains a specified element (elem)
// a certain number of times, based on provided options
func Contains(str, elem string, options ContainsOptions) bool {
	if !options.IgnoreCase {
		return strings.Count(str, elem) > options.MinOccurences
	}

	str = strings.ToLower(str)
	elem = strings.ToLower(elem)
	return strings.Count(str, elem) > options.MinOccurences
}