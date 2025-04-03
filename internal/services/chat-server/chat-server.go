package chatserver

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Repository interface {
	Create(ctx context.Context, usernames []string) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, from string, text string) error
}
type ChatServer struct {
	repo Repository
}

func New(repo Repository) *ChatServer {
	return &ChatServer{repo: repo}
}
func (s *ChatServer) Create(ctx context.Context, usernames []string) (int64, error) {
	chatID, err := s.repo.Create(ctx, usernames)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	return chatID, nil
}
func (s *ChatServer) Delete(ctx context.Context, id int64) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *ChatServer) SendMessage(ctx context.Context, from string, text string) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
