package main

import (
	"context"
	"users_service/configs"
	"users_service/grpc"
	"users_service/grpc/clients"
	"users_service/pkg/logger"
	"users_service/service"
	"users_service/storage"

	"fmt"
	"net"
)

func main() {
	cfg := configs.Load()

	log := logger.NewLogger(cfg.ServiceName, cfg.LoggerLevel, cfg.LogPath)
	defer logger.Cleanup(log)

	storage, err := storage.New(context.Background(), cfg, &log)
	if err != nil {
		log.Panic("error while creating storage in main", logger.Error(err))
		return
	}
	defer storage.Close()

	clients, err := clients.NewClients(cfg)
	if err != nil {
		log.Panic("error while creating clients in main", logger.Error(err))
		return
	}
	services := service.NewServiceManager(storage, clients, log)

	server := grpc.SetUpServer(services, log)
	listener, err := net.Listen("tcp",
		cfg.UserServiceGrpcHost+cfg.UserServiceGrpcPort,
	)
	if err != nil {
		log.Panic("error while creating listener for user service", logger.Error(err))
		return
	}
	defer listener.Close()

	fmt.Printf("User service is listening on port %s...\n",
		cfg.UserServiceGrpcHost+cfg.UserServiceGrpcPort)
	if err := server.Serve(listener); err != nil {
		log.Fatal("Error with listening user server", logger.Error(err))
	}
}
