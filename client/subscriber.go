package client

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"lab/golang-mqtt-chat-engine/config"
	"lab/golang-mqtt-chat-engine/models"
	"lab/golang-mqtt-chat-engine/utils"
	"time"
)

type Subscriber struct {
	Client     mqtt.Client
	Username   string
	RuleEngine *utils.RuleEngine
}

func NewSubscriber(username string) *Subscriber {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.AppConfig.Broker.URL)
	opts.SetClientID(username + "_sub")
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return &Subscriber{
		Client:     c,
		Username:   username,
		RuleEngine: &utils.RuleEngine{},
	}
}

func (s *Subscriber) Subscribe(topic string) {
	token := s.Client.Subscribe(topic, config.AppConfig.Broker.QoS, func(client mqtt.Client, msg mqtt.Message) {
		var chatMsg models.ChatMessage
		_ = json.Unmarshal(msg.Payload(), &chatMsg)

		// Ephemeral message TTL
		if chatMsg.Type == "ephemeral" {
			elapsed := time.Since(chatMsg.Timestamp)
			if elapsed.Seconds() > float64(config.AppConfig.Broker.EphemeralTTL) {
				return
			}
		}

		// Typing indicator
		if chatMsg.Type == "typing" {
			fmt.Printf("[%s] %s is typing...\n", chatMsg.Topic, chatMsg.Sender)
			return
		}

		// Rule engine
		if !s.RuleEngine.ProcessMessage(chatMsg) {
			return
		}

		fmt.Printf("[%s][%s] %s: %s\n", chatMsg.Topic, chatMsg.Type, chatMsg.Sender, chatMsg.Payload)
	})
	token.Wait()
}

func (s *Subscriber) Disconnect() {
	s.Client.Disconnect(250)
}
