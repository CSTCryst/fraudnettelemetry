package fnt

import (
	"sync"

	"github.com/CSTCryst/fraudnettelemetry/internal"
)

var chkPool sync.Pool = sync.Pool{
	New: func() any {
		return new(Chk)
	},
}

type Chk struct {
	TS    int64         `json:"ts"`
	ETEID []interface{} `json:"eteid"`
	TTS   int64         `json:"tts"`
}

func NewChkBuilder() *Chk {
	return chkPool.Get().(*Chk)
}

func (m *Chk) String() (string, error) {
	if m == nil {
		return "", internal.NullStructError("chk")
	}
	return json.MarshalToString(&m)
}

func (m *Chk) Reset(put bool) {
	if m == nil {
		return
	}
	m.ETEID = nil
	m.TS = 0
	m.TTS = 0
}
