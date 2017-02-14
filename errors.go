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

// NewExitError makes a new *ExitError with error and code
func NewExitError(err error, exitCode ...int) *ExitError {
	if err == nil {
		return NewExitStringError("", exitCode...)
	}
	return NewExitStringError(err.Error(), exitCode...)
}

// NewExitStringError makes a new *ExitError with string and code
func NewExitStringError(err string, exitCodes ...int) *ExitError {
	var exitCode int = 1
	if exitCodes != nil && len(exitCodes) != 0 {
		exitCode = exitCodes[0]
	}
	return &ExitError{
		ExitError: *cli.NewExitError(err, exitCode),
	}
}

func (e ExitError) GetError() error {
	return errors.New(e.Error())
}

func (e ExitError) IsBreak() bool {
	return e.ExitCode() == 0 && e.Error() == "break"
}
