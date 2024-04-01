package base

import "net/http"

type BaseResp struct {
	Code     int    `json:"code"`
	Messasge string `json:"message"`
	Data     any    `json:"data,omitempty"`
}

func BuildSuccessResp(data any) (resp *BaseResp) {

	resp = &BaseResp{
		Code:     0,
		Messasge: "ok",
		Data:     data,
	}

	return
}

func BuildFailedResp(code int, msg string) (resp *BaseResp) {

	resp = &BaseResp{
		Code:     code,
		Messasge: msg,
		Data:     nil,
	}

	return
}

func BuildInternalErrorResp(msg string) (resp *BaseResp) {

	resp = &BaseResp{
		Code:     http.StatusInternalServerError,
		Messasge: msg,
		Data:     nil,
	}

	return
}

func BuildBadRequestResp(msg string) (resp *BaseResp) {

	resp = &BaseResp{
		Code:     http.StatusBadRequest,
		Messasge: msg,
		Data:     nil,
	}

	return
}
