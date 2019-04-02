package cli

import (
	"errors"
	"gopkg.in/urfave/cli.v1"
)

const (
	DATA_NAME    = "Name"
	DATA_USAGE   = "Usage"
	DATA_VERSION = "Version"
)

// metadata of cli included items like name, usage and version.
var appData = map[string]string{
	DATA_NAME:    "solverenv",
	DATA_USAGE:   "Provide an environment where can solve problems",
	DATA_VERSION: "v0.1.0",
}

var appCommands = []Commander{
	newInitCommand(),
}

var (
	ErrEmptyCliEngine = errors.New("App not contain cli engine.")
	ErrNotContainItem = errors.New("Not contain item in collection.")
)

type Runner interface {
	Run(args []string) error
}

type App struct {
	cli      *cli.App
	metaData map[string]string
}

func (app App) Name() (string, error) {
	return app.getData(DATA_NAME)
}

func (app App) Usage() (string, error) {
	return app.getData(DATA_USAGE)
}

func (app App) Version() (string, error) {
	return app.getData(DATA_VERSION)
}

func (app App) getData(key string) (string, error) {
	item, exists := app.metaData[key]
	if !exists {
		return "", ErrNotContainItem
	}

	return item, nil
}

func (app *App) Run(args []string) error {
	if app.cli == nil {
		return ErrEmptyCliEngine
	}

	return app.cli.Run(args)
}

func NewApp() Runner {
	app := newAppInstance(cli.NewApp(), appData)

	updateCliMetadata(app)

	return &app
}

func newAppInstance(cliEngine *cli.App, appMetadata map[string]string) *App {
	return App{
		cli:      cliEngine,
		metaData: appMetadata,
	}
}

func updateCliMetadata(app *App) error {
	cli.Name = app.Name()
	cli.Usage = app.Usage()
	cli.Version = app.Version()

	return nil
}
