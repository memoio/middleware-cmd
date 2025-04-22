package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const Version = "0.1.0"

var BuildFlag string

var versionCmd = &cli.Command{
	Name:    "version",
	Usage:   "print version",
	Aliases: []string{"V"},
	Action: func(_ *cli.Context) error {
		fmt.Println(Version + "+" + BuildFlag)
		return nil
	},
}
