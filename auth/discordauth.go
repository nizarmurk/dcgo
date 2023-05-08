package auth

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nizarmurk/dcgo/env"
)

func DiscordAuth() {
	_, err := discordgo.New(env.BOTNAME + env.TOKEN)
	if err != nil {
		fmt.Println(err)
		return
	}
}
