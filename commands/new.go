package commands

import "gopkg.in/urfave/cli.v1"

func New() cli.Command {
	return cli.Command{
		Name:    "new",
		Usage:   "Create new environment ex) ps-cli new <options> <src>",
		Action: func(context *cli.Context) error {
			template := context.String("templates")
			src := context.Args().Get(0)

			makeDir(template, src)
			copyTemplateFiles(template, src)

			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "templates, t",
				Value: "cpp",
				Usage: "Select env template",
			},
		},
	}
}

func makeDir(template, src string) {

}

func copyTemplateFiles(template, src string) {

}

