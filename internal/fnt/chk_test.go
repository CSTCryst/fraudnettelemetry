package fnt_test

import (
	"testing"

	"github.com/CSTCryst/fraudnettelemetry/internal"
	"github.com/CSTCryst/fraudnettelemetry/internal/fnt"
)

func BenchmarkChk(b *testing.B) {
	internal.BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				chk_test := fnt.NewChkBuilder()
				chk_test.TS = 27
				chk_test.TTS = 29
				chk_test.ETEID = []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
				chk_test.Reset(true)
			}
		})
	})
}
