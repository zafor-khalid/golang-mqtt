package main

import (
	"flag"
	"fmt"
	"lab/golang-mqtt-chat-engine/client"
	"lab/golang-mqtt-chat-engine/config"
	"os"
)

func main() {
	config.LoadConfig("config.yml")

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go pub --user=<name>")
		fmt.Println("  go run main.go sub --user=<name> --topic=<topic>")
		return
	}

	mode := os.Args[1]

	switch mode {

	case "pub":
		runPublisher()

	case "sub":
		runSubscriber()

	default:
		fmt.Println("Unknown mode. Use 'pub' or 'sub'")
	}
}

func runPublisher() {
	user := flag.String("user", "guest", "username")
	topic := flag.String("topic", "general", "default topic")
	flag.CommandLine.Parse(os.Args[2:])

	pub := client.NewPublisher(*user)
	defer pub.Disconnect()

	fmt.Println("Publisher started")
	fmt.Println("Commands:")
	fmt.Println("/topic <name>")
	fmt.Println("/ephemeral <msg>")
	fmt.Println("/typing")
	fmt.Println("exit")

	currentTopic := *topic

	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		if input == "/typing" {
			pub.Publish(currentTopic, "", "typing", false)
			continue
		}

		if input == "/ephemeral" {
			var msg string
			fmt.Scanln(&msg)
			pub.Publish(currentTopic, msg, "ephemeral", false)
			continue
		}

		if input == "/topic" {
			var newTopic string
			fmt.Scanln(&newTopic)
			currentTopic = newTopic
			fmt.Println("Switched topic:", currentTopic)
			continue
		}

		pub.Publish(currentTopic, input, "normal", true)
	}
}

func runSubscriber() {
	user := flag.String("user", "guest", "username")
	topic := flag.String("topic", "general", "topic to subscribe")
	flag.CommandLine.Parse(os.Args[2:])

	sub := client.NewSubscriber(*user)
	defer sub.Disconnect()

	fmt.Println("Subscriber started")
	fmt.Println("Listening on topic:", *topic)

	sub.Subscribe(*topic)

	select {} // block forever
}
