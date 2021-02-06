package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	dgo "github.com/bwmarrin/discordgo"
)

func sending(b *dgo.Session, src, dst string, min, max int64) {
	for {
		// Change the waiting time
		waitSec := int64(rand.Intn(int(max-min))) + min
		waitingTime, err := time.ParseDuration(fmt.Sprintf("%ds", waitSec))
		if err != nil {
			log.Fatal(err)
		}

		select {
		case <-time.After(waitingTime):
			sendAndMoveFile(
				src,
				dst,
				b,
				"719171339788877857",
			)
		}
	}
}
