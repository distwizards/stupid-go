package rand

import (
	"testing"
)

func TestInsecureRand(t *testing.T) {
	expectedLen := 10
	// Run this a few times to make sure it isn't a fluke
	for i := 0; i < 5; i++ {
		val := InsecureRand(expectedLen)
		if len(val) != expectedLen {
			t.Errorf("Expected string of length %d, got %d", expectedLen, len(val))
		}
	}
}

func TestInsecureRandSafe(t *testing.T) {
	expectedLen := 10
	// Run this a few times to make sure it isn't a fluke
	for i := 0; i < 5; i++ {
		val := InsecureRandSafe(expectedLen)
		if len(val) != expectedLen {
			t.Errorf("Expected string of length %d, got %d", expectedLen, len(val))
		}
	}
}

func BenchmarkInsecureRand10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRand(10)
	}
}

func BenchmarkInsecureRand100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRand(100)
	}
}

func BenchmarkInsecureRand1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRand(1000)
	}
}

func BenchmarkInsecureRand10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRand(10000)
	}
}

func BenchmarkInsecureRand100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRand(100000)
	}
}

func BenchmarkInsecureRandSafe10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRandSafe(10)
	}
}

func BenchmarkInsecureRandSafe100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRandSafe(100)
	}
}

func BenchmarkInsecureRandSafe1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRandSafe(1000)
	}
}

func BenchmarkInsecureRandSafe10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRandSafe(10000)
	}
}

func BenchmarkInsecureRandSafe100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsecureRandSafe(100000)
	}
}
