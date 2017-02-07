package cli_mate

import "github.com/urfave/cli"

type Context struct {
	cli.Context
}

func NewContext(c *cli.Context) *Context {
	return &Context{
		Context: *c,
	}
}
