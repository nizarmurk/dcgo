package auth

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/nizarmurk/dcgo/env"
)

func DiscordAuth() {
	session, err := discordgo.New(env.TOKEN)
	if err != nil {
		fmt.Println(err)
		return
	}
	session.Open()

	// Signal-Handler hinzufügen
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Auf Signale warten
	<-sigChan

}
