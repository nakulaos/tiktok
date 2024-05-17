package xvalidate

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

var PaswordValidation validator.Func = func(fl validator.FieldLevel) bool {
	// 校验字符是否含有大写字母，小写字母，数字，特殊字符
	password, ok := fl.Field().Interface().(string)
	if ok {
		flagLower := false
		flagUpper := false
		flagNumeric := false
		flagPunch := false

		for _, value := range password {
			if unicode.IsLower(value) {
				flagLower = true
				break
			}
		}

		for _, value := range password {
			if unicode.IsUpper(value) {
				flagUpper = true
				break
			}
		}

		for _, value := range password {
			if unicode.IsNumber(value) {
				flagNumeric = true
				break
			}
		}

		for _, value := range password {
			if unicode.IsPunct(value) {
				flagPunch = true
				break
			}
		}

		return flagNumeric && flagUpper && flagLower && flagPunch

	}
	return false
}
