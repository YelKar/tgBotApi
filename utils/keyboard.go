package utils

type Keyboard struct {
	Keyboard        []KeyboardRow `json:"keyboard"`
	ResizeKeyboard  bool          `json:"resize_keyboard"`
	OneTimeKeyboard bool          `json:"one_time_keyboard"`
	// input_field_placeholder string
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

type InlineKeyboard struct {
}

type InlineKeyboardButton struct {
}
