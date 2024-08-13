package challenges

import (
	"cryptopals/utils"
	"fmt"
)

func challengeNine() {
	toPad := []byte("YELLOW SUBMARINE")
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	padded := utils.PKCS7Pad(toPad, 20)

	fmt.Printf("Expected: %q\n", expected)
	fmt.Printf("Actual:   %q\n", padded)
	fmt.Printf("UnPadded: %q\n", utils.PKCS7UnPad(padded))
}
