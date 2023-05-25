package utils

type User struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
