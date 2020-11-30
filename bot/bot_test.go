package bot

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/go-microbot/viber/api"
	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

var errMock = errors.New("error")

type testStrategy struct {
	fn     func(strategy *testStrategy)
	events chan models.CallbackEvent
	errs   chan error
}

func (t testStrategy) Listen() {
	if t.fn != nil {
		t.fn(&t)
	}
}

func (t testStrategy) Callbacks() chan models.CallbackEvent {
	return t.events
}

func (t testStrategy) Errors() chan error {
	return t.errs
}

func (t testStrategy) Stop() {
	close(t.events)
	close(t.errs)
}

func newTestStrategy(fn func(strategy *testStrategy)) testStrategy {
	return testStrategy{
		fn:     fn,
		events: make(chan models.CallbackEvent),
		errs:   make(chan error),
	}
}

func Test_NewViberBot(t *testing.T) {
	botAPI := api.NewViberAPI("")
	require.NotNil(t, botAPI)
	bot := NewViberBot(&botAPI)
	require.NotNil(t, bot)
	require.NotNil(t, bot.api)
	require.NotNil(t, bot.eventsChan)
	require.NotNil(t, bot.errChan)
}

func TestViberBot_API(t *testing.T) {
	botAPI := api.NewViberAPI("")
	require.NotNil(t, botAPI)
	bot := NewViberBot(&botAPI)
	require.NotNil(t, bot.api)
	require.Equal(t, &botAPI, bot.API())
}

func TestViberBot_WaitForUpdates(t *testing.T) {
	botAPI := api.NewViberAPI("")
	require.NotNil(t, botAPI)
	bot := NewViberBot(&botAPI)
	require.NotNil(t, bot.api)
	var wg sync.WaitGroup
	wg.Add(1)
	strategy := newTestStrategy(func(strategy *testStrategy) {
		time.Sleep(time.Second)
		strategy.events <- models.CallbackEvent{
			Event: "message",
			Message: &models.Message{
				Text: "first message",
			},
		}
		time.Sleep(time.Second)

		strategy.errs <- errors.New("some error")
		time.Sleep(time.Second)

		strategy.events <- models.CallbackEvent{
			Event: "message",
			Message: &models.Message{
				Text: "second message",
			},
		}
		time.Sleep(time.Second)
		wg.Done()
	})

	messages := make([]string, 0)
	errs := make([]error, 0)
	go func() {
		for {
			select {
			case msg, ok := <-bot.eventsChan:
				require.True(t, ok)
				require.NotNil(t, msg)
				require.NotNil(t, msg.Message)
				messages = append(messages, msg.Message.Text)
			case err, ok := <-bot.errChan:
				require.True(t, ok)
				require.Error(t, err)
				errs = append(errs, err)
			}
		}
	}()

	go bot.WaitForUpdates(strategy)
	wg.Wait()

	expErrors := []error{
		errors.New("some error"),
	}
	expMessages := []string{"first message", "second message"}
	require.Equal(t, expMessages, messages)
	require.Equal(t, expErrors, errs)
}

func TestViberBot_Stop(t *testing.T) {
	botAPI := api.NewViberAPI("")
	require.NotNil(t, botAPI)
	bot := NewViberBot(&botAPI)
	require.NotNil(t, bot.api)
	var wg sync.WaitGroup
	wg.Add(1)
	strategy := newTestStrategy(func(strategy *testStrategy) {
		_, ok := <-strategy.events
		require.False(t, ok)
		wg.Done()
	})
	go bot.WaitForUpdates(strategy)
	time.Sleep(time.Second * 3)
	bot.Stop()
	wg.Wait()
}

func TestViberBot_Callbacks(t *testing.T) {
	botAPI := api.NewViberAPI("")
	require.NotNil(t, botAPI)
	bot := NewViberBot(&botAPI)
	require.NotNil(t, bot.api)
	events, errs := bot.Callbacks()
	require.NotNil(t, events)
	require.NotNil(t, errs)
	require.Equal(t, bot.eventsChan, events)
	require.Equal(t, bot.errChan, errs)
}
