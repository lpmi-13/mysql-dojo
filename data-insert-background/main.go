package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

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
	ctx   = context.TODO()
	delay = 300
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/transactions?charset=utf8mb4&parseTime=True&loc=Local"

	concurrentExecutions, _ := strconv.Atoi(os.Args[1])

	ch := make(chan string)

	for i := 0; i < concurrentExecutions; i++ {
		go insertData(dsn, ch)
	}

	for {
		go insertData(<-ch, ch)
	}
}

func insertData(dsn string, ch chan string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	referral := FakeData{
		Currency:          "dollars",
		Email:             "yep@dep.com",
		ReferralCode:      "292",
		ReferralSubmitted: "yes",
		Referree:          "doobie",
		Referrer:          "scoobie",
		Username:          "yeerrrps",
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	err = db.Table("referrals").Create(referral).Error
	if err != nil {
		log.Fatal(err)
	}

	ch <- dsn
}
