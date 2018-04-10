package weather

import "testing"

func Benchmark_toDigitalPing(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toDigitalPing(89)
	}
}

func Benchmark_toDigitalPack(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toDigitalPack(89)
	}
}

func Benchmark_toDigitalBell(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toDigitalBell(89)
	}
}
