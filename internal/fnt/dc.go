package fnt

import (
	"sync"
)

var dcPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return new(dc)
	},
}

type dc struct {
	Screen *dcScreen `json:"screen"`
	UA     string    `json:"ua"`
}

func NewDCBuilder(userAgent string) *dc {
	m := dcPool.Get().(*dc)
	m.Screen = dcScreenPool.Get().(*dcScreen)
	m.UA = userAgent
	return m
}

func (m *dc) String() (string, error) {
	return json.MarshalToString(&m)
}

func (m *dc) Reset(put bool) {
	m.UA = ""
	m.Screen.Reset(put)
	if put {
		dcPool.Put(m)
	}
}
