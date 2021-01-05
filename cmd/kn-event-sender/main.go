package main

import (
	"os"

	"github.com/cardil/kn-event/internal/cli/ics"
	"github.com/cardil/kn-event/internal/cli/retcode"
	"github.com/cardil/kn-event/internal/configuration"
)

// ExitFunc will be used to exit Go process in case of error.
var ExitFunc = os.Exit // nolint:gochecknoglobals

func main() {
	configuration.ConfigureSender()
	if err := ics.SendFromEnv(); err != nil {
		ExitFunc(retcode.Calc(err))
	}
}
