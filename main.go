package main

import (
	"github.com/gyuhwankim/solverenv/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
