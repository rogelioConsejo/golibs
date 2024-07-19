package helpers

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMakeRandomString(t *testing.T) {
	t.Parallel()
	t.Run("MakeRandomString returns a string of the correct size", testMakeRandomStringReturnsCorrectSize)
	t.Run("MakeRandomString returns a different string each time", testMakeRandomStringReturnsDifferentString)
}

func testMakeRandomStringReturnsCorrectSize(t *testing.T) {
	t.Parallel()
	stringSize := rand.Intn(100)
	randomString := MakeRandomString(stringSize)
	if len(randomString) != stringSize {
		t.Errorf("expected string of size %d, but got %d", stringSize, len(randomString))
	}
}

func testMakeRandomStringReturnsDifferentString(t *testing.T) {
	t.Parallel()
	stringSize := rand.Intn(100) + 5
	iterations := rand.Intn(1000) + 100
	fmt.Printf("string size: %d\n", stringSize)
	fmt.Printf("iterations: %d\n", iterations)
	createdStrings := make(map[string]bool)
	for _ = range iterations {
		randomString := MakeRandomString(stringSize)
		if _, ok := createdStrings[randomString]; ok {
			t.Errorf("generated the same string %s twice", randomString)
		}
		createdStrings[randomString] = true
	}
}
