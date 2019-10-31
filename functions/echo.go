package functions

import "github.com/bwmarrin/discordgo"

// Echo is a simple echo function. The Bot will just repeat what it got send.
func Echo(s *discordgo.Session, m *discordgo.MessageCreate, line string) {
	if line != "" {
		go s.ChannelMessageSend(m.ChannelID, line)
	} else {
		go s.ChannelMessageSend(m.ChannelID, "YOU SAID NOTHING")
	}
}
