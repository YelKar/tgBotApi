package utils

type CallbackQuery struct {
	ID      int64
	From    User
	Message Message
	Data    string
}
