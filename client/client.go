package client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

// ChatClient struct
type ChatClient struct {
	Client   mqtt.Client
	Username string
}

// NewClient creates MQTT client with LWT
func NewClient(username string) *ChatClient {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID(username)
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)
	opts.SetWill("chat/general", fmt.Sprintf("%s is offline", username), 1, false)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return &ChatClient{
		Client:   client,
		Username: username,
	}
}

// Subscribe to a topic
func (c *ChatClient) Subscribe(topic string) {
	token := c.Client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("[%s] %s: %s\n", topic, msg.Topic(), string(msg.Payload()))
	})
	token.Wait()
}

// Publish a message
func (c *ChatClient) Publish(topic, msg string, retained bool) {
	token := c.Client.Publish(topic, 1, retained, msg)
	token.Wait()
}

// Disconnect client
func (c *ChatClient) Disconnect() {
	c.Client.Disconnect(250)
}
