package govalidator

// IsString takes an empty interface (interface{}) as an argument
// and uses type assertion to check if it can be converted to a string type.
// The function returns `true`; otherwise, it returns `false`
func IsString(text interface{}) bool {
	_, isString := text.(string)
	return isString
}