package app

import (
	"log"

	grpcapp "github.com/nikita-reshetnyak/chat-server/internal/app/grpc"
	chatserver "github.com/nikita-reshetnyak/chat-server/internal/services/chat-server"
	"github.com/nikita-reshetnyak/chat-server/internal/storage/postgres"
)

type App struct {
	GrpcServer *grpcapp.App
}

func New(grpcPort string, configPath string) *App {
	conn, err := postgres.New(configPath)
	if err != nil {
		log.Fatalf("failed to connect to postgres:%v", err)
	}
	chatService := chatserver.New(conn)
	app := grpcapp.New(chatService, grpcPort)
	return &App{GrpcServer: app}
}
