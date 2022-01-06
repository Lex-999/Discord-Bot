package config

import (
	"os"
	"strings"
)

type Config struct {
	DiscordToken string
	//OwnerID        string
	CommandPrefix  string
	WelcomeChannel string
	VerifyChannel  string
}

var cfg *Config

// Load Initializes the configuration
func Load() {
	cfg = &Config{
		DiscordToken:  strings.TrimSpace(os.Getenv("DISCORD_BOT_TOKEN")),
		CommandPrefix: strings.TrimSpace(os.Getenv("COMMAND_PREFIX")),
		//OwnerID:        strings.TrimSpace(os.Getenv("OWNER_ID")),
		WelcomeChannel: strings.TrimSpace(os.Getenv("WELCOME_CHANNEL")),
		VerifyChannel:  strings.TrimSpace(os.Getenv("VERIFY_CHANNEL")),
	}

	if len(cfg.DiscordToken) == 0 {
		panic("The environment variable 'DISCORD_BOT_TOKEN' must not be empty!")
	}
	if len(cfg.CommandPrefix) != 1 {
		cfg.CommandPrefix = "!"
	}
	//if len(cfg.OwnerID) == 0 {
	//	panic("The environment variable 'OWNER_ID' must not be empty!")
	//}
	if len(cfg.WelcomeChannel) == 0 {
		cfg.WelcomeChannel = "NONE"
	}
	if len(cfg.VerifyChannel) == 0 {
		cfg.VerifyChannel = "NONE"
	}
}

// Get returns the configuration
func Get() *Config {
	return cfg
}
