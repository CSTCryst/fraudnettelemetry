package device_screen

import "sync"

var screenPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return new(screen)
	},
}

type screen struct {
	LogicalResolution  [2]uint32
	PhysicalResolution [2]uint32
	ColorDepth         uint8
}

func NewScreenBuilder() *screen {
	return screenPool.Get().(*screen)
}

func (m *screen) Reset(put bool) {
	if m == nil {
		return
	}
	m.LogicalResolution = [2]uint32{}
	m.PhysicalResolution = [2]uint32{}
	m.ColorDepth = 0
	if put {
		screenPool.Put(m)
	}
}
