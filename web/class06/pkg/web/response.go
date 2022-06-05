package web

import "strconv"

type Response struct {
	Code string `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewResponse(code int64, data interface{}, err string) Response {
	if code < 300 {
		return Response{strconv.FormatInt(code, 10), data, ""}
	}

	return Response{strconv.FormatInt(code, 10), nil, err}
}