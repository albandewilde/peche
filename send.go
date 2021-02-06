package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"

	dgo "github.com/bwmarrin/discordgo"
)

func sendAndMoveFile(src, dst string, s *dgo.Session, channelID string) error {
	// Read files in source directory
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	// Check if there is file to send
	if len(files) <= 0 {
		return nil
	}

	// Randomly choose a file in the source directory
	file, err := chooseFile(files)
	if err != nil {
		return err
	}

	// Open the file for reading
	f, err := os.Open(src + file.Name())
	defer f.Close()
	if err != nil {
		return err
	}
	// Send that file
	_, err = s.ChannelFileSend(channelID, f.Name(), f)
	if err != nil {
		return err
	}

	// Move the file to the destination directory
	move(src, dst, file.Name())

	return nil
}

func move(src, dst, fle string) {
	// Use shell command because os.Rename return `invalid cross-device link`
	cmd := exec.Command("mv", src+fle, dst+fle)
	cmd.Run()
}

func chooseFile(imgs []os.FileInfo) (os.FileInfo, error) {
	// Get a random index
	nbImgs := len(imgs)
	idx := rand.Intn(nbImgs)

	return imgs[idx], nil
}
