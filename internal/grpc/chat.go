package chatgrpc

import (
	"context"
	"fmt"

	"github.com/nikita-reshetnyak/chat-server/gen/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Chat interface {
	Create(ctx context.Context, usernames []string) (int64, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
	SendMessage(ctx context.Context, from string, text string) (*emptypb.Empty, error)
}
type serverApi struct {
	chat_v1.UnimplementedChatV1Server
	chat Chat
}

func Register(server *grpc.Server, chat Chat) {
	chat_v1.RegisterChatV1Server(server, &serverApi{chat: chat})
}

func (s *serverApi) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	chatID, err := s.chat.Create(ctx, req.Usernames)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &chat_v1.CreateResponse{Id: chatID}, nil
}
func (s *serverApi) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	_, err := s.chat.Delete(ctx, req.Id)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("%w", err)
	}
	return &emptypb.Empty{}, nil
}
func (s *serverApi) SendMessage(ctx context.Context, req *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
