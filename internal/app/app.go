// Package app is the main package for the application.
package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/acopw-cli/internal/meta"
	"git.sr.ht/~jamesponddotco/acopw-go"
	"github.com/urfave/cli/v2"
)

// Run is the entry point for the application.
func Run() int {
	app := cli.NewApp()
	app.Name = meta.Name
	app.Version = meta.Version
	app.Usage = meta.Description
	app.HideHelpCommand = true

	app.Commands = []*cli.Command{
		{
			Name:    "diceware",
			Aliases: []string{"d"},
			Usage:   "generate a random diceware password",
			Action:  DicewareAction,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:    "length",
					Aliases: []string{"l"},
					Usage:   "the number of words in the password",
					Value:   uint(acopw.DefaultDicewareLength),
				},
				&cli.StringFlag{
					Name:    "separator",
					Aliases: []string{"s"},
					Usage:   "the string to use to separate words",
					Value:   " ",
				},
				&cli.BoolFlag{
					Name:    "capitalized",
					Aliases: []string{"c"},
					Usage:   "whether to capitalize a random word in the password",
					Value:   true,
				},
			},
		},
		{
			Name:    "random",
			Aliases: []string{"r"},
			Usage:   "generate a random password",
			Action:  RandomAction,
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:    "length",
					Aliases: []string{"l"},
					Usage:   "the number of characters in the password",
					Value:   acopw.DefaultRandomLength,
				},
				&cli.BoolFlag{
					Name:    "numbers",
					Aliases: []string{"n"},
					Usage:   "whether to include numbers in the password",
					Value:   true,
				},
				&cli.BoolFlag{
					Name:    "uppercase",
					Aliases: []string{"u"},
					Usage:   "whether to include uppercase letters in the password",
					Value:   true,
				},
				&cli.BoolFlag{
					Name:    "lowercase",
					Aliases: []string{"L"},
					Usage:   "whether to include lowercase letters in the password",
					Value:   true,
				},
				&cli.BoolFlag{
					Name:    "symbols",
					Aliases: []string{"s"},
					Usage:   "whether to include symbols in the password",
					Value:   true,
				},
				&cli.StringSliceFlag{
					Name:    "exclude",
					Aliases: []string{"e"},
					Usage:   "a character to exclude from the password",
					Value:   nil,
				},
			},
		},
		{
			Name:    "pin",
			Aliases: []string{"p"},
			Usage:   "generate a random numeric PIN",
			Action:  PINAction,
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:    "length",
					Aliases: []string{"l"},
					Usage:   "the number of digits in the PIN",
					Value:   acopw.DefaultPINLength,
				},
			},
		},
		{
			Name:    "uuid",
			Aliases: []string{"u"},
			Usage:   "generate a random UUID, a.k.a, UUIDv4",
			Action:  UUIDAction,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		return 1
	}

	return 0
}
