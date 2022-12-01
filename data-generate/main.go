package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/bxcodec/faker/v3"
)

type FakeData struct {
	Currency          string `faker:"currency"`
	Email             string `faker:"email"`
	ReferralCode      string `faker:"uuid_hyphenated"`
	ReferralSubmitted string `faker:"timestamp"`
	Referree          string `faker:"name"`
	Referrer          string `faker:"name"`
	UserName          string `faker:"username"`
}

var NumberOfRecordsToLog = 10000

func main() {
	numberOfRecords, err := strconv.Atoi(os.Args[1])
	// this is for making it easier to test locally vs running in a container
	outPutDirectory := os.Args[2]

	jsonData := []FakeData{}

	for i := 0; i < numberOfRecords+1; i++ {
		a := FakeData{}

		err := faker.FakeData(&a)
		if err != nil {
			log.Println(err)
		}

		if i%1000 == 0 {
			log.Printf("generated record %d out of %d\n", i, numberOfRecords)
		}

		jsonData = append(jsonData, a)

		if i > 0 {
			if (i % NumberOfRecordsToLog) == 0 {
				comment := fmt.Sprintf("writing out to file: record-%d.json", i)
				log.Println(comment)

				jsonBlob, _ := json.MarshalIndent(jsonData, "", " ")
				_ = ioutil.WriteFile(fmt.Sprintf("%s/record-%d.json", outPutDirectory, i), jsonBlob, 0o644)
				jsonData = []FakeData{}
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finished writing to file.")
}
