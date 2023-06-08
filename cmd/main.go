package main

import (
	"Impact/config"
	"Impact/models"
	"Impact/pkg/logger"
	"Impact/storage/postgres"
	"context"
	"fmt"
)

func main() {
	cfg := config.Load()

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	//room, err := pgStore.Room().GetRoom(context.Background(), 3)
	//if err != nil {
	//	logger.Error(err)
	//}
	//fmt.Println(room)

	rooms, err := pgStore.Room().GetRooms(context.Background(), models.RoomsRequest{
		Search:   "r",
		Type:     "",
		Page:     1,
		PageSize: 10,
	})

	//SELECT * FROM room
	//WHERE name LIKE ''
	//AND type LIKE ''
	//ORDER BY id
	//OFFSET 0
	//LIMIT 10

	if err != nil {
		logger.Error(err)
	}

	fmt.Println(rooms)

}
