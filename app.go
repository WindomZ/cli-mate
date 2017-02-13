package cli_mate

import (
	"github.com/urfave/cli"
	"os"
)

type App struct {
	App      *cli.App
	Action   CommandAction
	action   CommandAction
	flags    []*Flag
	commands []*Command
}

func NewApp(name, version string) *App {
	a := cli.NewApp()
	a.Name = name
	a.Version = version
	return &App{
		App: a,
	}
}

func (app *App) register() *App {
	// define default App.Action if it`s nil
	if app.Action == nil {
		app.Action = func(c *Context) ExitCoder {
			return nil
		}
	}

	// define default App.action
	app.action = func(c *Context) ExitCoder {
		for _, f := range app.flags {
			if err := f.Action(c, f); err != nil {
				if err.IsBreak() {
					break
				}
				return err
			}
		}
		if app.Action == nil {
		} else if err := app.Action(c); err != nil {
			return err
		}
		return nil
	}

	// App.action adapter cli.App.Action
	app.App.Action = func(c *cli.Context) error {
		if err := app.action(NewContext(c)); err != nil {
			return err.GetError()
		}
		return nil
	}

	// register cli.App.Flags
	app.App.Flags = make([]cli.Flag, len(app.flags))
	for _, f := range app.flags {
		app.App.Flags = append(app.App.Flags, f.register().Flag)
	}

	// register cli.App.Commands
	app.App.Commands = make([]cli.Command, len(app.commands))
	for _, c := range app.commands {
		app.App.Commands = append(app.App.Commands, c.register().Command)
	}

	return app
}

func (a *App) Name() string {
	return a.App.Name
}

func (a *App) Version() string {
	return a.App.Version
}

func (a *App) Run(arguments []string) error {
	return a.register().App.Run(arguments)
}

func (a *App) RunOSArgs() error {
	return a.Run(os.Args)
}

func (a *App) AddFlag(f Flag) {
	a.flags = append(a.flags, &f)
}

func (a *App) AddFlags(fs []Flag) {
	for _, f := range fs {
		a.AddFlag(f)
	}
}

func (a *App) AddCommand(c Command) {
	a.commands = append(a.commands, &c)
}

func (a *App) AddCommands(cs []Command) {
	for _, c := range cs {
		a.AddCommand(c)
	}
}
