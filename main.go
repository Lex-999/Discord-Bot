package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wrldguard-bot/config"
)

func main() {
	config.Load()

	// Create Discord Bot Session
	bot, err := discordgo.New("Bot " + config.Get().DiscordToken)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		println("err")
		panic(err)
	}

	// Safely close down the Discord session
	defer bot.Close()
	defer func() {
		log.Printf("Shutting Down")
		time.Sleep(1 * time.Second)
	}()

	// Register the handleMessage func
	bot.AddHandler(HandleMessage)
	bot.AddHandler(onGuildMemberAdd)

	// Identify necessary bot intents
	bot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers

	// Open a websocket connection to Discord.
	err = bot.Open()
	if err != nil {
		log.Println("Error Connecting!", err)
		return
	}

	_ = bot.UpdateListeningStatus(fmt.Sprint("Music!"))

	log.Println("Successfully Connected!")
	log.Println("Command Prefix is = " + config.Get().CommandPrefix)

	// Wait till the bot is killed
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-channel

	log.Println("Terminating Bot")
}

func HandleMessage(bot *discordgo.Session, message *discordgo.MessageCreate) {
	var commandPrefix = config.Get().CommandPrefix

	// Ignore messages from the bot itself
	if message.Author.Bot || message.Author.ID == bot.State.User.ID {
		return
	}

	latency := bot.HeartbeatLatency()
	if message.Content == commandPrefix+"ping" {
		_, err := bot.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Latency: %s", latency))
		if err != nil {
			log.Println("Failed to send latency message", err)
			return
		}
	}
}

func onGuildMemberAdd(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	guild, _ := session.Guild(event.GuildID)

	// Time delay so that welcome message is sent after member has joined
	time.Sleep(4 * time.Second)

	if config.Get().WelcomeChannel != "NONE" && config.Get().VerifyChannel != "NONE" {
		_, err := session.ChannelMessageSend(config.Get().WelcomeChannel, "Welcome to "+guild.Name+" "+event.Mention()+
			" To access the server, verify in <#"+config.Get().VerifyChannel+">")
		if err != nil {
			log.Println("Failed to send welcome message", err)
			return
		}
	} else if config.Get().WelcomeChannel != "NONE" && config.Get().VerifyChannel == "NONE" {
		_, err := session.ChannelMessageSend(config.Get().WelcomeChannel, "Welcome to "+guild.Name+" "+event.Mention())
		if err != nil {
			log.Println("Failed to send welcome message", err)
			return
		}
	} else {
		return
	}
}
