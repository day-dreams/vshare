package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ParamInvalid = 5001
)

var (
	ErrParamInvalid = status.Error(codes.Code(ParamInvalid), "参数错误")
)

func ErrParam(msg string) error {
	return status.Error(codes.Code(ParamInvalid), msg)
}
