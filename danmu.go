package main

type DanmuType uint8

const (
	NORMAL = iota
)

// 单个弹幕信息
type Danmu []interface{}

// TODO:处理时间
func (d Danmu) ShowTime() float64 {
	return d[0].(float64)
}
func (d Danmu) Type() DanmuType {
	return d[1].(DanmuType)
}
func (d Danmu) ID() string {
	return d[3].(string)
}
func (d Danmu) Color() int {
	return d[2].(int)
}
func (d Danmu) Content() string {
	return d[4].(string)
}

type DanmuData []Danmu
