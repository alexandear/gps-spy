package server

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tidwall/buntdb"
	"google.golang.org/grpc"

	"github.com/devchallenge/spy-api/internal/gen/restapi"
	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/service/gps"
	"github.com/devchallenge/spy-api/internal/service/handler"
	"github.com/devchallenge/spy-api/internal/service/specnomery"
	"github.com/devchallenge/spy-api/internal/service/together"
	"github.com/devchallenge/spy-api/internal/service/violator"
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

		log.Printf("Opening db=%s", config.buntdbPath)
		db, err := buntdb.Open(config.buntdbPath)
		if err != nil {
			return errors.Wrap(err, "failed to open buntdb in memory")
		}

		var specnomeryGRPC specnomery.AllowedUsersClient
		conn, err := grpc.Dial(config.specnomeryServer, grpc.WithInsecure())
		if err != nil {
			specnomeryGRPC = specnomery.EmptyAllowedUsersClient{}
			log.Printf("Did not connect: %v", err)
		} else {
			specnomeryGRPC = specnomery.NewAllowedUsersClient(conn)
			log.Printf("GRPC connected to=%s", config.specnomeryServer)
		}
		defer util.Close(conn)

		storage := storage.New(db)
		defer util.Close(storage)
		gps := gps.New(storage)
		together := together.New(storage)
		violator := violator.New(specnomeryGRPC)
		handler := handler.New(gps, together, violator)
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
