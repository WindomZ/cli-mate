package cli_mate

import (
	"flag"
	"github.com/urfave/cli"
)

type FlagAction func(c *Context, f *Flag) ExitCoder

type Flag struct {
	Flag     cli.Flag
	FlagName string
	Action   FlagAction
}

func (f *Flag) Apply(set *flag.FlagSet) {
	f.Flag.Apply(set)
}

func (f Flag) GetName() string {
	return f.FlagName
}

func (f *Flag) register() *Flag {
	return f
}
