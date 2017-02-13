package cli_mate

import (
	"errors"
	"github.com/urfave/cli"
)

// BreakExit just for break action of command execute
var BreakExit *ExitError = &ExitError{
	ExitError: *cli.NewExitError("break", 0),
}

type ExitCoder interface {
	cli.ExitCoder
	GetError() error
	IsBreak() bool
}

type ExitError struct {
	cli.ExitError
}

// NewExitError makes a new *ExitError
func NewExitError(err error, exitCode int) *ExitError {
	if err == nil {
		return &ExitError{
			ExitError: *cli.NewExitError("", exitCode),
		}
	}
	return &ExitError{
		ExitError: *cli.NewExitError(err.Error(), exitCode),
	}
}

// NewDefaultExitError makes a new *ExitError with no exit code
func NewDefaultExitError(err error) *ExitError {
	return NewExitError(err, 1)
}

func (e ExitError) GetError() error {
	return errors.New(e.Error())
}

func (e ExitError) IsBreak() bool {
	return e.ExitCode() == 0 && e.Error() == "break"
}
