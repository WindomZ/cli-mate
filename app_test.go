package cli_mate

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestApp_SimpleFlag(t *testing.T) {
	var result string

	app := NewApp()

	app.AddFlag(&Flag{
		Flag: cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
		FlagName: "lang",
		Action: func(c *Context, f *Flag) ExitCoder {
			value := c.String(f.GetName())
			if len(value) == 0 {
				return nil
			}

			name := "Null"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}

			switch value {
			case "spanish":
				result = c.Print("Hola", name)
				break
			default:
				result = c.Print("Hello", name)
				break
			}

			return nil
		},
	})

	app.Run([]string{"", "-lang", "english", "world"})

	assert.Equal(t, result, "Hello world")
}
