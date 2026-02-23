package client

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"lab/golang-mqtt-chat-engine/config"
	"lab/golang-mqtt-chat-engine/models"
	"time"
)

type Publisher struct {
	Client   mqtt.Client
	Username string
}

func NewPublisher(username string) *Publisher {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.AppConfig.Broker.URL)
	opts.SetClientID(username + "_pub")
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return &Publisher{
		Client:   c,
		Username: username,
	}
}

func (p *Publisher) Publish(topic, payload, msgType string, retained bool) {
	msg := models.ChatMessage{
		Sender:    p.Username,
		Topic:     topic,
		Payload:   payload,
		Timestamp: time.Now(),
		Type:      msgType,
	}
	data, _ := json.Marshal(msg)
	token := p.Client.Publish(topic, config.AppConfig.Broker.QoS, retained, data)
	token.Wait()
}

func (p *Publisher) Disconnect() {
	p.Client.Disconnect(250)
}
