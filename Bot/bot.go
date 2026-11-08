package bot

import (
	"fmt"
	config "qme/Config"
	database "qme/Database"
	handlers "qme/Handlers"
	logger "qme/Logger"

	"github.com/bwmarrin/discordgo"
)

func Init() {
	if !database.Init() {
		logger.ErrorText("BOTINIT", "Failed to initialise database")
		panic("Failed to initialise database")
	}

	if !sessionInit() {
		logger.ErrorText("BOTINIT", "Failed to initialise discord session")
		panic("Failed to initialise discord session")
	}

	if !addHandlers() {
		logger.ErrorText("BOTINIT", "Failed to add discord event handlers")
		panic("Failed to add discord event handlers")
	}

	if !sessionOpen() {
		logger.ErrorText("BOTINIT", "Failed to open discord session")
		panic("Failed to open discord session")
	}
}

func sessionInit() bool {
	session, err := discordgo.New("Bot " + config.Settings.DISCORDBOTTOKEN)
	if err != nil {
		logger.Error("BOTINIT", err)
		return false
	} else if session == nil {
		logger.Error("BOTINIT", fmt.Errorf("discord session was nil"))
		return false
	}

	config.Session = session
	return true
}

func sessionOpen() bool {
	err := config.Session.Open()
	if err != nil {
		logger.Error("BOTINIT", err)
		return false
	}

	if config.Session == nil {
		logger.Error("BOTINIT", err)
		return false
	}

	config.Settings.DISCORDSESSIONID = config.Session.Token
	return true
}

func addHandlers() bool {
	config.Session.AddHandler(handlers.NewMessage) //   New Messages
	config.Session.AddHandler(handlers.NewGuild)   //	Server connected to the bot
	return true
}
