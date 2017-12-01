package callbacks

import (
	"fmt"
	"strings"

	utils "../utils"
	"github.com/bwmarrin/discordgo"
)

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	if utils.MentionsContainsUser(m.Mentions, s.State.User) {
		c, err := s.State.Channel(m.ChannelID)
		utils.CheckError(err, "Couldn't get the Channel with the ID: ", m.ChannelID)

		_, err = s.State.Guild(c.GuildID)
		utils.CheckError(err, "Couldn't get the Server with the ID: ", c.GuildID)

		split := strings.SplitN(m.ContentWithMentionsReplaced(), " ", 3)
		command, line := split[1], ""

		if len(split) == 3 {
			line = split[2]
		}

		fmt.Println("Got command", command, "with line", line)
	}
}
