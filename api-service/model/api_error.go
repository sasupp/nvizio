package model

import "fmt"

const UnknownErrorCode = 5001

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}
