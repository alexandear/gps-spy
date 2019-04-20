package specnomery

import (
	"context"

	"google.golang.org/grpc"
)

type EmptyAllowedUsersClient struct {
	AllowedUsersClient
}

func Check(_ context.Context, _ *AllowedUsersRequest, _ ...grpc.CallOption) (*AllowedUsersReply, error) {
	return &AllowedUsersReply{Result: true}, nil
}
