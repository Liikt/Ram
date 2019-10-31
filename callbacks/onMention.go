package callbacks

import (
	"strings"

	f "../functions"
	utils "../utils"
	"github.com/bwmarrin/discordgo"
)

var (
	commChan  = make(chan string)
	recording = false
)

// OnMessage will be executed every time the discord session will receive a message
func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	if m.ChannelID == "276583899088289793" {
		s.MessageReactionAdd(m.ChannelID, m.Message.ID, "🐀")
	}

	f.Shitpost(s, m)

	if utils.ArrayContainsUser(m.Mentions, s.State.User) {
		c, err := s.State.Channel(m.ChannelID)
		utils.CheckError(err, "Couldn't get the Channel with the ID: ", m.ChannelID)

		_, err = s.State.Guild(c.GuildID)
		utils.CheckError(err, "Couldn't get the Server with the ID: ", c.GuildID)

		split := strings.SplitN(m.ContentWithMentionsReplaced(), " ", 3)
		command, line := split[1], ""

		if len(split) == 3 {
			line = split[2]
		}

		switch command {
		case "echo":
			f.Echo(s, m, line)
		case "record":
			if !recording && (m.Author.Username == "Endmonaut" || m.Author.Username == "Liikt") {
				go f.Record(s, m, line, commChan)
				recording = true
			} else {
				s.ChannelMessageSend(m.ChannelID, "I CURRENTLY AM RECORDING YOU FUCK")
			}
		case "pause":
			commChan <- "pause"
		case "resume":
			commChan <- "resume"
		case "stop":
			commChan <- "stop"
			recording = false
		case "debug":
			f.Debug(s, m)
		default:
			go s.ChannelMessageSend(m.ChannelID, "I AM LIVING CANCER")
		}
	}
}
