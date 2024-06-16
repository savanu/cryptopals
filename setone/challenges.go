package setone

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func Run(challengeNum int) {
	switch challengeNum {
	case 1:
		challengeOne()
	case 2:
		challengeTwo()
	case 3:
		challengeThree()
	}
}

// Encodes a given hex string as base64.
//
// Hex String -> Bytes -> Base 64 String
func challengeOne() {
	answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	payload := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	bytes, _ := hex.DecodeString(payload)
	encoded := base64.StdEncoding.EncodeToString(bytes)

	if answer != encoded {
		log.Fatalf("Failure. Expected: %s Actual: %s", answer, encoded)
	} else {
		fmt.Printf("Challenge One: %s\n", encoded)
	}
}

func challengeTwo() {
	arrOne := "1c0111001f010100061a024b53535009181c"
	arrTwo := "686974207468652062756c6c277320657965"
	answer := "746865206b696420646f6e277420706c6179"

	arrOneBytes, _ := hex.DecodeString(arrOne)
	arrTwoBytes, _ := hex.DecodeString(arrTwo)

	resultBytes, _ := xorBytes(arrOneBytes, arrTwoBytes)

	result := hex.EncodeToString(resultBytes)

	if result != answer {
		log.Fatalf("Failure. Expected: %s Actual: %s", answer, result)
	} else {
		fmt.Printf("Challenge Two: %s\n", result)
	}
}

func challengeThree() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decoded, _ := hex.DecodeString(input)

	keyExpanded := make([]byte, len(decoded))
	highest := 0
	var cipher []byte

	for i := 1; i <= 255; i++ {
		for j := range keyExpanded {
			keyExpanded[j] = byte(i)
		}

		text, _ := xorBytes(keyExpanded, decoded)
		if count := common_letter_count(text); count >= highest {
			highest = count
			cipher = text
		}
	}

	fmt.Printf("Challenge Three: %s\n", string(cipher[:]))
}

func xorBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("arrays are not equals length. a: %d, b: %d", len(a), len(b))
	}

	result := make([]byte, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}

	return result, nil
}

func common_letter_count(letters []byte) int {
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
