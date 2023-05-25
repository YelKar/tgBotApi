package utils

type TGResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
	JSON string `json:",omitempty"`
}

type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}

type Query struct {
	Offset  int `json:"offset,omitempty"`
	Limit   int `json:"limit,omitempty"`
	Timeout int `json:"timeout,omitempty"`
}
