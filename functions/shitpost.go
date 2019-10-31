package functions

import (
	"math/rand"
	"strings"

	utils "../utils"
	"github.com/bwmarrin/discordgo"
)

type embeddedShitpost struct {
	colour      int
	link        string
	description string
}

type reactionalShitpost struct {
	reaction string
}

var shitMap = make(map[string]interface{})

func handleEmbed(embed embeddedShitpost, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: embed.description,
		Color:       embed.colour,
		Image:       &discordgo.MessageEmbedImage{URL: embed.link},
	})
}

func handleNormal(embed reactionalShitpost, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, embed.reaction)
}

// Shitpost handles the shitposting of the bot. Whenever a received message contains a key of the
// `shitMap` the bot will send the corresponding shitpost
func Shitpost(s *discordgo.Session, m *discordgo.MessageCreate) {
	shitMap = map[string]interface{}{
		"woah":                embeddedShitpost{0xEB8D43, "http://i0.kym-cdn.com/photos/images/newsfeed/001/279/920/054.png", "WOAH"},
		"desire":              embeddedShitpost{0xEB8D43, getDesireLink(), "OH DESIRE"},
		"friendship is magic": embeddedShitpost{0x8958A7, "https://cdn.discordapp.com/attachments/285537911414325249/291701362218237962/MagicIsHeresy.jpg", "AND MAGIC IS HERESY!"},
		"reee":                embeddedShitpost{0xFF0000, "http://i1.kym-cdn.com/entries/icons/original/000/017/318/angry_pepe.jpg", "REEEEEEEEE"},
		"do it":               embeddedShitpost{0x000000, "https://cdn.discordapp.com/attachments/302709996041273344/302713643106304000/dew_it.jpg", ""},
		"dew it":              embeddedShitpost{0x000000, "https://cdn.discordapp.com/attachments/302709996041273344/302713643106304000/dew_it.jpg", ""},
		"nuu":                 reactionalShitpost{":flag_de:"},
		"gib":                 reactionalShitpost{":flag_gr:"},
		"balans":              reactionalShitpost{":flag_ru:"},
		"murica":              reactionalShitpost{":flag_us:"},
		"fuck yeah":           reactionalShitpost{":eagle:"},
		"your mom":            reactionalShitpost{"NO YOUR MOM!"},
		"dumb":                reactionalShitpost{"NO YOU'RE DUMB!"},
	}

	for name, f := range shitMap {
		switch f := f.(type) {
		case embeddedShitpost:
			go func(key string, embed embeddedShitpost) {
				if strings.Contains(strings.ToLower(m.Content), key) {
					handleEmbed(embed, s, m)
				}
			}(name, f)
		case reactionalShitpost:
			go func(key string, embed reactionalShitpost) {
				if strings.Contains(strings.ToLower(m.Content), key) {
					handleNormal(embed, s, m)
				}
			}(name, f)
		}

	}

}

func getDesireLink() string {
	desireLinks := utils.LoadDesire()
	return (*desireLinks)[rand.Intn(len(*desireLinks))]
}
