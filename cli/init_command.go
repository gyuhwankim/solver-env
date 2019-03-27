package cli

import (
	"gopkg.in/urfave/cli.v1"			
)

var commandData := map[string]string{
	"Name": "init",
	"Usage": "Create new root workspace",
}
type InitCommand struct {
	
}

func (cmd *InitCommand) CliCommand() cli.Command {
	return cli.Command{
		Name: commandData["Name"]
		Usage: commandData["Usage"]
	}	
} 