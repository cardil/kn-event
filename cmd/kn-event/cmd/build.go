package cmd

import (
	"fmt"
	"time"

	"github.com/cardil/kn-event/internal"
	"github.com/spf13/cobra"
)

var (
	buildTypeFlag        string
	buildTypeFlagDefault = "dev.knative.cli.plugin.event.generic"
	buildCmd             = func() *cobra.Command {
		c := &cobra.Command{
			Use:   "build",
			Short: "Builds a CloudEvent and print it to stdout",
			Run: func(cmd *cobra.Command, args []string) {
				now := time.Now().Format(time.RFC3339Nano)
				if len(args) > 0 {
					cmd.Printf(`{
						"specversion": "1.0",
						"type": "org.example.ping",
						"id": "71830",
						"time": "%s",
						"source": "/api/v1/ping",
						"dataContentType": "application/json",
						"data": {
							"person": {
								"name": "Chris",
								"email": "ksuszyns@example.com"
							},
							"ping": 123,
							"ref": "321"
						}
					}
					`, now)
				} else {
					cmd.Printf(`{
						"specversion": "1.0",
						"type": "dev.knative.cli.plugin.event.generic",
						"id": "1566bdca-a3ed-4365-a871-6bafbbbb4548",
						"time": "%s",
						"source": "%s/%s",
						"dataContentType": "application/json",
						"data": {}
					}
					`, now, internal.PluginName, internal.Version)
				}
			},
		}
		c.Flags().StringVarP(
			&buildTypeFlag, "type", "t", buildTypeFlagDefault,
			fmt.Sprintf("specify a type of a CloudEvent (default is: %s)", buildTypeFlagDefault),
		)
		return c
	}()
)
