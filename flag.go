package cli_mate

import (
	"flag"
	"github.com/urfave/cli"
)

// FlagAction the function for Flag action
type FlagAction func(c *Context, f *Flag) ExitCoder

// Flag is a common interface related to parsing flags in cli.
type Flag struct {
	Flag     cli.Flag
	FlagName string
	Action   FlagAction
}

// Apply Flag settings to the given flag set
func (f *Flag) Apply(set *flag.FlagSet) {
	f.Flag.Apply(set)
}

// GetName get the name of Flag
func (f Flag) GetName() string {
	return f.FlagName
}

// register get Flag After the registration
func (f *Flag) register() *Flag {
	return f
}
