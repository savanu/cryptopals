package utils

import "fmt"

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

		for l := 7; l >= 0; l-- {
			mask := byte((1 << 7) >> l)

			if (j & mask) != (k & mask) {
				distance++
			}
		}
	}

	return distance, nil
}

func XorBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("arrays are not equals length. a: %d, b: %d", len(a), len(b))
	}

	result := make([]byte, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}

	return result, nil
}

func MostProbableKey(xorEnc []byte) []byte {
	keyExpanded := make([]byte, len(xorEnc))
	highest := 0
	var probableKey int

	for i := 1; i <= 255; i++ {
		for j := range keyExpanded {
			keyExpanded[j] = byte(i)
		}

		text, _ := XorBytes(keyExpanded, xorEnc)
		if count := CommonLetterCount(text); count >= highest {
			highest = count
			probableKey = i
		}
	}

	probableKeyArr := make([]byte, len(xorEnc))
	for i := 0; i < len(xorEnc); i++ {
		probableKeyArr[i] = byte(probableKey)
	}

	return probableKeyArr
}

func CommonLetterCount(letters []byte) int {
	count := 0
	for _, letter := range letters {
		if is_common_letter(letter) {
			count += 1
		}
	}

	return count
}

func is_common_letter(x byte) bool {
	switch x {
	case 'E':
		return true
	case 'T':
		return true
	case 'A':
		return true
	case 'O':
		return true
	case 'I':
		return true
	case 'N':
		return true
	case 'S':
		return true
	case 'H':
		return true
	case 'R':
		return true
	case 'D':
		return true
	case 'L':
		return true
	case 'U':
		return true
	case 'C':
		return true
	case 'M':
		return true
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
	case 'c':
		return true
	case 'm':
		return true
	default:
		return false
	}
}
