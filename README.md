# Discord-Bot

This is a simple discord bot written in Golang.

### Features

It displays a welcome message to users who have joined the discord server. \
It also has a ping command to retrieve the latency of the bot.

## Usage

| Environment variable | Description                                                                                 | Required | Default |
| --- |---------------------------------------------------------------------------------------------| --- |---------|
| DISCORD_BOT_TOKEN | Discord Bot Token                                                                           | yes | `""`    |
| COMMAND_PREFIX | Prefix for the bot. Must be exactly 1 character, or it will default to `!`                  | no | `!`     |
| WELCOME_CHANNEL | Channel that greets new users to the discord                                                | no | `NONE`  |
| VERIFY_CHANNEL | Channel users will be redirected to, so they can verify. (Extension to the welcome command) | no | `NONE`  |


## Getting started

### Discord

1. Go to https://discord.com/developers/applications
2. Create a new application
3. Create a bot in your application
4. Retreive the bot's token and set it as the `DISCORD_BOT_TOKEN` environment variable
5. Go to `https://discordapp.com/oauth2/authorize?client_id=<YOUR_BOT_CLIENT_ID>&scope=bot&permissions=36785216`
6. Add the bot to your server

### Bot commands

Assuming `COMMAND_PREFIX` is not defined or is set to `!`.

#### Help

```
!help
OUTPUT:
Displays all commands
```

#### Ping

```
!ping
OUTPUT:
Retreives the bot's latency
```

## Prerequisites

If you want to run the bot locally, you'll need the following: \
`go (version 1.17.5 is stable)`

### Install go

#### Linux
`Debian/Ubuntu` \
sudo apt-get update \
sudo apt-get install golang-go. 

`Arch` \
sudo pacman -Syu \
sudo pacman -S go

To check the version type: \
go version

You can also manually install Go here: https://go.dev/doc/install

#### Windows

Follow this guide: https://go.dev/doc/install

## Run the Bot

Go to the project directory and run:

```
go build
```

When you have specified the desired environment variables, run the following command:

```
go run .
```

Examples:
```
DISCORD_BOT_TOKEN=token COMMAND_PREFIX=! WELCOME_CHANNEL=ID VERIFY_CHANNEL=ID go run .
DISCORD_BOT_TOKEN=token go run .
```