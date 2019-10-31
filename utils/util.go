package utils

import (
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
	yaml "gopkg.in/yaml.v2"
)

// CheckError will check if the error is nil and if not log the error
func CheckError(e error, msg ...interface{}) {
	if e != nil {
		log.Fatal(msg...)
		return
	}
}

// ArrayContainsUser will check whether a given user is contained in a list of discord users
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

// LoadDesire will load all links to a desire picture defined in `pictures/desire.yml`
func LoadDesire() *[]string {
	ret := new([]string)
	b, err := ioutil.ReadFile("pictures/desire.yml")
	CheckError(err, "Couldn't load pictures/desire.yml")
	yaml.Unmarshal(b, ret)
	return ret
}
