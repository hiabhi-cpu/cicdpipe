package mailing

import (
	"log"
	"os"
	"strings"

	mailrepo "github.com/hiabhi-cpu/mailing_repo/mail_repo"
	"github.com/joho/godotenv"
)

func Mailing(subject, body string, filePath []string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	from := os.Getenv("GMAIL")
	password := os.Getenv("MAIL_PASS") // 16-character app password
	tomail := strings.Split(os.Getenv("TO_MAIL"), ",")

	mailrepo.Mail(from, tomail, []string{}, password, subject, body, filePath)
}
