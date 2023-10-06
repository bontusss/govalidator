package main

import (
	"github.com/bontusss/govalidator"
)

func main() {
	str := "Yourinputstring"
	locale := "en-US"
	options := map[string]interface{}{
		"ignore": "abc",
	}

	result, err := govalidator.IsAlpha(str, locale, options)
	if err != nil {
		panic(err)
	}
	if result {
		println("input is alpha")
	} else {
		println("input is not alpha")
	}
}
