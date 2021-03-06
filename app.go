package cli_mate

import (
	"github.com/urfave/cli"
	"os"
)

// App is the main structure of a cli application
type App struct {
	App      *cli.App
	Action   CommandAction
	action   CommandAction
	flags    []*Flag
	commands []*Command
}

// NewApp creates a new cli Application with Name, Usage and Version
func NewApp(name, usage, version string) *App {
	a := cli.NewApp()
	a.Name = name
	a.Usage = usage
	a.Version = version
	return &App{
		App: a,
	}
}

// register get App After the registration
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
	app.App.Flags = make([]cli.Flag, 0, len(app.flags))
	for _, f := range app.flags {
		app.App.Flags = append(app.App.Flags, f.register().Flag)
	}

	// register cli.App.Commands
	app.App.Commands = make([]cli.Command, 0, len(app.commands))
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

// Run is the entry point to the cli app
func (a *App) Run(arguments []string) error {
	return a.register().App.Run(arguments)
}

// RunOSArgs same as Run with arguments from os
func (a *App) RunOSArgs() error {
	return a.Run(os.Args)
}

// AddFlag add a Flag to list of Flags
func (a *App) AddFlag(f Flag) *App {
	a.flags = append(a.flags, &f)
	return a
}

// AddFlags add an array of Flags to list of Flags
func (a *App) AddFlags(fs []Flag) *App {
	for _, f := range fs {
		a.AddFlag(f)
	}
	return a
}

// AddCommand add a child Command to list of child Commands
func (a *App) AddCommand(c Command) *App {
	a.commands = append(a.commands, &c)
	return a
}

// AddCommands add an array of child Commands to list of child Commands
func (a *App) AddCommands(cs []Command) *App {
	for _, c := range cs {
		a.AddCommand(c)
	}
	return a
}

// Clear clear Commands and Flags to the initial state
func (a *App) Clear() {
	a.flags = a.flags[:0]
	a.commands = a.commands[:0]
}
