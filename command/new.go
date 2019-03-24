package command

import (
	"fmt"
	"github.com/otiai10/copy"
	"gopkg.in/urfave/cli.v1"
	"os"
)

const commandName = "new"
const commandUsage = "Create new environment ex) solverenv new <options> <src>"

func CreateNew() cli.Command {
	return cli.Command{
		Name:  commandName,
		Usage: commandUsage,
		Action: func(context *cli.Context) error {
			template := context.String("templates")
			path := context.Args().Get(0)

			return createEnvironment(template, path)
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

func createEnvironment(template, path string) error {
	src := fmt.Sprintf("templates/%s", template)
	dest := fmt.Sprintf("src/%s", path)

	if !isExistsDir(src) {
		return fmt.Errorf(fmt.Sprintf("Does not exists template on %s", src))
	}

	copyDir(src, dest)

	return nil
}

func isExistsDir(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func copyDir(src, dest string) error {
	return copy.Copy(src, dest)
}
