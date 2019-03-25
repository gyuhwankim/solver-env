package cli

import (
	"gopkg.in/urfave/cli.v1"	
)

type Commander interface {
	CliCommand() cli.Command	
}