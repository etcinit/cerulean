package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/etcinit/cerulean/app"
	"github.com/etcinit/tumble/client"
	"github.com/facebookgo/inject"
	"github.com/jacobstr/confer"
)

func main() {
	// Create the configuration
	// In this case, we will be using the environment and some safe defaults
	config := confer.NewConfig()
	config.SetDefault("database.driver", "sqlite")
	config.SetDefault("database.file", ":memory:")
	config.ReadPaths("config/main.yml", "config/main.production.yml")
	config.AutomaticEnv()

	// For Chromabits, we want to include the Tumblr blog, so we will setup the
	// client here so we can sue it later.
	tumblr := client.NewWithKey(config.GetString("tumblr.key"))

	// Next, we setup the dependency graph
	// In this example, the graph won't have many nodes, but on more complex
	// applications it becomes more useful.
	var g inject.Graph
	var cerulean app.Cerulean
	g.Provide(
		&inject.Object{Value: config},
		&inject.Object{Value: tumblr},
		&inject.Object{Value: &cerulean},
	)
	g.Populate()

	// Setup the command line application
	app := cli.NewApp()
	app.Name = "cerulean"
	app.Usage = "Chromabits Blog Engine"

	// Set version and authorship info
	app.Version = "0.0.2"
	app.Author = "Eduardo Trujillo <ed@chromabits.com>"

	// Setup the default action. This action will be triggered when no
	// subcommand is provided as an argument
	app.Action = func(c *cli.Context) {
		fmt.Println("Usage: cerulean [global options] command [command options] [arguments...]")
	}

	app.Commands = []cli.Command{
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "migrates the database",
			Action:  cerulean.Migrator.Run,
		},
		{
			Name:    "serve",
			Aliases: []string{"s", "server", "listen"},
			Usage:   "starts the API server",
			Action:  cerulean.Serve.Run,
		},
	}

	// Begin
	app.Run(os.Args)
}
