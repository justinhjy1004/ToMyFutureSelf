package smtp

import (
	"log"
	"testing"
)

func TestGenerateRandomNumber(t *testing.T) {
	// Define the test range
	min := 10
	max := 20

	// Run the test multiple times to ensure consistency
	for i := 0; i < 100; i++ {
		result := generateRandomNumber(min, max)

		// Check if the result is within the expected range
		if result < min || result >= max {
			t.Errorf("Generated number %d is out of range [%d, %d)", result, min, max)
		}
	}
}

func TestGenerateEmail(t *testing.T) {

	email, err := GenerateEmail("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	email.PrintEmail()
}
