package cli

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/urfave/cli.v1"
	"testing"
)

func TestRunWhenCliNil(t *testing.T) {
	app := App{cli: nil}
	expected := ErrEmptyCliEngine

	actual := app.Run([]string{})

	if expected != actual {
		t.Errorf("expected: %t, but actual: %t", expected, actual)
	}
}

func TestUpdateAppMetadata(t *testing.T) {
	metaData := map[string]string{
		"Name":    "test-solverenv",
		"Usage":   "just test solverenv",
		"Version": "test version",
	}

	cliEngine := cli.NewApp()
	app := newAppInstance(cliEngine, metaData)
	
	if err := updateCliMetadata(app); err != nil {
		t.Errorf("Occurred error: %s", err)
	}

	assert.Equal(t, cliEngine.Name, metaData["Name"])
	assert.Equal(t, cliEngine.Usage, metaData["Usage"])
	assert.Equal(t, cliEngine.Version, metaData["Version"])
}
