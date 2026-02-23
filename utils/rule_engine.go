package utils

import (
	"fmt"
	"lab/golang-mqtt-chat-engine/config"
	"lab/golang-mqtt-chat-engine/models"
	"strings"
)

type RuleEngine struct{}

func (r *RuleEngine) ProcessMessage(msg models.ChatMessage) bool {
	for _, word := range config.AppConfig.BannedWords {
		if strings.Contains(strings.ToLower(msg.Payload), word) {
			fmt.Printf("[RULE] Message blocked (banned word): %s\n", msg.Payload)
			return false
		}
	}
	// Log for analytics (simulated)
	fmt.Printf("[RULE] Logging message: %s in topic %s\n", msg.Payload, msg.Topic)
	return true
}
