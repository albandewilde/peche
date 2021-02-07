package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	dgo "github.com/bwmarrin/discordgo"
)

const (
	SRCDIR = "/img_src/"
	DSTDIR = "/img_dst/"
)

var (
	TKN         string   // Discord bot token
	MINTIME     int64    // Minimum time between two sending (in seconde)
	MAXTIME     int64    // Maximum time between two sending (in seconde)
	channelsIDs []string // Channel id to send pictures
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

	// Add no handers

	// Open a websocket connection to Discord and begin listening.
	err = b.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connecting...")

	// Wait for connecting
	<-time.After(time.Second * 5)

	fmt.Println("I'm logged in ! (Press CTRL-C to exit.)")

	// Get channels list
	channelsNames := readChannelsName()
	channelsIDs := channelsNameToID(b, channelsNames)

	// Start the sending
	go sending(b, SRCDIR, DSTDIR, MINTIME, MAXTIME, channelsIDs)

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	b.Close()
}

func readChannelsName() []string {
	return strings.Split(os.Getenv("channels"), ",")

}

func channelsNameToID(b *dgo.Session, chs []string) []string {
	guilds := b.State.Guilds
	ids := make([]string, 0)

	// Find channel id with their name
	for _, g := range guilds { // Loop on each guild
		for _, c := range g.Channels { // Loop on each channel of the guild
			for _, channelName := range chs { // Loop on each channel name we have
				if c.Name == channelName && !include(ids, c.ID) {
					ids = append(ids, c.ID)
				}
			}
		}
	}

	return ids
}
