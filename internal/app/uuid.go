package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/acopw-go"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

// UUIDAction is the action for the uuid command.
func UUIDAction(c *cli.Context) error {
	generator := &acopw.UUID{}

	uuid, err := generator.Generate()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	isTerminal := term.IsTerminal(int(os.Stdout.Fd()))

	if isTerminal {
		fmt.Fprintln(c.App.Writer, uuid)

		return nil
	}

	fmt.Fprint(c.App.Writer, uuid)

	return nil
}
