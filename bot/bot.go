package bot

import (
	"github.com/go-microbot/viber/api"
	"github.com/go-microbot/viber/models"
)

// ViberBot represents Viber Bot handler.
// https://developers.viber.com/docs/api/rest-bot-api/#get-started.
type ViberBot struct {
	api        api.Bot
	eventsChan chan models.CallbackEvent
	errChan    chan error
	strategy   EventStrategy
}

// NewViberBot returns new Viber Bot instance.
func NewViberBot(botAPI api.Bot) ViberBot {
	return ViberBot{
		api:        botAPI,
		eventsChan: make(chan models.CallbackEvent, defaultUpdatesLimit),
		errChan:    make(chan error),
	}
}

// API returns bot's API instance.
func (b *ViberBot) API() api.Bot {
	return b.api
}

// WaitForUpdates starts listening of new events using provided strategy.
func (b *ViberBot) WaitForUpdates(strategy EventStrategy) {
	b.strategy = strategy

	go strategy.Listen()

	for {
		select {
		case err, ok := <-strategy.Errors():
			if ok {
				b.errChan <- err
			}
		case event, ok := <-strategy.Callbacks():
			if ok {
				b.eventsChan <- event
			}
		}
	}
}

// Stop stops listening of new events.
func (b *ViberBot) Stop() {
	b.strategy.Stop()
}

// Callbacks returns callback and error channels.
func (b *ViberBot) Callbacks() (upds chan models.CallbackEvent, errs chan error) {
	return b.eventsChan, b.errChan
}
