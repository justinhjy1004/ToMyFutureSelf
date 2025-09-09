package smtp

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func generateRandomNumber(min int, max int) int {

	return rand.Intn(max-min) + min

}

func GenerateEmail(inputFileName string) (Email, error) {

	file, err := os.ReadFile(inputFileName)

	if err != nil {
		log.Fatal(err)
	}

	parsedEmail := strings.SplitN(string(file), "\n", 2)

	var ErrInputFormatError = fmt.Errorf("input format error! should in: \n # <subject> \n <body>")

	if len(parsedEmail) != 2 {
		return Email{}, ErrInputFormatError
	}

	now := time.Now()

	var daysFromNow int = generateRandomNumber(30, 240)
	sendTime := now.AddDate(0, 0, daysFromNow)

	email := Email{TimeCreated: now, TimeToBeSent: sendTime, Subject: parsedEmail[0], Body: parsedEmail[1]}

	return email, nil
}
