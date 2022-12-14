package main

import (
	log "github.com/sirupsen/logrus"

	"colorize/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Error(err)
	}
}
