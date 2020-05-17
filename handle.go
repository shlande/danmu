package main

import (
	"apicore"
	"context"
	"errors"
	"net/http"
)

func init() {
	apicore.AddHandler("/danmu/", func() apicore.Handler {
		return &DanmuHandler{}
	})
}

type DanmuHandler struct {
	VideoID int `json:"video_id"`
	MaxNum  int `json:"max_num"`
}

func (d *DanmuHandler) Handle(ctx context.Context, request *http.Request) context.Context {
	return d.mock(ctx)
}

func (d *DanmuHandler) IsValid() error {
	if d.VideoID == 0 {
		return errors.New("无效视频id")
	}
	return nil
}

func (d *DanmuHandler) mock(ctx context.Context) context.Context {
	var dm = append(make([]interface{}, 0, 5), 61.781, 0, 16777215, "be7658b2", "测试弹幕")
	var DanmuData = DanmuData([]Danmu{dm})
	return DanmuRep(ctx, DanmuData)
}
