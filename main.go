package main

import (
	"bufio"
	"fmt"
	"lab/golang-mqtt-chat-engine/client"
	"lab/golang-mqtt-chat-engine/config"
	"lab/golang-mqtt-chat-engine/utils"
	"os"
	"strings"
)

func main() {
	// Load YAML config
	config.LoadConfig("config.yml")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	chatClient := client.NewClient(username)
	utils.Info("%s connected to EMQX broker", username)

	currentRoom := config.AppConfig.Broker.DefaultRoom
	chatClient.Subscribe(currentRoom)

	fmt.Println("Commands: /room <name> | /ephemeral <message> | /typing | exit")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		if strings.HasPrefix(input, "/room") {
			parts := strings.Split(input, " ")
			if len(parts) > 1 {
				currentRoom = "chat/" + parts[1]
				chatClient.Subscribe(currentRoom)
				utils.Info("Switched to room: %s", currentRoom)
			}
			continue
		}

		if strings.HasPrefix(input, "/ephemeral") {
			msg := strings.TrimPrefix(input, "/ephemeral ")
			chatClient.Publish(currentRoom, msg, "ephemeral", false)
			continue
		}

		if input == "/typing" {
			chatClient.Publish(currentRoom, "", "typing", false)
			continue
		}

		// Normal message
		chatClient.Publish(currentRoom, input, "normal", false)
	}

	chatClient.Disconnect()
	utils.Info("%s disconnected", username)
}
