package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func (*CustomValidator) PhoneFormat() validator.Func {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().Interface().(string)
		re := regexp.MustCompile("^1[3456789]\\d{9}$")
		if re.FindString(data) == "" {
			return false
		} else {
			return true
		}
	}
}
