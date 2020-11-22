package cmd

import (
	"github.com/cardil/kn-event/internal/cli"
	"github.com/spf13/cobra"
)

var (
	target = &cli.TargetArgs{}

	sendCmd = func() *cobra.Command {
		c := &cobra.Command{
			Use:   "send",
			Short: "Builds and sends a CloudEvent to recipient",
			RunE: func(cmd *cobra.Command, args []string) error {
				options.OutWriter = cmd.OutOrStdout()
				options.ErrWriter = cmd.ErrOrStderr()
				ce, err := cli.CreateWithArgs(eventArgs)
				if err != nil {
					return err
				}
				return cli.Send(ce, target, options)
			},
		}
		addBuilderFlags(c)
		c.Flags().StringVarP(
			&target.URL, "to-url", "u", "",
			`Specify an URL to send event to. This option can't be used with 
--to option.`,
		)
		c.Flags().StringVarP(
			&target.Addressable, "to", "r", "",
			`Specify an addressable resource to send event to. This argument
takes format kind:apiVersion:name for named resources or
kind:apiVersion:labelKey1=value1,labelKey2=value2 for matching via a
label selector. This option can't be used with --to-url option.'`,
		)
		c.Flags().StringVarP(
			&target.Namespace, "namespace", "n", "",
			`Specify a namespace of addressable resource defined with --to
option. If this option isn't specified a current context namespace will be used
to find addressable resource. This option can't be used with --to-url option.'`,
		)
		return c
	}()
)
