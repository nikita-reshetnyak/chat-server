// main is the entry point of the application.
// It initializes and runs the program.
package main

import (
	"flag"
	"log"

	"github.com/nikita-reshetnyak/chat-server/internal/app"
	"github.com/nikita-reshetnyak/chat-server/internal/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}
func main() {
	flag.Parse()
	grpcConfig, err := config.NewGrpcConfig()
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
	application := app.New(grpcConfig.Address(), configPath)
	err = application.GrpcServer.Run()
	if err != nil {
		log.Fatalf("failed to run grpc server")
	}
}
