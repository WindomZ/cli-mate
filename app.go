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

func (a *App) Run(arguments []string) error {
	return a.App.Run(arguments)
}

func (a *App) RunOSArgs() error {
	return a.App.Run(os.Args)
}

func (a *App) AddFlag(f Flag) {
	a.App.Flags = append(a.App.Flags, f.Flag)
	a.Command.AddFlag(f)
}

func (a *App) AddFlags(fs []Flag) {
	a.Command.AddFlags(fs)
}

func (a *App) AddCommand(c Command) {
	a.App.Commands = append(a.App.Commands, c.GetCliSubCommand()...)
}

func (a *App) AddCommands(cs []Command) {
	if cs != nil && len(cs) != 0 {
		for _, c := range cs {
			a.App.Commands = append(a.App.Commands, c.GetCliSubCommand()...)
		}
	}
}
