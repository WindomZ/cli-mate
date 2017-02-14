package cli_mate

import (
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
			value := c.String(f.Name())
			if len(value) == 0 {
				return nil
			}

			name := "Null"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}

			switch value {
			case "spanish":
				result = c.Sprint("Hola", name)
			case "chinese":
				result = c.Sprint("你好", name)
			default:
				result = c.Sprint("Hello", name)
			}

			return nil
		},
	}

	command_add = Command{
		Command: cli.Command{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a word",
		},
		Action: func(c *Context) ExitCoder {
			result = c.Sprint("Added:", c.Args().First())
			return nil
		},
	}

	command_remove = Command{
		Command: cli.Command{
			Name:    "remove",
			Aliases: []string{"s"},
			Usage:   "remove a word",
		},
		Action: func(c *Context) ExitCoder {
			result = c.Sprint("Removed:", c.Args().First())
			return nil
		},
	}
)

func TestApp_NewApp(t *testing.T) {
	app = NewApp("wahaha", "just for testing", "0.0.1")
	assert.NotNil(t, app)
}

func TestApp_SimpleFlag(t *testing.T) {
	app.AddFlag(flag_lang)

	testApp_SimpleFlag(t)
}

func testApp_SimpleFlag(t *testing.T) {
	app.Run([]string{app.Name(), "-lang", "english", "world"})
	assert.Equal(t, result, "Hello world")
}

func TestApp_SimpleCommand(t *testing.T) {
	command_add.AddFlag(flag_lang)
	app.AddCommand(command_add)

	testApp_SimpleCommand(t)
}

func testApp_SimpleCommand(t *testing.T) {
	testApp_SimpleFlag(t)

	app.Run([]string{app.Name(), "add", "codes"})
	assert.Equal(t, result, "Added: codes")

	app.Run([]string{"", "add", "world", "-lang", "english", "world"})
	assert.Equal(t, result, "Added: world")
}

func TestApp_SimpleSubcommand(t *testing.T) {
	app.Clear()

	app.AddFlag(flag_lang)

	command_add.AddSubCommand(command_remove)
	app.AddCommand(command_add)

	testApp_SimpleSubcommand(t)
}

func testApp_SimpleSubcommand(t *testing.T) {
	testApp_SimpleCommand(t)

	app.Run([]string{"", "add", "remove", "world"})
	assert.Equal(t, result, "Removed: world")
}
