package models

import "time"

type ChatMessage struct {
	Sender    string    `json:"sender"`
	Topic     string    `json:"topic"` // group or private
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"` // normal / ephemeral / typing / offline
}
