package utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func CheckError(e error, msg ...interface{}) {
	if e != nil {
		log.Fatal(msg)
		return
	}
}

func MentionsContainsUser(arr []*discordgo.User, el *discordgo.User) bool {
	if len(arr) == 0 {
		return false
	}

	for _, element := range arr {
		if element.ID == el.ID {
			return true
		}
	}

	return false
}
