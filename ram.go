package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	utils "../Ram/utils"
	"github.com/bwmarrin/discordgo"
)

func main() {

	if utils.Key == "" {
		fmt.Println("No token provided. Please run: airhorn -t <bot token>")
		return
	}

	dg, err := discordgo.New("Bot " + utils.Key)
	utils.CheckError(err, "Error creating Discord session: ", err)

	dg.AddHandler(ready)

	err = dg.Open()
	utils.CheckError(err, "Error opening Discord session: ", err)

	fmt.Println("Ram is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "CANCER")
}
