package challenges

import (
	"cryptopals/utils"
	"encoding/hex"
	"fmt"
)

func challengeThree() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decoded, _ := hex.DecodeString(input)

	probableKey := utils.MostProbableKey(decoded)
	cipher := utils.XorByteArray(decoded, probableKey)

	fmt.Printf("Challenge Three: %s\n", string(cipher))
}
