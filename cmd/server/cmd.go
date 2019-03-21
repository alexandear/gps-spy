package server

import (
	"log"

	"github.com/devchallenge/spy-api/internal/restapi"
	"github.com/devchallenge/spy-api/internal/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Start a spy http server used by the mobile clients",
	RunE: func(cmd *cobra.Command, args []string) error {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			return errors.Wrap(err, "failed to embeded spec")
		}

		cmd.Long = swaggerSpec.Spec().Info.Description

		pflag.Parse()

		api := operations.NewSpyAPI(swaggerSpec)
		server := restapi.NewServer(api)

		defer func() {
			if err := server.Shutdown(); err != nil {
				log.Fatal(err)
			}
		}()

		server.ConfigureAPI()
		if err := server.Serve(); err != nil {
			return errors.Wrap(err, "failed to serve")
		}

		return nil
	},
}
