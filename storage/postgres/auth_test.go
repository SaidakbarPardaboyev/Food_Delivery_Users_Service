package postgres

import (
	"context"
	"users_service/configs"
	"users_service/pkg/logger"
)

func setupRepo() (*authRepo, error){
	cfg := configs.Load()
	db, err := ConnectDB(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	logger := logger.NewLogger(cfg.ServiceName, "debug", cfg.LogPath)

	return &authRepo{
		db:  db,
		log: logger,
	}, nil
}

