package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
