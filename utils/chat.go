package utils

type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
