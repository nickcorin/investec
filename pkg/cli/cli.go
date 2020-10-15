package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/nickcorin/ziggy"
	"github.com/nickcorin/ziggy/client"
	"github.com/nickcorin/ziggy/pkg/config"
	"github.com/nickcorin/ziggy/pkg/credentials"

	docker_credentials "github.com/docker/docker-credential-helpers/credentials"
)

// App contains the state of a CLI application.
type App struct {
	creds  *docker_credentials.Credentials
	conf   *config.Config
	router *CommandRouter
	ziggy  ziggy.Client
}

// New returns a new App instance.
func New() (*App, error) {
	var app App

	err := config.WriteDefault(config.DefaultConfigFile, false)
	if err != nil {
		return nil, fmt.Errorf("failed to write default config: %w", err)
	}

	conf, err := config.Load(config.DefaultConfigFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	creds, err := credentials.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get credentials: %w", err)
	}

	app.conf = conf
	app.creds = creds
	app.router = NewRouter()
	app.ziggy = client.NewHTTP(creds.Username, creds.Secret)

	RegisterRoutes(app.router)

	return &app, nil
}

func (app *App) Run(ctx context.Context, cmd string) {
	c, args, err := app.router.Search(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Run(ctx, app.ziggy, os.Stdout, args...)
}
