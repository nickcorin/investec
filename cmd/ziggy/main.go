package main

import (
	"log"

	"github.com/nickcorin/ziggy/client"
	"github.com/nickcorin/ziggy/pkg/credentials"
)

func main() {
	creds, err := credentials.Get()
	if err != nil {
		log.Fatal(err)
	}

	_ = client.NewHTTP(creds.Username, creds.Secret)
}
