package app

import (
	"context"

	grpcapp "github.com/nikita-reshetnyak/chat-server/internal/app/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type App struct {
	GrpcServer *grpcapp.App
}
type chat struct{}

func (s *chat) CreateRequest(ctx context.Context, usernames []string) (int64, error) {
	return 0, nil
}
func (s *chat) DeleteRequest(ctx context.Context, id int64) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *chat) SendMessageRequest(ctx context.Context, from string, text string) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func New(grpcPort string, configPath string) *App {
	app := grpcapp.New(&chat{}, grpcPort)
	return &App{GrpcServer: app}
}
