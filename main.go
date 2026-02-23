package main

import (
	"bufio"
	"fmt"
	"lab/golang-mqtt-chat-engine/client"
	"lab/golang-mqtt-chat-engine/utils"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	chatClient := client.NewClient(username)
	utils.Info("%s connected to EMQX broker", username)

	// Subscribe to a general topic
	chatClient.Subscribe("chat/general")

	fmt.Println("Type messages to send (type 'exit' to quit):")
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		chatClient.Publish("chat/general", fmt.Sprintf("%s: %s", username, text), false)
	}

	chatClient.Disconnect()
	utils.Info("%s disconnected", username)
}
