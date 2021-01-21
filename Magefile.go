// +build mage

package main

import (
	"fmt"
	"log"

	"github.com/cardil/kn-event/internal"
	"github.com/joho/godotenv"
	"github.com/wavesoftware/go-magetasks/pkg/checks"

	// mage:import
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"
	// mage:import
	_ "github.com/wavesoftware/go-magetasks/container"
)

// Default target is set to binary.
//goland:noinspection GoUnusedGlobalVariable
var Default = magetasks.Binary

func init() { //nolint:gochecknoinits
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	bins := []string{
		internal.PluginName,
		fmt.Sprintf("%s-sender", internal.PluginName),
	}
	for _, bin := range bins {
		config.Binaries = append(config.Binaries, config.Binary{Name: bin})
	}
	config.VersionVariablePath = "github.com/cardil/kn-event/internal.Version"
	checks.Revive()
	checks.Staticcheck()
}
