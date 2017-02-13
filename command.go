package cli_mate

import "github.com/urfave/cli"

type Command struct {
	Command     cli.Command
	Action      func(c *Context) ExitCoder
	flags       []Flag
	SubCommands []Command
}

func NewCommand(cmd *cli.Command) (_cmd *Command) {
	if cmd != nil {
		_cmd = &Command{
			Command: *cmd,
		}
	} else {
		_cmd = &Command{}
	}
	// define Command.Action
	_cmd.Action = func(c *Context) ExitCoder {
		if _cmd.flags == nil {
			return nil
		}
		for _, f := range _cmd.flags {
			if err := f.Action(c, f); err != nil {
				return err
			}
		}
		return nil
	}
	// Command.Action adapter cli.Command.Action
	_cmd.Command.Action = func(c *cli.Context) error {
		if err := _cmd.Action(NewContext(c)); err != nil {
			return err.GetError()
		}
		return nil
	}
	return
}

func (c *Command) AddFlag(f Flag) {
	c.flags = append(c.flags, f)
}

func (c *Command) AddFlags(fs []Flag) {
	for _, f := range fs {
		c.AddFlag(f)
	}
}

// AddSubCommand add a child command to list of child commands
func (c *Command) AddSubCommand(cmd Command) {
	c.SubCommands = append(c.SubCommands, cmd)
}

// getSubCommand get list of child cli.commands
func (c *Command) GetCliSubCommand() []cli.Command {
	if c.SubCommands == nil || len(c.SubCommands) == 0 {
		return []cli.Command{}
	}
	cs := make([]cli.Command, 0, len(c.SubCommands))
	for _, sub := range c.SubCommands {
		cs = append(cs, sub.Command)
	}
	return cs
}
