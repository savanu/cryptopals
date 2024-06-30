package setone

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func ChallengeSeven() {
	key := []byte("YELLOW SUBMARINE")

	base64Encoded, err := os.ReadFile("files/challenge_7_data.txt")
	if err != nil {
		log.Fatalf("Couldn't open file: %s", err.Error())
	}

	data := make([]byte, base64.StdEncoding.DecodedLen(len(base64Encoded)))
	dataSize, _ := base64.StdEncoding.Decode(data, base64Encoded)
	data = data[:dataSize]

	cipher, _ := aes.NewCipher(key)
	plainText := make([]byte, dataSize)
	for i := 0; i < dataSize; i += 16 {
		cipher.Decrypt(plainText[i:i+16], data[i:i+16])
	}

	fmt.Printf("%s\n", plainText)
}
