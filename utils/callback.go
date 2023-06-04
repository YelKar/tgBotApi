package utils

type CallbackQuery struct {
	ID      int64   `json:"id"`
	From    User    `json:"from"`
	Message Message `json:"message"`
	Data    string  `json:"data"`
}
