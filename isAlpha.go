package govalidator

import (
	"errors"
	"regexp"
)

var alpha = map[string]*regexp.Regexp{
	"en-US": regexp.MustCompile("^[A-Za-z]+$"),
	// Add other locales here
}

func IsAlpha(str, locale string, options map[string]interface{}) (bool, error) {
	ignore, ok := options["ignore"]
	if ok {
		switch ignore := ignore.(type) {
		case *regexp.Regexp:
			str = ignore.ReplaceAllString(str, "")
		case string:
			escaped := regexp.QuoteMeta(ignore)
			str = regexp.MustCompile("["+escaped+"]").ReplaceAllString(str, "")
		default:
			return false, errors.New("ignore should be an instance of a String or RegExp")
		}
	}

	if _, exists := alpha[locale]; exists {
		return alpha[locale].MatchString(str), nil
	}

	return false, errors.New("Invalid locale '" + locale + "'")
}
