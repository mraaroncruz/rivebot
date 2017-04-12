package main

import (
	"flag"
	"fmt"

	"github.com/pferdefleisch/rivebot/bot"
)

var rsDirectory = flag.String("rs", "./rivescripts", "Path to rivescript files directory")
var msg = flag.String("message", "", "Message for bot to parse")
var userID = flag.String("sessionid", "", "Session id for bot")

func main() {
	flag.Parse()
	brain := bot.NewBot(*rsDirectory)
	reply, err := brain.Reply(*userID, *msg)
	if err != nil {
		fmt.Printf("error %s\n", err)
		return
	}
	fmt.Println(reply)
}
