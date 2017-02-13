package cli_mate

import "github.com/urfave/cli"

// CommandAction the function for Command action
type CommandAction func(c *Context) ExitCoder

// Command is a subcommand for a cli.App
type Command struct {
	Command     cli.Command
	Action      CommandAction
	action      CommandAction
	flags       []*Flag
	subCommands []*Command
}

// register get Command After the registration
func (cmd *Command) register() *Command {
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

	// register cli.Command.Flags
	cmd.Command.Flags = make([]cli.Flag, 0, len(cmd.flags))
	for _, f := range cmd.flags {
		cmd.Command.Flags = append(cmd.Command.Flags, f.register().Flag)
	}

	// register cli.Command.Subcommands
	cmd.Command.Subcommands = make(cli.Commands, 0, len(cmd.subCommands))
	for _, c := range cmd.subCommands {
		cmd.Command.Subcommands = append(cmd.Command.Subcommands, c.register().Command)
	}

	return cmd
}

// AddFlag add a Flag to list of Flags
func (c *Command) AddFlag(f Flag) {
	c.flags = append(c.flags, &f)
}

// AddFlags add an array of Flags to list of Flags
func (c *Command) AddFlags(fs []Flag) {
	for _, f := range fs {
		c.AddFlag(f)
	}
}

// AddSubCommand add a child Command to list of child Commands
func (c *Command) AddSubCommand(cmd Command) {
	c.subCommands = append(c.subCommands, &cmd)
}

// AddSubCommands add an array of child Commands to list of child Commands
func (c *Command) AddSubCommands(cmds []Command) {
	for _, cmd := range cmds {
		c.AddSubCommand(cmd)
	}
}
