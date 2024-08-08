package challenges

import (
	"fmt"
	"slices"
)

func challengeNine() {
	toPad := []byte("YELLOW SUBMARINE")
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	padded := pkcs7Padding(toPad, 20)

	fmt.Printf("Expected: %q\n", expected)
	fmt.Printf("Actual:   %q\n", padded)
}

// Perform PCKS#7 padding on the given block
func pkcs7Padding(block []byte, size int) []byte {
	padSize := size - (len(block) % size)
	padded := slices.Clone(block)

	for range padSize {
		padded = append(padded, byte(padSize))
	}

	return padded
}
