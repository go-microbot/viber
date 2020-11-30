package main

import (
	"context"
	"fmt"

	"github.com/go-microbot/telegram/api"
	apiModels "github.com/go-microbot/telegram/api/models"
	"github.com/go-microbot/telegram/bot"
	"github.com/go-microbot/telegram/query"
)

const telegramBotToken = "<PASTE_YOUR_TOKEN_HERE>"

func main() {
	// init Bot API with token.
	botAPI := api.NewTelegramAPI(telegramBotToken)

	// create Bot instance.
	myBot := bot.NewTelegramBot(&botAPI)

	// delete webhook (if it was using before).
	if err := myBot.API().DeleteWebhook(context.Background()); err != nil {
		fmt.Printf("could not remove webhook: %v", err)
	}

	// start long polling.
	go myBot.WaitForUpdates(bot.NewUpdatesStrategyLongPolling(bot.LongPollingConfig{
		Timeout: 10,
		BotAPI:  &botAPI,
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
