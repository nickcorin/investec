package main

import (
	"log"
	"os"

	"github.com/nickcorin/ziggy"
	"github.com/nickcorin/ziggy/ziggyd"
	"github.com/nickcorin/ziggy/ziggyd/credentials"
)

func main() {
	creds, err := credentials.Get()
	if err != nil {
		log.Fatal(err)
	}

	client := ziggy.NewClient(creds.Username, creds.Secret)

	err = ziggyd.Run(os.Args[1], client, os.Args[2:]...)
	if err != nil {
		log.Fatal(err)
	}
}
