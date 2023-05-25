package utils

type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
	From      User   `json:"from"`
	Chat      Chat   `json:"chat"`
}

type SentMessage struct {
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}
