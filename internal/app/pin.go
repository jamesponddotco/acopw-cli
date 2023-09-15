package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/acopw-go"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

// PINAction is the action for the pin command.
func PINAction(c *cli.Context) error {
	generator := &acopw.PIN{
		Length: int(c.Uint("length")),
	}

	pin, err := generator.Generate()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	isTerminal := term.IsTerminal(int(os.Stdout.Fd()))

	if isTerminal {
		fmt.Fprintln(c.App.Writer, pin)

		return nil
	}

	fmt.Fprint(c.App.Writer, pin)

	return nil
}
