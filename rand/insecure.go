package rand

import (
	"math"
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

var (
	// This is the formula for the number of bits that is needed to represent
	// the number of characters in the letters
	indexBits = uint(math.Floor(math.Log2(float64(len(letters)))) + 1)

	// The mask to use to select random bits
	indexMask = int64(1)<<indexBits - 1

	// The number of indices we can read using 63 bits
	numReads = 63 / indexBits

	// A custom source for non-concurrent InsecureRand calls
	src = rand.NewSource(time.Now().UnixNano())
)

func init() {
	// Set the random seed for the default functions
	rand.Seed(time.Now().UnixNano())
}

// This optimized version is heavily based off of this wonderful Stack Overflow
// answer:
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

// StringUnsafe returns a random string of the specified length that is
// sufficiently random for use in identifiers but should not ever be used to
// seed cryptographic operations. It is not safe to use concurrently
func StringUnsafe(n int) string {
	return getString(n, src.Int63)
}

// String returns a random string of the specified length that is
// sufficiently random for use in identifiers but should not ever be used to
// seed cryptographic operations. It is safe to use concurrently
func String(n int) string {
	return getString(n, rand.Int63)
}

func getString(n int, f func() int64) string {
	b := make([]byte, n)
	bitCache := f()
	remainingReads := numReads
	for i := 0; i < n; {
		// If we are out of bits, read again
		if remainingReads == 0 {
			bitCache = f()
			remainingReads = numReads
		}
		// Because we may have one extra bit, the index could be out of range.
		// If it is in range, use it. Otherwise, throw away
		if index := int(bitCache & indexMask); index < len(letters) {
			b[i] = letters[index]
			i++
		}

		remainingReads--
		// shift off the number of bits we used
		bitCache >>= indexBits
	}
	return string(b)
}
