package main

import (
	"log"
	"os"
	"github.com/yanzay/tbot/v2"
)


func main() {
	bot := tbot.New(os.Getenv("TELEGRAM_BOT_TOKEN"),
		tbot.WithWebhook("https://chatgpt-telebot.onrender.com", ":8080"))
	c := bot.Client()
	bot.HandleMessage(".*yo.*", func(m *tbot.Message) {
	c.SendMessage(m.Chat.ID, "hello!")
	})
	log.Fatal(bot.Start())
}
