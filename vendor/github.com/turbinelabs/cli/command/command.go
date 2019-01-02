/*
Copyright 2018 Turbine Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// The command package provides an abstraction for a command-line application
// sub-command, a means to execute code when that sub-command is invoked, a
// means to report success/failure status of said code, and generic
// implementations of help and version sub-commands.
package command

//go:generate mockgen -source $GOFILE -destination mock_$GOFILE -package $GOPACKAGE --write_package_comment=false

import (
	"flag"
	"fmt"
)

// A Runner represents the executable code associated with a Cmd. Typically
// the struct implementing Runner will include whatever state is needed,
// configured by the Flags in the associated Cmd.
type Runner interface {
	// Execute code associated with a Cmd with the given arguments,
	// return exit status. The Cmd is provided here to avoid a circular
	// dependency between the Runner and the Cmd.
	Run(cmd *Cmd, args []string) CmdErr
}

// A Cmd represents a named sub-command for a command-line application.
type Cmd struct {
	Name        string       // Name of the Command and the string to use to invoke it
	Summary     string       // One-sentence summary of what the Command does
	Usage       string       // Usage options/arguments
	Description string       // Detailed description of command
	Flags       flag.FlagSet // Set of flags associated with this Cmd, which typically configure the Runner
	Runner      Runner       // The code to run when this Cmd is invoked
}

// Run invokes the Runner associated with this Cmd, passing the args remaining
// after flags are parsed out. Names of Flags that have been marked as required
// by wrapping their Usage strings in turbinelabs/nonstdlib/flag.Required() and
// for which no value has been set will be returned to the caller in a
// Cmd.BadInput.
func (c *Cmd) Run() CmdErr {
	if c.Runner == nil {
		return c.Error("No Runner specified")
	}

	return c.Runner.Run(c, c.Flags.Args())
}

// BadInputf produces a Cmd-scoped CmdErr with an exit code of 2, based on the
// given format string and args, which are passed to fmt.Sprintf.
func (c *Cmd) BadInputf(format string, args ...interface{}) CmdErr {
	return c.BadInput(fmt.Sprintf(format, args...))
}

// Errorf produces a Cmd-scoped CmdErr with an exit code of 1, based on the
// given format string and args, which are passed to fmt.Sprintf.
func (c *Cmd) Errorf(format string, args ...interface{}) CmdErr {
	return c.Error(fmt.Sprintf(format, args...))
}

// BadInput produces a Cmd-scoped CmdErr with an exit code of 2, based on the
// given args, which are passed to fmt.Sprint.
func (c *Cmd) BadInput(args ...interface{}) CmdErr {
	return CmdErr{c, CmdErrCodeBadInput, fmt.Sprintf("%s: %s", c.Name, fmt.Sprint(args...))}
}

// Error produces a Cmd-scoped CmdErr with an exit code of 1, based on the
// given args, which are passed to fmt.Sprint.
func (c *Cmd) Error(args ...interface{}) CmdErr {
	return CmdErr{c, CmdErrCodeError, fmt.Sprintf("%s: %s", c.Name, fmt.Sprint(args...))}
}

// CmdErrCode is the exit code for the application
type CmdErrCode uint32

const (
	//CmdErrCodeNoError is the CmdErrCode returns when there is no error
	CmdErrCodeNoError CmdErrCode = 0

	// CmdErrCodeError is the CmdErrCode returned for a generic error
	CmdErrCodeError = 1

	// CmdErrCodeBadInput is the CmdErrorCode returned for bad input
	CmdErrCodeBadInput = 2 // Bad Input Error
)

// CmdErr represents the exit status of a Cmd.
type CmdErr struct {
	Cmd     *Cmd       // The Cmd that produced the exit status. Can be nil for global errors
	Code    CmdErrCode // The exit code
	Message string     // Additional information if the Code is non-zero
}

// IsError returns true if the exit code is non-zero
func (err CmdErr) IsError() bool {
	return err.Code != CmdErrCodeNoError
}

var cmdErrNoErr = CmdErr{nil, CmdErrCodeNoError, ""}

// NoError returns the singleton unscoped CmdErr with an exit code of 0
func NoError() CmdErr {
	return cmdErrNoErr
}
