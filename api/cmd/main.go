package main

import (
	"backend/config"
	"backend/internal/delivery/docs"
	"backend/internal/delivery/routing"
	"backend/pkg/database"
	"backend/pkg/log"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	logger, loggerFile := log.InitLoggers()
	defer loggerFile.Close()
	logger.Info().Msg("Logger Initialized")

	router.Static("/static", "../static")
	logger.Info().Msg("Static Initialized")

	config.InitConfig()
	logger.Info().Msg("Config Initialized")

	_ = database.GetDB()
	logger.Info().Msg("Database Initialized")

	session := database.InitRedisSession()
	logger.Info().Msg("Session Initialized")

	routing.InitRouting(router, session)
	logger.Info().Msg("Routing Initialized")

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Info().Msg("Swagger Initialized")

	if err := router.Run("0.0.0.0:80"); err != nil {
		panic(fmt.Sprintf("Failed to run client: %s", err.Error()))
	}
}
