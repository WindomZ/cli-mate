package cli_mate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

var (
	app    *App
	result string
)

var (
	flag_lang = Flag{
		Flag: cli.StringFlag{
			Name:  "lang",
			Value: "chinese",
			Usage: "language for the greeting",
		},
		FlagName: "lang",
		Action: func(c *Context, f *Flag) ExitCoder {
			println("flag_lang:", fmt.Sprintf("%#v", c.Args()))
			value := c.String(f.GetName())
			if len(value) == 0 {
				return nil
			}
			println("flag_lang:", value)

			name := "Null"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}

			switch value {
			case "spanish":
				result = c.Print("Hola", name)
			case "chinese":
				result = c.Print("你好", name)
			default:
				result = c.Print("Hello", name)
			}

			return nil
		},
	}

	command_add = Command{
		Command: cli.Command{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
		},
		Action: func(c *Context) ExitCoder {
			println("command_add:", fmt.Sprintf("%#v", c.Args()))
			result = c.Print("Added:", c.Args().First())
			return nil
		},
	}
)

func TestApp_NewApp(t *testing.T) {
	app = NewApp("wahaha", "0.0.1")
	assert.NotNil(t, app)
}

func TestApp_SimpleFlag(t *testing.T) {
	TestApp_NewApp(t)

	app.AddFlag(flag_lang)

	app.Run([]string{app.Name(), "-lang", "english", "world"})
	assert.Equal(t, result, "Hello world")
}

func TestApp_SimpleCommand(t *testing.T) {
	TestApp_NewApp(t)

	app.AddCommand(command_add)

	app.Run([]string{app.Name(), "add", "codes"})
	assert.Equal(t, result, "Added: codes")

	//command_add.AddFlag(flag_lang)
	//
	//app.Run([]string{"", "add", "world", "-lang", "english", "world"})
	//assert.Equal(t, result, "Added: world")
}
