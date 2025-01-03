package fnt_test

import (
	"testing"

	"github.com/CSTCryst/fraudnettelemetry/internal"
	"github.com/CSTCryst/fraudnettelemetry/internal/fnt"
)

func BenchmarkScreen(b *testing.B) {
	internal.BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				screen_test := fnt.NewDCScreenBuilder()
				screen_test.AvailHeight = 724
				screen_test.AvailWidth = 1024
				screen_test.Reset(true)
			}
		})
	})
}
