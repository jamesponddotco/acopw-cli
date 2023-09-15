package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/acopw-go"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

// DicewareAction is the action for the diceware command.
func DicewareAction(c *cli.Context) error {
	generator := &acopw.Diceware{
		Length:     int(c.Uint("length")),
		Separator:  c.String("separator"),
		Capitalize: c.Bool("capitalized"),
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
