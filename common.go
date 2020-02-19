package common

import (
	"net/http"
)

// NewResult returns a *common.Result with given code & msg
func NewResult(code int32, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
	}
}

// EMPTY is a global instance for commn.Empty
var EMPTY = &Empty{}

// OK represents ok result
var OK = NewResult(int32(http.StatusOK), "")

// SimpleOK represents a ok simple respsonse
var SimpleOK = &SimpleResp{
	Res: OK,
}
