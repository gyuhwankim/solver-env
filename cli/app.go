package cli

import (
	"gopkg.in/urfave/cli.v1"
)

var appData = map[string]string{
	"Name": "solverenv",
	"Usage": "Provide an environment where can solve problems",
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
	
	app := App{cliEngine}
	
	return &app	
}