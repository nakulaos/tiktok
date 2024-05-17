package xvalidate

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

var MinLenValidation = func(fl validator.FieldLevel) bool {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	num, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}

	if len(s) >= num {
		return true
	}

	return false

}

var MaxLenValidation = func(fl validator.FieldLevel) bool {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	num, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}

	if len(s) <= num {
		return true
	}

	return false

}
