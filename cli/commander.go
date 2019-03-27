package cli

import (
	"gopkg.in/urfave/cli.v1"	
)

type Commander interface {
	MetaData() map[string]string
	Flags() []cli.Flag
	Action(ctx *cli.Context) error
}

func createCliCommand(cmd Commander) cli.Command {
	return cli.Command{
		Name: cmd.MetaData()["Name"],
		Usage: cmd.MetaData()["Usage"],
		Flags: cmd.Flags(),
		Action: cmd.Action,
	}
}