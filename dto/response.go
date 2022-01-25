package dto

type SuccessResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
