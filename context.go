package cli_mate

import (
	"fmt"
	"github.com/urfave/cli"
	"strings"
)

type Context struct {
	cli.Context
	printFunc   func(a ...interface{}) string
	printlnFunc func(a ...interface{}) string
}

func NewContext(c *cli.Context) *Context {
	return &Context{
		Context: *c,
		printFunc: func(a ...interface{}) string {
			s := strings.TrimSuffix(fmt.Sprintln(a...), "\n")
			fmt.Println(s)
			return s
		},
		printlnFunc: func(a ...interface{}) string {
			s := strings.TrimSuffix(fmt.Sprintln(a...), "\n")
			fmt.Println(s)
			return s
		},
	}
}

func (c *Context) Print(a ...interface{}) string {
	return c.printFunc(a...)
}

func (c *Context) Println(a ...interface{}) string {
	return c.printlnFunc(a...)
}

func (c *Context) Sprint(a ...interface{}) string {
	return strings.TrimSuffix(fmt.Sprintln(a...), "\n")
}

func (c *Context) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}

func (c *Context) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func (c *Context) ShowAppHelp() error {
	return cli.ShowAppHelp(&c.Context)
}

func (c *Context) ShowCommandHelp(command string) error {
	return cli.ShowCommandHelp(&c.Context, command)
}

func (c *Context) ShowSubcommandHelp() error {
	return cli.ShowSubcommandHelp(&c.Context)
}
