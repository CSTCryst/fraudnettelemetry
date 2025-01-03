package fnt_test

import (
	"testing"

	"github.com/CSTCryst/fraudnettelemetry/internal"
	"github.com/CSTCryst/fraudnettelemetry/internal/fnt"
)

func BenchmarkDC(b *testing.B) {
	internal.BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				dc_test := fnt.NewDCBuilder("TEST USER-AGENT")
				dc_test.Screen.ColorDepth = 24
				dc_test.Screen.Height = 1024
				dc_test.Screen.Width = 724
				dc_test.Reset(true)
			}
		})
	})
}
