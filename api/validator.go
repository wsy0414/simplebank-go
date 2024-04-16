package api

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

// PasswordValidator verify password
var PasswordValidator validator.Func = func(fieldVal validator.FieldLevel) bool {
	var (
		hasUpperCase = false
		hasLowerCase = false
		hasDigit     = false
		hasNoMark    = true
	)

	if pwd, ok := fieldVal.Field().Interface().(string); ok {
		if len(pwd) < 8 {
			return false
		}
		for _, s := range pwd {
			switch {
			case unicode.IsUpper(s):
				hasUpperCase = true
			case unicode.IsLower(s):
				hasLowerCase = true
			case unicode.IsDigit(s):
				hasDigit = true
			default:
				hasNoMark = false
			}
		}

		return hasLowerCase && hasUpperCase && hasDigit && hasNoMark
	}

	return false
}
