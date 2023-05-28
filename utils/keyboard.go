package utils

type Keyboard struct {
	Keyboard        []KeyboardRow `json:"keyboard"`
	ResizeKeyboard  bool          `json:"resize_keyboard"`
	OneTimeKeyboard bool          `json:"one_time_keyboard"`
	// input_field_placeholder string
}
type RemoveKeyboard struct {
	Remove    bool `json:"remove_keyboard"`
	Selective bool `json:"selective,omitempty"`
}
type KeyboardRow []KeyboardButton

type KeyboardButton struct {
	Text string `json:"text"`
}

func (kb *Keyboard) Add(rows ...KeyboardRow) {
	kb.Keyboard = append(kb.Keyboard, rows...)
}
func (kbr *KeyboardRow) Add(buttons ...KeyboardButton) {
	*kbr = append(*kbr, buttons...)
}

type InlineKeyboardRow []InlineKeyboardButton
type InlineKeyboard struct {
	Keyboard []InlineKeyboardRow `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url"`
	CallbackData string `json:"callback_data"`
}

func (kb *InlineKeyboard) Add(rows ...InlineKeyboardRow) {
	kb.Keyboard = append(kb.Keyboard, rows...)
}
func (kbr *InlineKeyboardRow) Add(buttons ...InlineKeyboardButton) {
	*kbr = append(*kbr, buttons...)
}
