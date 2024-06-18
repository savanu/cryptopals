package setone

import (
	"cryptopals/utils"
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

	resultBytes, _ := utils.XorBytes(arrOneBytes, arrTwoBytes)

	result := hex.EncodeToString(resultBytes)

	if result != answer {
		log.Fatalf("Failure. Expected: %s Actual: %s", answer, result)
	} else {
		fmt.Printf("Challenge Two: %s\n", result)
	}
}
