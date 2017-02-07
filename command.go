package cli_mate

import "github.com/urfave/cli"

type Command struct {
	Command     *cli.Command
	Action      func(c *Context) error
	flags       []*Flag
	SubCommands []*Command
}

func NewCommand(cmd *cli.Command) *Command {
	return &Command{
		Command: cmd,
	}
}
