package client

import "time"

type ChatMessage struct {
	Sender    string    `json:"sender"`
	Topic     string    `json:"topic"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}
