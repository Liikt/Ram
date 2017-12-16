package functions

import (
	"encoding/binary"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	utils "../utils"
	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const (
	sampleRate = 48000
	channels   = 1
)

func getCurrentVoiceChannel(s *discordgo.Session, user *discordgo.User, guild *discordgo.Guild) *discordgo.Channel {
	for _, vs := range guild.VoiceStates {
		if vs.UserID == user.ID {
			channel, _ := s.State.Channel(vs.ChannelID)
			return channel
		}
	}
	return nil
}

func Record(s *discordgo.Session, m *discordgo.MessageCreate, filename string, closeChan chan bool) {
	channel, _ := s.State.Channel(m.ChannelID)
	if channel == nil {
		log.WithFields(log.Fields{
			"channel": m.ChannelID,
			"message": m.ID,
		}).Warning("Failed to grab channel")
		return
	}

	guild, _ := s.State.Guild(channel.GuildID)
	if guild == nil {
		log.WithFields(log.Fields{
			"guild":   channel.GuildID,
			"channel": channel,
			"message": m.ID,
		}).Warning("Failed to grab guild")
		return
	}

	channelToJoin := getCurrentVoiceChannel(s, m.Author, guild)
	fmt.Println(channelToJoin.Name, channelToJoin.Recipients)

	if match, _ := !regexp.Match("^[A-Za-z0-9._]+$", filename); filename == "" || !match {
		filename = time.Now().Format("2006-02-Jan")
	}
	if !strings.HasSuffix(filename, ".pcm") {
		filename += ".pcm"
	}

	filename = "recordings/" + filename

	if channelToJoin == nil {
		log.Warning("Couldn't find the channel to join")
		s.ChannelMessageSend(m.ChannelID, "Couldn't find the channel to join")
	} else {
		mutex := &sync.Mutex{}
		packetArr := [][]int16{}
		voice, err := s.ChannelVoiceJoin(channelToJoin.GuildID, channelToJoin.ID, true, false)
		utils.CheckError(err, "Couldn't join the voice channel")
		defer voice.Disconnect()

		f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		recv := make(chan *discordgo.Packet, 2)
		go dgvoice.ReceivePCM(voice, recv)

		for {
			select {
			case <-closeChan:
				fmt.Println("Closing")
				mutex.Lock()
				for _, packet := range packetArr {
					binary.Write(f, binary.LittleEndian, packet)
				}
				mutex.Unlock()
				return

			case packet, _ := <-recv:
				packetArr = append(packetArr, packet.PCM)

			case <-time.After(1 * time.Second):
				if len(packetArr) > 0 {
					fmt.Println("Writing to file")
					go func() {
						mutex.Lock()
						tmp := packetArr
						packetArr = [][]int16{}
						for _, packet := range tmp {
							binary.Write(f, binary.LittleEndian, packet)
						}
						mutex.Unlock()
					}()
				}
			}
		}
	}
}
