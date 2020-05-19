package main

import (
	"apicore"
	"errors"
)

func init() {
	apicore.AddHandler(apicore.NewMatcher("/danmu/v3/", "GET"), func() apicore.Handler {
		return &GetDanmuHandler{}
	})
}

type GetDanmuHandler struct {
	VideoID int `json:"id"`
	MaxNum  int `json:"max"`
}

func (d *GetDanmuHandler) Handle(ctx apicore.Conn) {
	manuData := DanmuStore.Get(d.VideoID)
	ctx.SetRsp(DanmuRep(manuData))
	return
}

func (d *GetDanmuHandler) IsValid() error {
	if d.VideoID == 0 {
		return errors.New("无效视频id")
	}
	return nil
}

func (d *GetDanmuHandler) mock(ctx apicore.Conn) {
	var dm = append(make([]interface{}, 0, 5), 61.781, 0, 16777215, "be7658b2", "测试弹幕")
	var DanmuData = DanmuData([]Danmu{dm})
	ctx.SetRsp(DanmuRep(DanmuData))
}
