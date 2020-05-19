package main

import (
	"apicore"
	"errors"
)

func init() {
	apicore.AddHandler(apicore.NewMatcher("/danmu/v3/", "POST"), func() apicore.Handler {
		return &AddDanmuHandler{}
	})
}

type AddDanmuHandler struct {
	VedioID int `json:"id"`
	Author  string
	Time    float64
	Content string `json:"text"`
	Color   int
	Type    int
}

func (a *AddDanmuHandler) Handle(conn apicore.Conn) {
	danmu := Danmu{a.Time, a.Type, a.Color, a.Author, a.Content}
	DanmuStore.Add(a.VedioID, danmu)
	conn.SetRsp(DanmuRep(nil))
}

func (a *AddDanmuHandler) IsValid() error {
	if a.Content == "" {
		return errors.New("内容不能为空")
	}
	return nil
}
