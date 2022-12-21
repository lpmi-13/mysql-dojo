package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/bxcodec/faker"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type FakeData struct {
	Currency          string `faker:"currency"`
	Email             string `faker:"email"`
	ReferralCode      string `faker:"uuid_hyphenated"`
	ReferralSubmitted string `faker:"timestamp"`
	Referree          string `faker:"name"`
	Referrer          string `faker:"name"`
	Username          string `faker:"username"`
}

var (
	ctx       = context.TODO()
	BatchSize = 1000
	db        *gorm.DB
	err       error
)

func main() {
	host := os.Args[1]
	totalRecords, _ := strconv.Atoi(os.Args[2])
	dsn := "root:password@tcp(" + host + ":3306)/transactions?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < totalRecords; i++ {
		InsertData()

		if i%100 == 0 && i != 0 {
			log.Println("created record:", i*BatchSize)
		}
	}

	log.Println("all finished")
}

func InsertData() {
	referralBatch := []FakeData{}

	for i := 0; i < BatchSize; i++ {
		referral := FakeData{}

		err := faker.FakeData(&referral)
		if err != nil {
			log.Println(err)
		}

		referralBatch = append(referralBatch, referral)
	}

	err = db.Table("referrals").CreateInBatches(referralBatch, BatchSize).Error
	if err != nil {
		log.Fatal(err)
	}
}
