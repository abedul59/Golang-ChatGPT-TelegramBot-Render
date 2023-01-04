package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/yanzay/tbot/v2"
)

func getChatGPTresponse(ctx context.Context, question string) string {
	c := gogpt.NewClient(os.Getenv("OPENAI_TOKEN"))

	maxtokens, err0 := strconv.Atoi(os.Getenv("OPENAI_MAXTOKENS"))

	if err0 != nil {
		fmt.Println("Error during conversion")
		return "MaxTokens Conversion Error happened!"
	}

	req := gogpt.CompletionRequest{
		Model:       "text-davinci-003",
		MaxTokens:   maxtokens,
		Prompt:      question,
		Temperature: 0,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "You got an error!"
	} else {
		fmt.Println(resp.Choices[0].Text)

		return resp.Choices[0].Text
	}

}

func main() {
	ctx := context.Background()
	bot := tbot.New(os.Getenv("TELEGRAM_BOT_TOKEN"),
		tbot.WithWebhook("https://chatgpt-telebot.onrender.com", ":8080"))
	c := bot.Client()

	bot.HandleMessage(".*human:*", func(m *tbot.Message) {
		c.SendMessage(m.Chat.ID, "AI:" + getChatGPTresponse(ctx, m.Text))
	})
	log.Fatal(bot.Start())
}
