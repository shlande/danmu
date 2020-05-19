package main

func init() {
	DanmuStore = NewMemoryStore()
}

type Store interface {
	Add(videoID int, danmu Danmu)
	Get(videoID int) DanmuData
}

var DanmuStore Store

func NewMemoryStore() Store {
	return &MemoryStore{
		catches: make(map[int]DanmuData),
	}
}

type MemoryStore struct {
	catches map[int]DanmuData
}

func (m *MemoryStore) Add(videoID int, danmu Danmu) {
	var damnuData DanmuData
	var ok bool
	if damnuData, ok = m.catches[videoID]; !ok {
		damnuData = DanmuData{}
		m.catches[videoID] = damnuData
	}
	m.catches[videoID] = append(damnuData, danmu)
}

func (m *MemoryStore) Get(videoID int) DanmuData {
	if data, ok := m.catches[videoID]; ok {
		return data
	}
	return DanmuData{}
}
