package util

import (
	"common/defs"
	"fmt"
	"google.golang.org/grpc/codes"
	grpcerr "google.golang.org/grpc/status"
)

func MakeGrpcError(code int32, msg string) error {
	return grpcerr.Error(codes.Code(code), msg)
}

func FromGrpcError(err error) (int32, string) {
	s, _ := grpcerr.FromError(err)
	return int32(s.Code()), s.Message()
}

func Grpc2RstError(err error) ResultError {
	s, _ := grpcerr.FromError(err)
	return ResultError{Code: int32(s.Code()), Msg: s.Message()}
}

// //////////////////////////////////////////////
type ResultError struct {
	Code int32
	Msg  string
}

func (this *ResultError) IsOK() bool {
	return this.Code == defs.ErrOK
}

func (this *ResultError) Error() string {
	return fmt.Sprintf("error: code = %d, desc = %s", this.Code, this.Msg)
}

func (this *ResultError) ToGrpcError() error {
	return MakeGrpcError(this.Code, this.Msg)
}
