package utils

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"

	"github.com/bwmarrin/discordgo"
)

func CheckError(e error, msg ...interface{}) {
	if e != nil {
		log.Fatal(msg)
		return
	}
}

func ArrayContainsUser(arr []*discordgo.User, el *discordgo.User) bool {
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

func LoadDesire() *[]string {
	ret := new([]string)
	b, err := ioutil.ReadFile("pictures/desire.yml")
	CheckError(err, "Couldn't load pictures/desire.yml")
	yaml.Unmarshal(b, ret)
	return ret
}
