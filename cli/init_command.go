package cli

import (
	"gopkg.in/urfave/cli.v1"
)

type initCommand struct {
	metaData map[string]string
}

func (cmd initCommand) MetaDataValue(key string) string {
	return cmd.metaData[key]
}

func (cmd initCommand) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (cmd initCommand) Action(ctx *cli.Context) error {
	return nil
}

func newInitCommand() *initCommand {
	return &initCommand{
		metaData: map[string]string{
			NAME_KEY:  "init",
			USAGE_KEY: "Create new root workspace",
		},
	}
}
