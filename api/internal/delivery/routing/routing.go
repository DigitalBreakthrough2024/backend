package routing

import (
	"backend/internal/delivery/handlers"
	"backend/pkg/database"
	"github.com/gin-gonic/gin"
)

func InitRouting(engine *gin.Engine, session database.Session) {
	videoHandler := handlers.InitVideoHandler(session)

	baseGroup := engine.Group("/api")
	videoGroup := baseGroup.Group("/videos")

	videoGroup.GET("", videoHandler.GetVideosForUser)
	videoGroup.POST("", videoHandler.CreateReview)
}
