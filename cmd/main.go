package main

import (
	"Impact/api"
	"Impact/config"
	"Impact/pkg/logger"
	"Impact/storage/postgres"
	"context"
	"github.com/gin-gonic/gin"
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

	r := gin.New()

	// call logger
	r.Use(gin.Recovery(), gin.Logger())

	api.NewApi(r, &cfg, pgStore, log)

	err = r.Run(cfg.ServerHost + cfg.ServerPort)
	if err != nil {
		log.Panic("Error listening server: ", logger.Error(err))
		return
	}

	//f := time.Date(2023, 7, 8, 17, 0, 0, 0, time.UTC)
	//t := time.Date(2023, 9, 9, 17, 54, 0, 0, time.UTC)
	//a, err := pgStore.Booking().Check(context.Background(), f, t)
	//if err != nil {
	//	log.Fatal("Error listening server: ", logger.Error(err))
	//	return
	//}
	//fmt.Println(a)

	//booking, err := pgStore.Booking().BookRoom(context.Background(), 2, models.BookingRequest{
	//	Resident: models.Resident{Name: "Abdulloh"},
	//	Start:    f,
	//	End:      t,
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(booking)

}
