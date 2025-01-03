package internal

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchPerCoreConfigs(b *testing.B, f func(b *testing.B)) {
	b.Helper()
	coreConfigs := []int{1, 2, 4, 8, 12}
	for _, n := range coreConfigs {
		name := fmt.Sprintf("%d cores", n)
		b.Run(name, func(b *testing.B) {
			runtime.GOMAXPROCS(n)
			f(b)
		})
	}
}
