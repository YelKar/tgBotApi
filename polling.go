package tg

import (
	"fmt"
	"github.com/YelKar/tgBotApi/utils"
)

func (bot *Bot) Polling() <-chan struct{} {
	stop := make(chan struct{})
	bot.stopChannel = stop
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Polling stopped")
				return
			default:
				update, ok := bot.Get()
				if ok {
					bot.processUpdate(update)
				}
			}
		}
	}()
	return stop
}

func (bot *Bot) StopPolling() {
	bot.GetUpdates()
	bot.stopChannel <- struct{}{}
}

func (bot *Bot) WaitPolling() {
	<-bot.stopChannel
}

func (bot *Bot) processUpdate(update utils.Update) {
	for _, handler := range bot.Handlers {
		if handler.Filter(update.Message.Text) {
			handler.Handler(bot, update.Message)
			break
		}
	}
}
