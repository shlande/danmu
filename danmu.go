package main

import "time"

type DanmuType uint8

const (
	NORMAL = iota
)

// 单个弹幕信息
type Danmu []interface{}

// TODO:处理时间
func (d Danmu) ShowTime() time.Time {
	return time.Now()
}
func (d Danmu) Type() DanmuType {
	return NORMAL
}
func (d Danmu) SendTime() time.Time {
	return time.Now()
}
func (d Danmu) Color() string {
	return d[4].(string)
}
func (d Danmu) Content() string {
	return d[5].(string)
}

type DanmuData []Danmu
