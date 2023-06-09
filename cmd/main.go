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

	r := gin.New()

	// call logger
	r.Use(gin.Recovery(), gin.Logger())

	api.NewApi(r, &cfg, pgStore, log)

	err = r.Run(cfg.ServerHost + cfg.ServerPort)
	if err != nil {
		log.Panic("Error listening server: ", logger.Error(err))
		return
	}

	//booking, err := pgStore.Booking().BookRoom(context.Background(), 2, models.BookingRequest{
	//	Resident: models.Resident{Name: "Abdulloh"},
	//	Start:    time.Now(),
	//	End:      time.Now().Add(time.Hour * 24),
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(booking)

}
