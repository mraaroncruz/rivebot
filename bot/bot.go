package bot

import (
	"fmt"
	"regexp"

	rivescript "github.com/aichaos/rivescript-go"
	"github.com/aichaos/rivescript-go/sessions/redis"
	"github.com/pferdefleisch/burgerbot/messenger"
)

type Brain struct {
	bot *rivescript.RiveScript
}

func NewBot(rivescriptDirectoryPath string) *Brain {
	brain := Brain{}

	bot := rivescript.New(rivescript.WithUTF8())
	bot = rivescript.New(&rivescript.Config{
		SessionManager: redis.New(nil),
	})

	err := bot.LoadDirectory(rivescriptDirectoryPath)
	if err != nil {
		fmt.Printf("Error loading from file: %s\n", err)
	}
	bot.SortReplies()

	brain.bot = bot
	return &brain
}

func (brain *Brain) isIntent(text string) bool {
	matched, err := regexp.MatchString("^intent ", text)
	if err != nil {
		return false
	}
	return matched
}

func (brain *Brain) Reply(userID string, message string) (string, error) {
	reply, err := brain.bot.Reply(userID, message)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// if brain.isIntent(reply) {
	// 	return brain.Reply(userID, reply)
	// }

	return reply, nil
}

func (bot *Brain) HandleMessage(userID string, msg string) (string, error) {
	return bot.Reply(userID, msg)
}

func (bot *Brain) setLocation(msg messenger.FBMessage) {
	b := *bot.bot
	coords := msg.Message.Attachments[0].Payload.Coordinates
	b.SetUservar(msg.Sender.ID, "lat", fmt.Sprintf("%f", coords.Lat))
	b.SetUservar(msg.Sender.ID, "lng", fmt.Sprintf("%f", coords.Long))
}

func (bot *Brain) resetTopic(userID string) {
	bot.bot.SetUservar(userID, "topic", "random")
}
