package bot

import "github.com/go-microbot/viber/models"

const defaultUpdatesLimit = 20

// EventStrategy represents a strategy interface
// for determining how a Bot will receive new messages.
type EventStrategy interface {
	Callbacks() chan models.CallbackEvent
	Errors() chan error
	Listen()
	Stop()
}
