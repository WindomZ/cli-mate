package cli_mate

import "github.com/urfave/cli"

type App struct {
	App     *cli.App
	command Command
}

func NewApp() *App {
	return &App{
		App: cli.NewApp(),
	}
}
