package main

import (
	"apicore"
	"context"
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

func DanmuRep(ctx context.Context, data DanmuData) context.Context {
	var resp = &Response{
		Code: 0,
		Data: data,
	}
	return context.WithValue(ctx, "SYS_RESPONSE", resp)
}
