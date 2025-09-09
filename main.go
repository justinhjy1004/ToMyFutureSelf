package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/justinhjy1004/ToMyFutureSelf/smtp"
)

func main() {

	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	//log.Println(strings.SplitN(string(file), "\n", 2)[1])

	email := smtp.Email{TimeCreated: time.Now(), TimeToBeSent: time.Now().AddDate(0, 0, 1000), Subject: strings.SplitN(string(file), "\n", 2)[0], Body: strings.SplitN(string(file), "\n", 2)[1]}

	fmt.Println(email)
}
