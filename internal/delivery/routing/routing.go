package routing

import (
	"backend/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouting(engine *gin.Engine) {
	videoHandler := handlers.InitVideoHandler()

	baseGroup := engine.Group("/api")
	videoGroup := baseGroup.Group("/videos")

	videoGroup.GET("", videoHandler.GetVideosForUser)
	videoGroup.POST("", videoHandler.CreateReview)
}
