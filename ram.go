package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	cb "../Ram/callbacks"
	utils "../Ram/utils"
	"github.com/bwmarrin/discordgo"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if utils.Key == "" {
		fmt.Println("No key provided. Please add 'var Key string = <key>' to utils/secret.go")
		return
	}

	dg, err := discordgo.New("Bot " + utils.Key)
	utils.CheckError(err, "Error creating Discord session: ", err)

	dg.AddHandler(cb.Ready)
	dg.AddHandler(cb.OnMessage)

	err = dg.Open()
	utils.CheckError(err, "Error opening Discord session: ", err)

	fmt.Println("Ram is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
