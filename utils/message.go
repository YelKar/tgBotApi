package utils

type Message struct {
	MessageID      int      `json:"message_id"`
	Text           string   `json:"text,omitempty"`
	From           User     `json:"from"`
	Chat           Chat     `json:"chat"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	Sticker        *Sticker `json:"sticker,omitempty"`
	ForwardFrom    *Message `json:"forward_from,omitempty"`
}

func (msg *Message) Type() MessageType {
	var msgType MessageType = 0
	if msg.ForwardFrom != nil {
		msgType += ForwardedMessage
	}
	if msg.ReplyToMessage != nil {
		msgType += ReplyToMessage
	}
	if msg.Sticker != nil {
		msgType += StickerMessage
	}
	if msgType == 0 {
		msgType += InputMessage
	}
	return msgType
}

type SentMessage struct {
	ChatID    int         `json:"chat_id"`
	Text      string      `json:"text"`
	ParseMode string      `json:"parse_mode,omitempty"`
	Keyboard  interface{} `json:"reply_markup,omitempty"`
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
	Callback
	PollMessage
	StickerMessage
	TextMessage = InputMessage | EditedMessage | ForwardedMessage | ReplyToMessage
)
