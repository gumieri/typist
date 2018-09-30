package typist

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Typist has methods to interact with user
// from command-line with some configurable behaviors
type Typist struct {
	Quiet bool
	Out   io.Writer
	In    io.Reader
	Err   io.Writer
}

// Printf printing information according Typist configurations
func (t *Typist) Printf(format string, a ...interface{}) (n int, err error) {
	if t.Quiet {
		return
	}

	output := t.Out
	if output == nil {
		output = os.Stdout
	}

	return fmt.Fprintf(output, format, a...)
}

// Println printing information according Typist configurations
func (t *Typist) Println(a ...interface{}) (n int, err error) {
	if t.Quiet {
		return
	}

	output := t.Out
	if output == nil {
		output = os.Stdout
	}

	return fmt.Fprintln(output, a...)
}

// Readln wait from break-line input and return as string
func (t *Typist) Readln() (string, error) {
	input := t.In
	if input == nil {
		input = os.Stdin
	}

	return bufio.NewReader(input).ReadString('\n')
}

// Confirm wait for user response returning
// `true` if the answer was y or yes
func (t *Typist) Confirm(message string) bool {
	if t.Quiet {
		return false
	}

	t.Printf("%s [y/N] ", message)

	response, err := t.Readln()
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

// Errorf print strings to defined Err output
func (t *Typist) Errorf(format string, a ...interface{}) (n int, err error) {
	output := t.Err
	if output == nil {
		output = os.Stderr
	}

	return fmt.Fprintf(output, format, a...)
}

// Errorln print strings to defined Err output
func (t *Typist) Errorln(a ...interface{}) (n int, err error) {
	output := t.Err
	if output == nil {
		output = os.Stderr
	}

	return fmt.Fprintln(output, a...)
}

// Must checks for error.
// if not nil it exit the process according the configurations
func (t *Typist) Must(err error) {
	if err == nil {
		return
	}

	t.Errorln(err.Error())
	t.Finish(err)
}

// Finish exit the process with (TODO) the right ERRNO code
func (t *Typist) Finish(err error) {
	switch err {
	case nil:
		os.Exit(0)
	default:
		os.Exit(1)
	}
}
