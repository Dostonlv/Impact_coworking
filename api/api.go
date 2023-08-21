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

// @title Swagger Example API
// NewApi godoc
// @description This is a api gateway
// @version 1.0
func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handlers.NewHandler(cfg, store, logger)
	r.Use(customCORSMiddleware())

	// rooms api
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	r.GET("/api/rooms/:id", handler.GetByIDRoom)
	r.GET("/api/rooms", handler.GetRoomsList)

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	// booking api
	r.POST("/api/rooms/:id/book", handler.BookingRoom)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
