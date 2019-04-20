package violator

import (
	"context"

	"google.golang.org/grpc"

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

func (s *Service) Numbers() ([]string, error) {
	return []string{}, nil
}
