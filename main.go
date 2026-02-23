package main

import (
	"bufio"
	"flag"
	"fmt"
	"lab/golang-mqtt-chat-engine/client"
	"lab/golang-mqtt-chat-engine/config"
	"os"
	"strings"
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

	fmt.Println("type exit to exit")

	currentTopic := *topic

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
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
