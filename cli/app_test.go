package cli

import (
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
