package utils

type Callback struct {
	ID      int64
	From    User
	Message Message
	Data    string
}
