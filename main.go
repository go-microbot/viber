package main

import (
	"context"
	"fmt"

	"github.com/go-microbot/viber/api"
	"github.com/go-microbot/viber/bot"
)

func main() {
	const token = 

	// init Bot API with token.
	botAPI := api.NewViberAPI(token)

	info, err := botAPI.GetAccountInfo(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	return

	// create Bot instance.
	myBot := bot.NewViberBot(&botAPI)

	// start listening.
	go myBot.WaitForUpdates(bot.NewWebhookStrategy(bot.WebhookConfig{
		ServeURL: "localhost:8443", // server to catch Telegram requests.
	}))

	// listen Bot's events.
	events, errs := myBot.Callbacks()
	for {
		select {
		case event, ok := <-events:
			if !ok {
				fmt.Println("updates channel closed")
				return
			}

			fmt.Println(event)
			/*// reply "hello" message.
			_, err := myBot.API().SendMessage(context.Background(), apiModels.SendMessageRequest{
				ChatID:           query.NewParamAny(update.Message.Chat.ID),
				Text:             fmt.Sprintf("Hello, %s!", update.Message.From.Username),
				ReplyToMessageID: &update.Message.ID,
			})
			if err != nil {
				panic(err)
			}*/
		case err, ok := <-errs:
			if !ok {
				fmt.Println("errors channel closed")
				return
			}
			fmt.Println(err)
		}
	}
}
