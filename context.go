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
			s := fmt.Sprintln(a...)
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
