package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	sendCmd = func() *cobra.Command {
		c := &cobra.Command{
			Use:   "send",
			Short: "Builds and sends a CloudEvent to recipient",
			RunE: func(cmd *cobra.Command, args []string) error {
				return errors.New("not yet implemented")
			},
		}
		addBuilderFlags(c)
		return c
	}()
)
