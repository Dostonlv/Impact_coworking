package main

import (
	"Impact/config"
	"Impact/pkg/logger"
	"Impact/storage/postgres"
	"context"
	"fmt"
)

func main() {
	cfg := config.Load()

	var loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func(l logger.LoggerI) {
		err := logger.Cleanup(l)
		if err != nil {
			logger.Any("logger.Cleanup", logger.Error(err))
		}
	}(log)

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	//r := gin.New()
	//
	////call logger
	//r.Use(gin.Recovery(), gin.Logger())
	//
	//api.NewApi(r, &cfg, pgStore, log)
	//
	//err = r.Run(cfg.ServerHost + cfg.ServerPort)
	//if err != nil {
	//	log.Panic("Error listening server: ", logger.Error(err))
	//	return
	//}

	arr, err := pgStore.Booking().GetBookingRooms(context.Background(), 1)
	if err != nil {
		log.Fatal("Error listening server: ", logger.Error(err))
		return
	}
	fmt.Println(arr)

}
