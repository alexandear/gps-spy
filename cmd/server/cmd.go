package server

import (
	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tidwall/buntdb"

	"github.com/devchallenge/spy-api/internal/gen/restapi"
	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/service/gps"
	"github.com/devchallenge/spy-api/internal/service/handler"
	"github.com/devchallenge/spy-api/internal/service/together"
	"github.com/devchallenge/spy-api/internal/storage"
	"github.com/devchallenge/spy-api/internal/util"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Start a spy http server used by the mobile clients",
	RunE: func(cmd *cobra.Command, args []string) error {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			return errors.Wrap(err, "failed to embedded spec")
		}

		cmd.Long = swaggerSpec.Spec().Info.Description

		pflag.Parse()

		api := operations.NewSpyAPI(swaggerSpec)
		server := server{Server: restapi.NewServer(api)}
		defer util.Close(server)

		db, err := buntdb.Open(":memory:")
		if err != nil {
			return errors.Wrap(err, "failed to open buntdb in memory")
		}

		storage := storage.New(db)
		defer util.Close(storage)
		gps := gps.New(storage)
		together := together.New(storage)
		handler := handler.New(gps, together)
		handler.ConfigureHandlers(api)
		if err := server.Serve(); err != nil {
			return errors.Wrap(err, "failed to serve")
		}

		return nil
	},
}

type server struct {
	*restapi.Server
}

func (s server) Close() error {
	return s.Shutdown()
}
