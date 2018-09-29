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
}

// Type printing information according Typist configurations
func (t *Typist) Type(format string, a ...interface{}) (n int, err error) {
	if t.Quiet {
		return
	}

	output := t.Out
	if output == nil {
		output = os.Stdout
	}

	return fmt.Fprintf(output, format, a...)
}

// ReadLine wait from break-line input and return as string
func (t *Typist) ReadLine() (string, error) {
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

	t.Type("%s [y/N] ", message)

	response, err := t.ReadLine()
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}
