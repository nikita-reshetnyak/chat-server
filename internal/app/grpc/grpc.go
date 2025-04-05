package grpcapp

import (
	"fmt"
	"net"

	chatgrpc "github.com/nikita-reshetnyak/chat-server/internal/grpc"
	"google.golang.org/grpc"
)

type App struct {
	grpcServer *grpc.Server
	port       string
}

func New(chat chatgrpc.Chat, port string) *App {
	server := grpc.NewServer()
	chatgrpc.Register(server, chat)
	return &App{grpcServer: server, port: port}
}
func (a *App) Run() error {
	lis, err := net.Listen("tcp", a.port)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := a.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
