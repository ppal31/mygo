package server

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/ppal31/mygo/internal/config"
	"github.com/ppal31/mygo/internal/logger"
	"github.com/ppal31/mygo/internal/router"
	"github.com/ppal31/mygo/internal/server"
	"github.com/ppal31/mygo/internal/store"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

type command struct {
	envfile string
}

func (c *command) run(*kingpin.ParseContext) error {
	_ = godotenv.Load(c.envfile)
	config, err := load()
	if err != nil {
		logger.Error(context.Background(), "error loading config", "err", err)
		os.Exit(1)
	}
	// configure the log level
	setupLogger(config)
	s, err := initServer(config)
	if err != nil {
		logger.Error(context.Background(), "error in init server", "err", err)
		os.Exit(1)
	}
	return s.ListenAndServe(context.Background())
}

func initServer(config *config.AppConfig) (*server.Server, error) {
	db, err := store.Connect(config.Database.Driver, config.Database.Datasource, config.Seed)
	if err != nil {
		return nil, err
	}
	handler := router.New(config, db)
	serverServer := &server.Server{
		Addr:    config.Server.Bind,
		Host:    config.Server.Host,
		Handler: handler,
	}

	return serverServer, nil
}

func setupLogger(config *config.AppConfig) {
	logger.L = logger.NewLogger(config.Debug)
}

// Register the server command.
func Register(app *kingpin.Application) {
	c := new(command)
	cmd := app.Command("server", "starts the server").
		Action(c.run)

	cmd.Arg("envfile", "load the environment variable file").
		Default("").
		StringVar(&c.envfile)
}
