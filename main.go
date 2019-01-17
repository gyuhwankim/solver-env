package main

import (
	"fmt"
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
	app.Name = "solver-env"
	app.Usage = "Provide an environment where can solve problems"
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Usage:   "Create new environment ex) ps-cli new <options> <src>",
			Action: func(context *cli.Context) error {
				template := context.String("templates")
				src := context.Args().Get(0)

				makeWorkspace(template, src)
				copyTemplateFiles(template, src)

				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "templates, t",
					Value: "cpp",
					Usage: "Choice for the template",
				},
			},
		},
	}

	return app
}

func makeWorkspace(template, src string) {
	fmt.Println(template)
	fmt.Println(src)
}

func copyTemplateFiles(template, src string) {

}
