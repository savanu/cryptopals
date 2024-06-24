package utils

import (
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
