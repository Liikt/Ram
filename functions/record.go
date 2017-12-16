package functions

import (
	"encoding/binary"
	"fmt"
	"os"
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

func Record(s *discordgo.Session, m *discordgo.MessageCreate, close chan bool) {
	fileName := "Testfile.pcm"

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

	if channelToJoin == nil {
		log.Warning("Couldn't find the channel to join")
		s.ChannelMessageSend(m.ChannelID, "Couldn't find the channel to join")
	} else {
		voice, err := s.ChannelVoiceJoin(channelToJoin.GuildID, channelToJoin.ID, true, false)
		utils.CheckError(err, "Couldn't join the voice channel")
		defer voice.Disconnect()

		f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		mutex := &sync.Mutex{}
		packetArr := [][]int16{}

		recv := make(chan *discordgo.Packet, 2)
		go dgvoice.ReceivePCM(voice, recv)

		for {
			select {
			case packet, _ := <-recv:
				fmt.Println("SSRC:", packet.SSRC, "Sequence:", packet.Sequence, "Timestamp", packet.Timestamp, "Type:", packet.Type, "Opus:", packet.Opus, "PCM:", packet.PCM)
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
			case <-close:
				fmt.Println("Closing")
				return
			}
		}
	}
}
