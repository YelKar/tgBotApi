package tg

import (
	"bytes"
	"fmt"
	"github.com/YelKar/tgBotApi/utils"
	"net/http"
)

var url = "https://api.telegram.org/bot%s/%s"
var lastUpdate = 0

type Bot struct {
	TOKEN       string
	Handlers    []Handler
	stopChannel chan struct{}
}

type Handler struct {
	Filter  func(string) bool
	Handler func(*Bot, utils.Message)
}

func (bot *Bot) AddHandler(handler Handler) {
	bot.Handlers = append(bot.Handlers, handler)
}

func (bot *Bot) Get() (utils.Update, bool) {
	resp := bot.GetUpdates()

	if len(resp.Result) > 0 {
		lastUpdate = resp.Result[0].ID
		return resp.Result[0], true
	}
	return utils.Update{}, false
}

func (bot *Bot) SetWebhook(webhookUrl string) {
	http.Post(
		fmt.Sprintf(url, bot.TOKEN, "setWebhook"),
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf(`{"url": "%s"}`, webhookUrl))),
	)
}

func (bot *Bot) DeleteWebhook(webhookUrl string) {
	http.Post(
		fmt.Sprintf(url, bot.TOKEN, "setWebhook"),
		"application/json",
		bytes.NewReader([]byte(fmt.Sprintf(`{"url": "%s"}`, webhookUrl))),
	)
}
