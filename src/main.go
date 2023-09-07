package main

import (
	"flag"
)

var (
	// BotToken Discord bot token
	BotToken = flag.String("t", "", "Bot authorization token")
	// ChannelID ID of the announcement channel
	ChannelID = flag.String("c", "", "ID of the testing channel")
)

func main() {
	flag.Parse()

	InitDB()
	BotInit()
	ServeAPI()
}
