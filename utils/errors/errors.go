package errors

import "fmt"

const (
	LOW = iota
	MIDDLE
	HIGH
)

type Error struct {
	Code  int
	Text  string
	Level int8
}

var Nil Error

func (e Error) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Text)
}

var (
	NoHandlers = Error{
		Code:  0,
		Text:  "No handlers\n\tUse Bot.AddHandler(Handler) to add handlers",
		Level: HIGH,
	}
	NoUpdates = Error{
		Code:  1,
		Text:  "no updates",
		Level: LOW,
	}
	WebhookIsActive = Error{
		Code: 409,
		Text: "can't use getUpdates method while webhook is active\n" +
			"\tUse Bot.DeleteWebhook() to delete the webhook",
		Level: HIGH,
	}
)
