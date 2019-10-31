package callbacks

import (
	utils "../utils"
	"github.com/bwmarrin/discordgo"
)

// Ready will be executed once the Bot is initialized. Currently this is only used for shitposting
func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "CANCER")
	channels, err := s.GuildChannels("258726617361285131")
	utils.CheckError(err, "Couldn't get the Server with the ID: ", "283066637928890379")

	for _, channel := range channels {
		go func(c *discordgo.Channel) {
			if c.Name == "dumb-bot-shit" {
				_, _ = s.ChannelMessageSend(c.ID, "I HAVE RETURNED WITH MORE CANCER")
			}
		}(channel)
	}
}
