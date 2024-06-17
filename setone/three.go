package setone

import (
	"cryptopals/utils"
	"encoding/hex"
	"fmt"
)

func ChallengeThree() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decoded, _ := hex.DecodeString(input)

	cipher := utils.MostProbableText(decoded)

	fmt.Printf("Challenge Three: %s\n", string(cipher[:]))
}
