package xvalidate

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
	ErrorCode int
}

func NewValidator() *Validator {
	v := Validator{}
	// set default error code to 3
	v.ErrorCode = 100001
	v.Validator = validator.New()
	v.Validator.RegisterValidation("password", PaswordValidation)
	v.Validator.RegisterValidation("min_len", MinLenValidation)
	v.Validator.RegisterValidation("max_len", MaxLenValidation)

	return &v
}

func (v *Validator) Validate(data any) string {
	err := v.Validator.Struct(data)
	if err == nil {
		return ""
	}

	errs, ok := err.(validator.ValidationErrors)

	if ok {
		return GetValidMsg(errs, data)
	}

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return invalid.Error()
	}

	return ""
}

// SetValidatorErrorCode sets the error code for validator when errors occurs
func (v *Validator) SetValidatorErrorCode(code int) {
	v.ErrorCode = code
}

func GetValidMsg(err error, obj any) string {
	// obj 的类型
	// TypeOf 返回i的动态类型信息，即reflect.Type类型
	getObj := reflect.TypeOf(obj)

	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := fmt.Sprintf("RequestErr." + fmt.Sprintf(f.Tag.Get("msg")))
				return msg
			}
		}
	}
	return err.Error()
}
