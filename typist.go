package typist

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Config of the Typist to be informend when creating one
type Config struct {
	Quiet bool
	Err   io.Writer
	Out   io.Writer
	In    io.Reader
}

// Typist has methods to interact with user
// from command-line with some configurable behaviors
type Typist struct {
	Config *Config
}

// New create a new Typist with the informed Config
func New(config *Config) *Typist {
	if config.In == nil {
		config.In = os.Stdin
	}

	if config.Out == nil {
		config.Out = os.Stdout
	}

	return &Typist{Config: config}
}

func (t *Typist) errput() (errput io.Writer) {
	errput = t.Config.Err
	if errput != nil {
		return
	}

	errput = os.Stderr
	return
}

func (t *Typist) output() (output io.Writer) {
	output = t.Config.Out
	if output != nil {
		return
	}

	output = os.Stdout
	return
}

func (t *Typist) input() (input io.Reader) {
	input = t.Config.In
	if input != nil {
		return
	}

	input = os.Stdin
	return
}

// Outf printing information according Typist Config
func (t *Typist) Outf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(t.output(), format, a...)
}

// Outln printing information according Typist Config
func (t *Typist) Outln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(t.output(), a...)
}

// Readln wait from break-line input and return as string
func (t *Typist) Readln() (string, error) {
	return bufio.NewReader(t.input()).ReadString('\n')
}

// Copy redirect a io.Reader input to the defined output
func (t *Typist) Copy(src io.Reader) (int64, error) {
	return io.Copy(t.Config.Out, src)
}

// Confirm wait for user response returning
// `true` if the answer was y or yes
func (t *Typist) Confirm(message string) bool {
	if t.Config.Quiet {
		return false
	}

	t.Infof("%s [y/N] ", message)

	response, err := t.Readln()
	if err != nil {
		return false
	}

	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

// Infof print string to defined Err output
func (t *Typist) Infof(format string, a ...interface{}) (int, error) {
	return t.Errorf(format, a...)
}

// Errorf print strings to defined Err output
func (t *Typist) Errorf(format string, a ...interface{}) (int, error) {
	if t.Config.Quiet {
		return 0, nil
	}

	return fmt.Fprintf(t.errput(), format, a...)
}

// Infoln print strings to defined Err output
func (t *Typist) Infoln(a ...interface{}) (int, error) {
	return t.Errorln(a...)
}

// Errorln print strings to defined Err output
func (t *Typist) Errorln(a ...interface{}) (int, error) {
	if t.Config.Quiet {
		return 0, nil
	}

	return fmt.Fprintln(t.errput(), a...)
}

// Must checks for error. It accept any number of parameters and check if the last one is nil.
// if not nil it exit the process according the Config
func (t *Typist) Must(params ...interface{}) {
	lastParam := params[len(params)-1]
	if lastParam == nil {
		return
	}

	err := lastParam.(error)
	t.Errorln(err.Error())
	t.Exit(err)
}

// Exitln exit the process with the informed formatting string
func (t *Typist) Exitln(a ...interface{}) {
	_, err := t.Errorln(a...)
	t.Exit(err)
}

// Exitf exit the process with the informed formatting string
func (t *Typist) Exitf(format string, a ...interface{}) {
	_, err := t.Errorf(format, a...)
	t.Exit(err)
}

// Exit exit the process with (TODO) the right ERRNO code
func (t *Typist) Exit(err error) {
	switch err {
	case nil:
		os.Exit(0)
	default:
		os.Exit(1)
	}
}
