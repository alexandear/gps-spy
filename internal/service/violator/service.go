package violator

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/devchallenge/spy-api/internal/model"
	"github.com/devchallenge/spy-api/internal/service/specnomery"
)

type Specnomery interface {
	Check(ctx context.Context, in *specnomery.AllowedUsersRequest, opts ...grpc.CallOption) (*specnomery.AllowedUsersReply, error)
}

type Service struct {
	specnomery Specnomery
}

func New(specnomery Specnomery) *Service {
	return &Service{
		specnomery: specnomery,
	}
}

func (s *Service) Numbers(from, to time.Time, minSpeed int, minCoordinate, maxCoordinate model.Coordinate) ([]string, error) {
	// TODO
	return []string{}, nil
}
