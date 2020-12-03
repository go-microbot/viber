package bot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/go-microbot/viber/models"
	"github.com/stretchr/testify/require"
)

func Test_NewWebhookStrategy(t *testing.T) {
	hook := NewWebhookStrategy(WebhookConfig{})
	require.NotNil(t, hook)
	require.NotNil(t, hook.cfg)
	require.NotNil(t, hook.errorsChan)
	require.NotNil(t, hook.eventsChan)
	require.NotNil(t, hook.doneChan)
}

func TestStrategyWebhook_Callbacks(t *testing.T) {
	hook := NewWebhookStrategy(WebhookConfig{})
	require.NotNil(t, hook)
	require.NotNil(t, hook.eventsChan)
	require.Equal(t, hook.eventsChan, hook.Callbacks())
}

func TestStrategyWebhook_Errors(t *testing.T) {
	hook := NewWebhookStrategy(WebhookConfig{})
	require.NotNil(t, hook)
	require.NotNil(t, hook.errorsChan)
	require.Equal(t, hook.errorsChan, hook.Errors())
}

func TestStrategyWebhook_Listen(t *testing.T) {
	hook := NewWebhookStrategy(WebhookConfig{
		ServeURL: "localhost:8443",
	})
	require.NotNil(t, hook)

	errs := make([]error, 0)
	messages := make([]string, 0)
	go func() {
		for {
			select {
			case msg, ok := <-hook.eventsChan:
				require.True(t, ok)
				require.NotNil(t, msg)
				require.NotNil(t, msg.Message)
				messages = append(messages, msg.Message.Text)
			case err, ok := <-hook.errorsChan:
				require.True(t, ok)
				require.Error(t, err)
				errs = append(errs, err)
			}
		}
	}()

	go hook.Listen()

	url := "http://" + hook.cfg.ServeURL
	// message 1.
	data, err := json.Marshal(models.CallbackEvent{
		Event: "message",
		Message: &models.Message{
			Text: "message 1",
		},
	})
	require.NoError(t, err)
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	require.NoError(t, err)
	closeBody(t, resp.Body)

	// error 1.
	require.NoError(t, err)
	resp, err = http.Post(url, "application/json", bytes.NewReader([]byte("invalid")))
	require.NoError(t, err)
	closeBody(t, resp.Body)

	// message 2.
	data, err = json.Marshal(models.CallbackEvent{
		Event: "message",
		Message: &models.Message{
			Text: "message 2",
		},
	})
	require.NoError(t, err)
	resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	require.NoError(t, err)
	closeBody(t, resp.Body)

	// error 2.
	require.NoError(t, err)
	resp, err = http.Post(url, "application/json", bytes.NewReader([]byte("data invalid")))
	require.NoError(t, err)
	closeBody(t, resp.Body)

	expMessages := []string{"message 1", "message 2"}
	require.Equal(t, expMessages, messages)
	require.Equal(t, 2, len(errs))
	for i := range errs {
		_, ok := errs[i].(*json.SyntaxError)
		require.True(t, ok)
	}
}

func TestStrategyWebhook_Stop(t *testing.T) {
	hook := NewWebhookStrategy(WebhookConfig{
		ServeURL: "localhost:8440",
	})
	require.NotNil(t, hook)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, ok := <-hook.eventsChan
		require.False(t, ok)
		wg.Done()
	}()

	go hook.Listen()
	time.Sleep(time.Second * 2)
	hook.Stop()
	wg.Wait()
}

func Test_validateMAC(t *testing.T) {
	t.Run("parse signature error", func(t *testing.T) {
		err := validateMAC(nil, "TT", "key")
		require.Error(t, err)
		require.EqualError(t, err, fmt.Errorf("could not decode signature: %v", errors.New("encoding/hex: invalid byte: U+0054 'T'")).Error())
	})
	t.Run("mac is invalid error", func(t *testing.T) {
		err := validateMAC(nil, "", "key")
		require.Error(t, err)
		require.EqualError(t, err, fmt.Errorf("mac is invalid").Error())
	})
	t.Run("all ok", func(t *testing.T) {
		key := "1234"
		message := []byte("test")

		mac := hmac.New(sha256.New, []byte(key))
		_, err := mac.Write(message)
		require.NoError(t, err)

		signature := hex.EncodeToString(mac.Sum(nil))

		err = validateMAC(message, signature, key)
		require.NoError(t, err)
	})
}

func closeBody(t *testing.T, body io.ReadCloser) {
	err := body.Close()
	require.NoError(t, err)
}
