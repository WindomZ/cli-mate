package cli_mate

import "github.com/urfave/cli"

type CommandAction func(c *Context) ExitCoder

type Command struct {
	Command     cli.Command
	Action      CommandAction
	action      CommandAction
	flags       []Flag
	subCommands []Command
}

func (cmd *Command) run() {
	// define default Command.Action if it`s nil
	if cmd.Action == nil {
		cmd.Action = func(c *Context) ExitCoder {
			return nil
		}
	}
	// define default Command.action
	cmd.action = func(c *Context) ExitCoder {
		for _, f := range cmd.flags {
			if err := f.Action(c, f); err != nil {
				if err.IsBreak() {
					break
				}
				return err
			}
		}
		if cmd.Action == nil {
		} else if err := cmd.Action(c); err != nil {
			return err
		}
		return nil
	}
	// Command.action adapter cli.Command.Action
	cmd.Command.Action = func(c *cli.Context) error {
		if err := cmd.action(NewContext(c)); err != nil {
			return err.GetError()
		}
		return nil
	}
}

func (c *Command) AddFlag(f Flag) {
	c.flags = append(c.flags, f)
	c.Command.Flags = append(c.Command.Flags, f.Flag)
}

func (c *Command) AddFlags(fs []Flag) {
	for _, f := range fs {
		c.AddFlag(f)
	}
}

// AddSubCommand add a child command to list of child commands
func (c *Command) AddSubCommand(cmd Command) {
	c.subCommands = append(c.subCommands, cmd)
	c.Command.Subcommands = c.Command.Subcommands[:0]
	c.Command.Subcommands = append(c.Command.Subcommands, c.getCliSubCommand()...)
}

// getCliSubCommand get list of child cli.commands
func (c *Command) getCliSubCommand() []cli.Command {
	if c.subCommands == nil || len(c.subCommands) == 0 {
		return []cli.Command{}
	}
	cs := make([]cli.Command, 0, len(c.subCommands))
	for _, sub := range c.subCommands {
		cs = append(cs, sub.Command)
	}
	return cs
}
