package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/nickcorin/ziggy/pkg/cli"
)

func main() {
	app, err := cli.New()
	if err != nil {
		log.Fatal(err)
	}


	app.Run(context.Background(), strings.Join(os.Args[1:], ""))
}
