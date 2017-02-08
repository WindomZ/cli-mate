package cli_mate

import "github.com/urfave/cli"

type Command struct {
	Command     *cli.Command
	Action      func(c *Context) ExitCoder
	flags       []*Flag
	SubCommands []*Command
}

func NewCommand(cmd *cli.Command) *Command {
	_cmd := &Command{
		Command: cmd,
	}
	_cmd.Action = func(c *Context) ExitCoder {
		for _, f := range _cmd.flags {
			if err := f.Action(c, f); err != nil {
				return err
			}
		}
		return nil
	}
	return _cmd
}

func (c *Command) AddFlag(f *Flag) {
	if f != nil {
		c.flags = append(c.flags, f)
	}
}

func (c *Command) AddFlags(fs []*Flag) {
	for _, f := range fs {
		c.AddFlag(f)
	}
}

func (c *Command) AddSubCommand(cmd *Command) {
	if cmd != nil {
		c.SubCommands = append(c.SubCommands, cmd)
	}
}
