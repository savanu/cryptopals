package utils

import (
	"bytes"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	strOne := []byte("this is a test")
	strTwo := []byte("wokka wokka!!!")
	expectedDist := 37

	dist, err := HammingDistance(strOne, strTwo)
	if err != nil {
		t.Fatalf("Hamming distance failed due to unexpected error: %s\n", err)
	}

	if dist != expectedDist {
		t.Fatalf("Hamming distance incorrect. Expected: %d Actual: %d", expectedDist, dist)
	}
}

func TestPKCS7PaddingTest(t *testing.T) {
	toPad := []byte("YELLOW SUBMARINE")
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	padded := PKCS7Padding(toPad, 20)

	if !bytes.Equal(padded, expected) {
		t.Fatalf("Padding test failed. Expected: %q ; Actual: %q\n", expected, padded)
	}
}
