package main

import (
	"flag"
)

var (
	BotToken = flag.String("t", "", "Bot authorization token")
	//GuildID   = flag.String("g", "", "ID of the testing guild")
	ChannelID = flag.String("c", "", "ID of the testing channel")
)

func main() {
	flag.Parse()

	InitDB()
	BotInit()
	ServeAPI()

	//var wg sync.WaitGroup
	//
	//go func() {
	//	defer wg.Done()
	//	BotInit()
	//}()
	//wg.Add(1)
	//
	//go func() {
	//	defer wg.Done()
	//	ServeAPI()
	//}()
	//wg.Add(1)
	//
	//wg.Wait()

}
