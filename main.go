package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6052149255:AAE1LZCldrWM1_Pz9krcQjAQqPeoIzmgNHo") //(os.Getenv("6052149255:AAE1LZCldrWM1_Pz9krcQjAQqPeoIzmgNHo"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// encode the message using SHA256
		hash := sha256.Sum256([]byte(update.Message.Text))
		encodedMessage := hex.EncodeToString(hash[:])

		// send the encoded message back to the user
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, encodedMessage)
		bot.Send(msg)
	}
}
