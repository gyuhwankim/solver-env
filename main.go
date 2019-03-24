package main

import (
	"github.com/gyuhwankim/solverenv/command"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := createApplication()

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func createApplication() *cli.App {
	app := cli.NewApp()
	app.Name = "solverenv"
	app.Usage = "Provide an environment where can solve problems"
	app.Commands = []cli.Command{
		command.CreateNew(),
	}

	return app
}
