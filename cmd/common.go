package cmd

import "github.com/urfave/cli/v2"

var CommonCmd = []*cli.Command{
	checkCmd,
	signmCmd,
	signTxCmd,
	sendCmd,
	gendidCmd,
	loginCmd,
	idCmd,
	versionCmd,
}
