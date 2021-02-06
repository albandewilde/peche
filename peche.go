package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	dgo "github.com/bwmarrin/discordgo"
)

const (
	SRCDIR = "/img_src/"
	DSTDIR = "/img_dst/"
)

var (
	TKN     string // Discord bot token
	MINTIME int64  // Minimum time between two sending (in seconde)
	MAXTIME int64  // Maximum time between two sending (in seconde)
)

func init() {
	TKN = os.Getenv("tkn")
	if TKN == "" {
		log.Fatal("No discord token found in environment variable `tkn`.")
	}

	var err error
	MINTIME, err = strconv.ParseInt(os.Getenv("min_time"), 10, 64)
	MAXTIME, err = strconv.ParseInt(os.Getenv("max_time"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	if MINTIME > MAXTIME {
		log.Fatal("`min_time` must be lower or equal to `max_time`")
	}
}

func main() {
	// Creating the bot instance
	b, err := dgo.New("Bot " + TKN)
	if err != nil {
		log.Fatal(err)
	}

	// Start the sending
	go sending(b, SRCDIR, DSTDIR, MINTIME, MAXTIME)

	// Add no handelers

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
