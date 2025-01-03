package fnt

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var (
	dcScreenPool sync.Pool = sync.Pool{
		New: func() any {
			return new(dcScreen)
		},
	}
	json jsoniter.API = jsoniter.ConfigFastest
)

type dcScreen struct {
	ColorDepth  uint8  `json:"colorDepth"`
	PixelDepth  uint8  `json:"pixelDepth"`
	Height      uint32 `json:"height"`
	Width       uint32 `json:"width"`
	AvailHeight uint32 `json:"availHeight"`
	AvailWidth  uint32 `json:"availWidth"`
}

func NewDCScreenBuilder() *dcScreen {
	return dcScreenPool.Get().(*dcScreen)
}

func (m *dcScreen) String() (string, error) {
	return json.MarshalToString(&m)
}

func (m *dcScreen) Reset(put bool) {
	if m == nil {
		return
	}
	m.AvailHeight = 0
	m.AvailWidth = 0
	m.ColorDepth = 0
	m.Height = 0
	m.PixelDepth = 0
	m.Width = 0
	if put {
		dcScreenPool.Put(m)
	}
}
