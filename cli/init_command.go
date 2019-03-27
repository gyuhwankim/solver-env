package cli

import (
	"gopkg.in/urfave/cli.v1"			
)

type initCommand struct {
	metaData map[string]string
}

func (cmd initCommand) MetaData(key string) string {
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
			"Name": "init",
			"Usage": "Create new root workspace",
		},
	}
}