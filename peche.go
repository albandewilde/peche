package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	dgo "github.com/bwmarrin/discordgo"
)

const (
	SRCDIR = "/img_src/"
	DSTDIR = "/img_dst/"
)

var TKN string // Discord bot token

func init() {
	TKN = os.Getenv("tkn")
	if TKN == "" {
		log.Fatal("No discord token found in environment variable `tkn`.")
	}
}

func main() {
	// Creating the bot instance
	b, err := dgo.New("Bot " + TKN)
	if err != nil {
		log.Fatal(err)
	}

	// Open a websocket connection to Discord and begin listening.
	err = b.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("I'm logged in ! (Press CTRL-C to exit.)")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	b.Close()
}
