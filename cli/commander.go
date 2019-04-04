package cli

import (
	"gopkg.in/urfave/cli.v1"
)

type Commander interface {
	MetaDataValue(key string) string
	Flags() []cli.Flag
	Action(ctx *cli.Context) error
}

func createCliCommand(cmd Commander) cli.Command {
	return cli.Command{
		Name:   cmd.MetaDataValue("Name"),
		Usage:  cmd.MetaDataValue("Usage"),
		Flags:  cmd.Flags(),
		Action: cmd.Action,
	}
}
