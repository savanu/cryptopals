package setone

import (
	"encoding/hex"
	"fmt"
	"log"
)

func ChallengeTwo() {
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
