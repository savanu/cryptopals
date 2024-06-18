package setone

import (
	"cryptopals/utils"
	"encoding/hex"
	"fmt"
)

func ChallengeFive() {
	var text string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	var expected string = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	key := [3]byte{'I', 'C', 'E'}
	keyExtended := make([]byte, len(text))

	for i := 0; i < len(text); i++ {
		keyExtended[i] = key[i%3]
	}

	cipher, _ := utils.XorBytes([]byte(text), keyExtended)
	cipherHex := hex.EncodeToString(cipher)

	fmt.Printf("Calculated: %v\n", cipherHex)
	fmt.Printf("Expected:   %v\n", expected)
	fmt.Printf("Matches? %v\n", cipherHex == expected)
}
