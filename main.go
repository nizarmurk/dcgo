package main

import (
	"github.com/nizarmurk/dcgo/auth"
	"github.com/nizarmurk/dcgo/core"
	"github.com/nizarmurk/dcgo/env"
)

func main() {
	core.Core()
	auth.DiscordAuth()
	env.Env()
}
