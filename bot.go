package tgBotApi

import (
	"bytes"
	"fmt"
	"github.com/YelKar/tgBotApi/utils"
	"github.com/YelKar/tgBotApi/utils/errors"
	"net/http"
)

var url = "https://api.telegram.org/bot%s/%s"
var lastUpdate = 0

type Bot struct {
	TOKEN            string
	MessageHandlers  MessageHandlers
	CallbackHandlers CallbackHandlers
	stopChannel      chan struct{}
}
type MessageHandlers []MessageHandler
type CallbackHandlers []CallbackHandler

func (mh *MessageHandlers) Add(handler MessageHandler) {
	*mh = append(*mh, handler)
}
func (ch *CallbackHandlers) Add(handler CallbackHandler) {
	*ch = append(*ch, handler)
}

type MessageHandler struct {
	Filter      func(*utils.Message) bool
	Handler     func(*Bot, *utils.Message)
	MessageType utils.MessageType
}

type CallbackHandler struct {
	Filter  func(*utils.CallbackQuery) bool
	Handler func(*Bot, *utils.CallbackQuery)
}

func (bot *Bot) AddMessageHandler(handler MessageHandler) {
	bot.MessageHandlers = append(bot.MessageHandlers, handler)
}

func (bot *Bot) AddCallbackHandler(handler CallbackHandler) {
	bot.CallbackHandlers = append(bot.CallbackHandlers, handler)
}

func (bot *Bot) Get() (utils.Update, errors.Error) {
	resp := bot.GetUpdates()
	if len(resp.Result) > 0 {
		lastUpdate = resp.Result[0].ID
		return resp.Result[0], errors.Nil
	} else if resp.ErrorCode == 409 {
		return utils.Update{}, errors.WebhookIsActive
	}
	return utils.Update{}, errors.NoUpdates
}

func (bot *Bot) SetWebhook(webhookUrl string) {
	_, err := http.Post(
		fmt.Sprintf(url, bot.TOKEN, "setWebhook"),
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf(`{"url": "%s"}`, webhookUrl))),
	)
	if err != nil {
		panic(err)
	}
}

func (bot *Bot) DeleteWebhook(webhookUrl string) {
	_, err := http.Post(
		fmt.Sprintf(url, bot.TOKEN, "setWebhook"),
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf(`{"url": "%s"}`, webhookUrl))),
	)
	if err != nil {
		panic(err)
	}
}
