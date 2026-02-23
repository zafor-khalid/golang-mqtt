package client

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"lab/golang-mqtt-chat-engine/config"
	"lab/golang-mqtt-chat-engine/utils"
	"time"
)

type ChatClient struct {
	Client   mqtt.Client
	Username string
}

// NewClient connects to EMQX with LWT
func NewClient(username string) *ChatClient {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.AppConfig.Broker.URL)
	opts.SetClientID(username)
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)
	opts.SetWill(config.AppConfig.Broker.DefaultRoom,
		fmt.Sprintf("%s has disconnected unexpectedly", username),
		config.AppConfig.Broker.QoS,
		false,
	)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return &ChatClient{
		Client:   client,
		Username: username,
	}
}

// Subscribe to a chat room
func (c *ChatClient) Subscribe(room string) {
	token := c.Client.Subscribe(room, config.AppConfig.Broker.QoS, func(client mqtt.Client, msg mqtt.Message) {
		var chatMsg ChatMessage
		err := json.Unmarshal(msg.Payload(), &chatMsg)
		if err != nil {
			utils.Error("Failed to parse message: %v", err)
			return
		}

		// Ephemeral messages: ignore if TTL expired
		if chatMsg.Type == "ephemeral" {
			elapsed := time.Since(chatMsg.Timestamp)
			if elapsed.Seconds() > float64(config.AppConfig.Broker.EphemeralTTL) {
				return
			}
		}

		fmt.Printf("[%s][%s] %s: %s\n", chatMsg.Room, chatMsg.Type, chatMsg.Sender, chatMsg.Payload)
	})
	token.Wait()
}

// Publish a message
func (c *ChatClient) Publish(room, text, msgType string, retained bool) {
	chatMsg := ChatMessage{
		Sender:    c.Username,
		Room:      room,
		Payload:   text,
		Timestamp: time.Now(),
		Type:      msgType,
	}

	payload, _ := json.Marshal(chatMsg)
	token := c.Client.Publish(room, config.AppConfig.Broker.QoS, retained, payload)
	token.Wait()
}

// Disconnect client
func (c *ChatClient) Disconnect() {
	c.Client.Disconnect(250)
}
