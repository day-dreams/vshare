package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	vld = validator.New()
)

// Validate 封装validator，校验结构体参数
func Validate(param interface{}) error {
	if param == nil {
		return nil
	}

	v := vld.Struct(param)

	if errs, ok := v.(validator.ValidationErrors); ok {
		fields := []string{}
		for _, err := range errs {
			fields = append(fields, err.Field())
		}
		msg := "参数不合法: " + strings.Join(fields, ",")
		return ErrParam(msg)
	}

	return nil
}
