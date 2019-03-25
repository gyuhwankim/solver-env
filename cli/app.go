package cli

import (
	"errors"
	"gopkg.in/urfave/cli.v1"
)

var appData = map[string]string{
	"Name": "solverenv",
	"Usage": "Provide an environment where can solve problems",
	"Version": "v0.1.0",
}

var appCommands = []Commander{
	
} 

var ErrEmptyCliEngine = errors.New("App not contain cli engine.")

type Runner interface {
	Run(args []string) error
}

type App struct {
	cli *cli.App
}

func (app *App) Run(args []string) error {
	if app.cli == nil {
		return ErrEmptyCliEngine		
	}
	
	return app.cli.Run(args)	
}

func NewApp() Runner {
	cliEngine := cli.NewApp()
	cliEngine.Name = appData["Name"]
	cliEngine.Usage = appData["Usage"]
	cliEngine.Version = appData["Version"]
	cliEngine.Commands = []cli.Command{}

	for _, command := range appCommands {
		cliEngine.Commands = append(cliEngine.Commands, command.CliCommand())
	}
	
	app := App{cli: cliEngine}
	
	return &app	
}