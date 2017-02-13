package cli_mate

import (
	"github.com/urfave/cli"
	"os"
)

type App struct {
	App      *cli.App
	Action   CommandAction
	action   CommandAction
	flags    []Flag
	commands []Command
}

func NewApp() *App {
	return &App{
		App: cli.NewApp(),
	}
}

func (a *App) run() {
	// define default App.Action if it`s nil
	if a.Action == nil {
		a.Action = func(c *Context) ExitCoder {
			return nil
		}
	}
	// define default App.action
	a.action = func(c *Context) ExitCoder {
		for _, f := range a.flags {
			if err := f.Action(c, f); err != nil {
				if err.IsBreak() {
					break
				}
				return err
			}
		}
		if a.Action == nil {
		} else if err := a.Action(c); err != nil {
			return err
		}
		return nil
	}
	// App.action adapter cli.App.Action
	a.App.Action = func(c *cli.Context) error {
		if err := a.action(NewContext(c)); err != nil {
			return err.GetError()
		}
		return nil
	}
}

func (a *App) Run(arguments []string) error {
	a.run()
	for _, c := range a.commands {
		c.run()
	}
	return a.App.Run(arguments)
}

func (a *App) RunOSArgs() error {
	return a.Run(os.Args)
}

func (a *App) AddFlag(f Flag) {
	a.flags = append(a.flags, f)
	a.App.Flags = append(a.App.Flags, f.Flag)
}

func (a *App) AddFlags(fs []Flag) {
	for _, f := range fs {
		a.AddFlag(f)
	}
}

func (a *App) AddCommand(c Command) {
	a.commands = append(a.commands, c)
	a.App.Commands = append(a.App.Commands, c.Command)
}

func (a *App) AddCommands(cs []Command) {
	for _, c := range cs {
		a.AddCommand(c)
	}
}
