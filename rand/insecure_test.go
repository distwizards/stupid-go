package rand

import (
	"testing"
)

func TestStringUnsafe(t *testing.T) {
	// Yes, this test is not entirely statistically sound, but it provides some
	// basic validation that it isn't returning the same string. Also, the
	// math/rand package is fully tested against statistical distributions,
	// meaning we should be safe in assuming we get sufficiently random data
	// here
	expectedLen := 10
	var existingMap = make(map[string]bool)
	// Run this a few times to make sure it isn't a fluke
	for i := 0; i < 50; i++ {
		val := StringUnsafe(expectedLen)
		if len(val) != expectedLen {
			t.Errorf("Expected string of length %d, got %d", expectedLen, len(val))
		}
		if existingMap[val] {
			t.Error("Found preexisting string of the same value")
		}
		existingMap[val] = true
	}
}

func TestString(t *testing.T) {
	// Yes, this test is not entirely statistically sound, but it provides some
	// basic validation that it isn't returning the same string. Also, the
	// math/rand package is fully tested against statistical distributions,
	// meaning we should be safe in assuming we get sufficiently random data
	// here
	expectedLen := 10
	var existingMap = make(map[string]bool)
	// Run this a few times to make sure it isn't a fluke
	for i := 0; i < 50; i++ {
		val := String(expectedLen)
		if len(val) != expectedLen {
			t.Errorf("Expected string of length %d, got %d", expectedLen, len(val))
		}
		if existingMap[val] {
			t.Error("Found preexisting string of the same value")
		}
		existingMap[val] = true
	}
}

func BenchmarkStringUnsafe10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringUnsafe(10)
	}
}

func BenchmarkStringUnsafe100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringUnsafe(100)
	}
}

func BenchmarkStringUnsafe1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringUnsafe(1000)
	}
}

func BenchmarkStringUnsafe10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringUnsafe(10000)
	}
}

func BenchmarkStringUnsafe100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringUnsafe(100000)
	}
}

func BenchmarkString10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(10)
	}
}

func BenchmarkString100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(100)
	}
}

func BenchmarkString1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(1000)
	}
}

func BenchmarkString10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(10000)
	}
}

func BenchmarkString100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(100000)
	}
}
