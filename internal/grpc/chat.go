package chatgrpc

import (
	"context"

	"github.com/nikita-reshetnyak/chat-server/gen/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Chat interface {
	CreateRequest(ctx context.Context, usernames []string) (int64, error)
	DeleteRequest(ctx context.Context, id int64) (*emptypb.Empty, error)
	SendMessageRequest(ctx context.Context, from string, text string) (*emptypb.Empty, error)
}
type serverApi struct {
	chat_v1.UnimplementedChatV1Server
	chat Chat
}

func Register(server *grpc.Server, chat Chat) {
	chat_v1.RegisterChatV1Server(server, &serverApi{chat: chat})
}

func (s *serverApi) CreateRequest(ctx context.Context, usernames []string) (int64, error) {
	return 0, nil
}
func (s *serverApi) DeleteRequest(ctx context.Context, id int64) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *serverApi) SendMessageRequest(ctx context.Context, from string, text string) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
