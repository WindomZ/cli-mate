package cli_mate

import (
	"github.com/urfave/cli"
	"os"
)

type App struct {
	App     *cli.App
	Command *Command
}

func NewApp() *App {
	_app := &App{
		App:     cli.NewApp(),
		Command: NewCommand(nil),
	}
	_app.App.Action = func(c *cli.Context) error {
		if err := _app.Command.Action(NewContext(c)); err != nil {
			return err.GetError()
		}
		return nil
	}
	return _app
}

func (a *App) AddFlag(f *Flag) {
	if f != nil {
		a.App.Flags = append(a.App.Flags, f.Flag)
		a.Command.AddFlag(f)
	}
}

func (a *App) AddFlags(fs []*Flag) {
	a.Command.AddFlags(fs)
}

func (a *App) Run(arguments []string) error {
	return a.App.Run(arguments)
}

func (a *App) RunOSArgs() error {
	return a.App.Run(os.Args)
}
