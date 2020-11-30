package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/go-microbot/telegram/api"
	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/bot"
	"github.com/go-microbot/telegram/form"
	"github.com/go-microbot/telegram/query"
)

const telegramBotToken = "<PASTE_YOUR_TOKEN_HERE>"

func main() {
	// init Bot API with token.
	botAPI := api.NewTelegramAPI(telegramBotToken)

	// create Bot instance.
	myBot := bot.NewTelegramBot(&botAPI)

	// read certificate data.
	data, err := ioutil.ReadFile("telegram_test.key")
	if err != nil {
		panic(err)
	}

	// set webhook.
	req := apiModels.SetWebhookRequest{
		Certificate: form.NewPartText(string(data)),
		URL:         query.NewParamString("https://53ec7fc0c840.ngrok.io"), // ngrok, you need to use your server URL.
	}
	err = myBot.API().SetWebhook(context.Background(), req)
	if err != nil {
		panic(err)
	}

	// start listening.
	go myBot.WaitForUpdates(bot.NewUpdatesStrategyWebhook(bot.WebhookConfig{
		ServeURL: "localhost:8443", // server to catch Telegram requests.
	}))

	// listen Bot's updates.
	updates, errs := myBot.Updates()
	for {
		select {
		case update, ok := <-updates:
			if !ok {
				fmt.Println("updates channel closed")
				return
			}

			// reply "hello" message.
			_, err := myBot.API().SendMessage(context.Background(), apiModels.SendMessageRequest{
				ChatID:           query.NewParamAny(update.Message.Chat.ID),
				Text:             fmt.Sprintf("Hello, %s!", update.Message.From.Username),
				ReplyToMessageID: &update.Message.ID,
			})
			if err != nil {
				panic(err)
			}
		case err, ok := <-errs:
			if !ok {
				fmt.Println("errors channel closed")
				return
			}
			fmt.Println(err)
		}
	}
}
