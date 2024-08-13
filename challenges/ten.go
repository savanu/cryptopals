package challenges

import (
	"crypto/aes"
	"cryptopals/utils"
	"fmt"
	"os"
)

func challengeTen() {
	file, _ := os.ReadFile("files/challenge_10_data.txt")
	payload := utils.Base64Decode(file)
	var key = []byte("YELLOW SUBMARINE")
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	dec := aesCBCDecrypt(key, iv, payload)

	fmt.Printf("%s\n", dec)
}

func aesCBCDecrypt(key, iv, payload []byte) []byte {
	enc, _ := aes.NewCipher(key)
	clear := make([]byte, 0, 16)

	carry := iv
	dst := make([]byte, 16)
	for i := 0; i < len(payload); i += 16 {
		enc.Decrypt(dst, payload[i:i+16])
		xor, _ := utils.XorByteArrays(dst, carry)
		clear = append(clear, xor...)
		carry = payload[i : i+16]
	}

	return utils.PKCS7UnPad(clear)
}

func aesCBCEncrypt(key, iv, payload []byte) []byte {
	message := utils.PKCS7Pad(payload, 16)
	enc, _ := aes.NewCipher(key)

	final := make([]byte, len(message))
	carry := iv
	for i := 0; i < len(message); i += 16 {
		xor, _ := utils.XorByteArrays(carry, message[i:i+16])
		enc.Encrypt(final[i:i+16], xor)
		carry = final[i : i+16]
	}

	return final
}
