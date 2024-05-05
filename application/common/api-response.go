package applicationCommon

import (
	"time"
)

type ApiResponse struct {
	Code      int    `json:"code"`
	IsError   bool   `json:"isError"`
	Message   string `json:"message"`
	Result    any    `json:"result"`
	Timestamp int64  `json:"timestamp"`
}

func ApiResponseSuccess(message string, result any) *ApiResponse {
	res := &ApiResponse{
		Code:      200,
		IsError:   false,
		Message:   message,
		Result:    result,
		Timestamp: time.Now().UnixMilli(),
	}

	return res
}

func ApiResponseError(code int, message string, result any) *ApiResponse {
	res := &ApiResponse{
		Code:      code,
		IsError:   true,
		Message:   message,
		Result:    result,
		Timestamp: time.Now().UnixMilli(),
	}

	return res
}
