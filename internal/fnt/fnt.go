package fnt

import (
	"fmt"
	"sync"

	"github.com/CSTCryst/fraudnettelemetry/internal"
	"github.com/colduction/rfc3986"
)

var (
	fntBasePool sync.Pool = sync.Pool{
		New: func() interface{} {
			return new(fntBase)
		},
	}
)

type IFNT interface {
	String(urlEncode bool) (string, error)
	Reset(put bool)
}

type fntBase struct {
	SCVersion          string            `json:"SC_VERSION"`
	SyncStatus         string            `json:"syncStatus"`
	F                  string            `json:"f"`
	S                  string            `json:"s"`
	Chk                *Chk              `json:"chk"`
	DC                 string            `json:"dc"`
	WV                 bool              `json:"wv"`
	WebIntegrationType string            `json:"web_integration_type"`
	CookieEnabled      bool              `json:"cookie_enabled"`
	D                  map[string]string `json:"d"`
}

func NewFNTBaseBuilder() *fntBase {
	m := fntBasePool.Get().(*fntBase)
	m.Chk = NewChkBuilder()
	return m
}

func (m *fntBase) String(urlEncode bool) (string, error) {
	if m == nil {
		return "", internal.NullStructError("FNTBase")
	}
	switch m.SCVersion {
	case "0.1.9", "2.0.1":
		data, err := json.MarshalToString(&m)
		if err != nil {
			return "", err
		}
		if urlEncode {
			return rfc3986.QueryEscape(data), nil
		}
		return data, nil
	case "2.0.4":
		data, err := json.MarshalToString(&m)
		if err != nil {
			return "", err
		}
		out := fmt.Sprintf("{\"fn_sync_data\":\"%s\"}", rfc3986.QueryEscape(data))
		if urlEncode {
			return rfc3986.QueryEscape(out), nil
		}
		return out, nil
	default:
		return "", internal.UndefinedVersionError(m.SCVersion)
	}
}

func (m *fntBase) Reset(put bool) {
	if m == nil {
		return
	}
	m.Chk.Reset(put)
	m.Chk = nil
	m.D = nil
	m.DC = ""
	m.F = ""
	m.S = ""
	m.SCVersion = ""
	if put {
		fntBasePool.Put(m)
	}
}
