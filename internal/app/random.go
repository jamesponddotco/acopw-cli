package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/acopw-go"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

// RandomAction is the action for the random command.
func RandomAction(c *cli.Context) error {
	generator := &acopw.Random{
		ExcludedCharset: c.StringSlice("excluded"),
		Length:          int(c.Uint("length")),
		UseLower:        c.Bool("lowercase"),
		UseUpper:        c.Bool("uppercase"),
		UseNumbers:      c.Bool("numbers"),
		UseSymbols:      c.Bool("symbols"),
	}

	password, err := generator.Generate()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	isTerminal := term.IsTerminal(int(os.Stdout.Fd()))

	if isTerminal {
		fmt.Fprintln(c.App.Writer, password)

		return nil
	}

	fmt.Fprint(c.App.Writer, password)

	return nil
}
