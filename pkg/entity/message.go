package entity

import "encoding/json"

type Message struct {
	SenderID string `json:"sender_id"`
	To       string `json:"to"`
	Content  string `json:"content"`
}

func (m *Message) ToJson() []byte {
	bytes, _ := json.Marshal(m)
	return bytes
}
