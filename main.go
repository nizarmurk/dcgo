package main

import (
	"github.com/nizarmurk/dcgo/auth"
	"github.com/nizarmurk/dcgo/core"
	"github.com/nizarmurk/dcgo/env"
)

func main() {
	env.Env()
	core.Core()
	auth.DiscordAuth()
}
