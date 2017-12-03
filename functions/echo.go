package functions

import "github.com/bwmarrin/discordgo"

func Echo(s *discordgo.Session, m *discordgo.MessageCreate, line string) {
	if line != "" {
		go s.ChannelMessageSend(m.ChannelID, line)
	} else {
		go s.ChannelMessageSend(m.ChannelID, "YOU SAID NOTHING")
	}
}
