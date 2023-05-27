package utils

type Message struct {
	MessageID      int      `json:"message_id"`
	Text           string   `json:"text"`
	From           User     `json:"from"`
	Chat           Chat     `json:"chat"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	Sticker        *Sticker `json:"sticker"`
	ForwardFrom    *Message `json:"forward_from"`
}

type SentMessage struct {
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

type Sticker struct {
	Name         string `json:"set_name"`
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Emoji        string `json:"emoji"`
	IsAnimated   bool   `json:"is_animated"`
	IsVideo      bool   `json:"is_video"`
}

type MessageType int

const (
	InputMessage = 1 << iota
	EditedMessage
	ForwardedMessage
	ReplyToMessage
	CallbackQuery
	PollMessage
	StickerMessage
	TextMessage = InputMessage + EditedMessage + ForwardedMessage + ReplyToMessage
)
