package api

import (
	_ "Impact/api/docs"
	"Impact/api/handlers"
	"Impact/config"
	"Impact/pkg/logger"
	"Impact/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handlers.NewHandler(cfg, store, logger)
	// rooms api
	r.GET("/api/rooms/:id", handler.GetByIDRoom)
	r.GET("/api/rooms", handler.GetRoomsList)

	// booking api
	r.POST("/api/rooms/:id/book", handler.BookingRoom)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
