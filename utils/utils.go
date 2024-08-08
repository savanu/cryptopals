package utils

import (
	"fmt"
)

var masks = []byte{
	0b00000001,
	0b00000010,
	0b00000100,
	0b00001000,
	0b00010000,
	0b00100000,
	0b01000000,
	0b10000000,
}

func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("input arrays are not same length. a: %d b: %d", len(a), len(b))
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		j := a[i]
		k := b[i]

		if j == k {
			continue
		}

		for _, mask := range masks {
			if (j & mask) != (k & mask) {
				distance++
			}
		}
	}

	return distance, nil
}

func XorByteArray(a []byte, b byte) []byte {
	result := make([]byte, len(a))

	for i, v := range a {
		result[i] = v ^ b
	}

	return result
}

func XorByteArrays(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("arrays are not equals length. a: %d, b: %d", len(a), len(b))
	}

	result := make([]byte, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}

	return result, nil
}

func MostProbableKey(xorEnc []byte) byte {
	highest := 0
	var probableKey int

	for i := 1; i <= 255; i++ {
		text := XorByteArray(xorEnc, byte(i))
		if count := CommonLetterCount(text); count >= highest {
			highest = count
			probableKey = i
		}
	}

	return byte(probableKey)
}

func CommonLetterCount(letters []byte) int {
	count := 0
	for _, letter := range letters {
		if isCommonLetter(letter) {
			count += 1
		}
	}

	return count
}

func isCommonLetter(x byte) bool {
	switch x {
	case 'e':
		return true
	case 't':
		return true
	case 'a':
		return true
	case 'o':
		return true
	case 'i':
		return true
	case 'n':
		return true
	case 's':
		return true
	case 'h':
		return true
	case 'r':
		return true
	case 'd':
		return true
	case 'l':
		return true
	case 'u':
		return true
	case ' ':
		return true
	default:
		return false
	}
}
