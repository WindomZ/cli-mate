package cli_mate

import (
	"flag"
	"github.com/urfave/cli"
)

type Flag struct {
	Flag     cli.Flag
	FlagName string
}

func (f *Flag) Apply(set *flag.FlagSet) {
	f.Flag.Apply(set)
}

func (f Flag) GetName() string {
	return f.FlagName
}
