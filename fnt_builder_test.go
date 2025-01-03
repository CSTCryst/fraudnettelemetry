package fraudnettelemetry_test

import (
	"testing"

	"github.com/CSTCryst/fraudnettelemetry"
	"github.com/CSTCryst/fraudnettelemetry/internal"
)

func BenchmarkFNTBuilder(b *testing.B) {
	internal.BenchPerCoreConfigs(b, func(b *testing.B) {
		b.RunParallel(func(b *testing.PB) {
			for b.Next() {
				vvv := fraudnettelemetry.V2_0_1.NewFNTBuilder()
				// vvv.String(true)
				// vvv.SetData(true)
				// vvv.SetDC("Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Mobile/15E148 Safari/604.1")
				// vvv.SetD([]string{"admin@gmail.com", "pass@321"}, true)
				_, _ = vvv.Generate([]string{"admin@gmail.com", "pass@321"}, []string{"UL_CHECKOUT_INPUT_EMAIL", "UL_CHECKOUT_INPUT_PASSWORD"}, "EC-4N177043M70444703", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Mobile/15E148 Safari/604.1", true)
				vvv.Reset(true)
			}
		})
	})
}
