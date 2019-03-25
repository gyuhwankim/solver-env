package cli

import (
	"gopkg.in/urfave/cli.v1"
)

var appData = map[string]string{
	"Name": "solverenv",
	"Usage": "Provide an environment where can solve problems",
	"Version": "v0.1.0",
}

type Runner interface {
	Run(args []string) error
}

type App struct {
	cli *cli.App
}

func (app *App) Run(args []string) error {
	return app.cli.Run(args)	
}

func NewApp() Runner {
	cliEngine := cli.NewApp()
	cliEngine.Name = appData["Name"]
	cliEngine.Usage = appData["Usage"]
	cliEngine.Version = appData["Version"]
	cliEngine.Commands = []cli.Command{
				
	}
	
	app := App{cli: cliEngine}
	
	return &app	
}