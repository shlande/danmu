package main

import (
	"apicore"
	"encoding/json"
	"errors"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (r *Response) Decode() []byte {
	by, err := json.Marshal(r)
	if err != nil {
		return apicore.NewServerErrorResponse(errors.New("内部错误")).Decode()
	}
	return by
}

func DanmuRep(data DanmuData) apicore.Response {
	return &Response{
		Code: 0,
		Data: data,
	}
}
