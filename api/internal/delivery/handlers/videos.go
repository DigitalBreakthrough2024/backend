package handlers

import (
	"backend/internal/models/dto"
	"backend/pkg/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type VideoHandler struct {
	session database.Session
}

func InitVideoHandler(session database.Session) VideoHandler {
	return VideoHandler{
		session: session,
	}
}

// GetVideosForUser
// @Summary Get Videos for User
// @Description Returns a list of 10 mock videos for the user
// @Tags Videos
// @Accept json
// @Produce json
// @Success 200 {array} dto.Video "List of videos"
// @Router /api/videos [get]
func (v VideoHandler) GetVideosForUser(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil || sessionID == "" {
		sessionID = uuid.New().String()
		c.SetCookie("session_id", sessionID, 3600, "/", "", false, true)
	}

	// Замоканные данные о видео
	videos := []dto.Video{
		{ID: "1", Name: "Video 1", Category: "Category 1", Date: parseDate("2024-01-01"), Duration: 120},
		{ID: "2", Name: "Video 2", Category: "Category 2", Date: parseDate("2024-01-02"), Duration: 150},
		{ID: "3", Name: "Video 3", Category: "Category 3", Date: parseDate("2024-01-03"), Duration: 180},
		{ID: "4", Name: "Video 4", Category: "Category 4", Date: parseDate("2024-01-04"), Duration: 210},
		{ID: "5", Name: "Video 5", Category: "Category 5", Date: parseDate("2024-01-05"), Duration: 240},
		{ID: "6", Name: "Video 6", Category: "Category 6", Date: parseDate("2024-01-06"), Duration: 270},
		{ID: "7", Name: "Video 7", Category: "Category 7", Date: parseDate("2024-01-07"), Duration: 300},
		{ID: "8", Name: "Video 8", Category: "Category 8", Date: parseDate("2024-01-08"), Duration: 330},
		{ID: "9", Name: "Video 9", Category: "Category 9", Date: parseDate("2024-01-09"), Duration: 360},
		{ID: "10", Name: "Video 10", Category: "Category 10", Date: parseDate("2024-01-10"), Duration: 390},
	}

	err = v.session.Set(c.Request.Context(), sessionID, database.SessionData{
		Vector:  make([]float32, 624),
		Watched: make([]string, 0),
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, videos)
}

// CreateReview
// @Summary Create Review
// @Description Returns a new list of 10 recommended videos based on user reviews
// @Tags Videos
// @Accept json
// @Produce json
// @Param reviews body []dto.VideoReview true "List of video reviews"
// @Success 200 {array} dto.Video "List of recommended videos"
// @Failure 400 {object} gin.H "Invalid request"
// @Router /api/videos [post]
func (v VideoHandler) CreateReview(c *gin.Context) {
	var reviews []dto.VideoReview

	if err := c.ShouldBindJSON(&reviews); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	fmt.Println(reviews)

	recommendedVideos := []dto.Video{
		{ID: "11", Name: "Video 11", Category: "Category 11", Date: parseDate("2024-01-11"), Duration: 400},
		{ID: "12", Name: "Video 12", Category: "Category 12", Date: parseDate("2024-01-12"), Duration: 420},
		{ID: "13", Name: "Video 13", Category: "Category 13", Date: parseDate("2024-01-13"), Duration: 440},
		{ID: "14", Name: "Video 14", Category: "Category 14", Date: parseDate("2024-01-14"), Duration: 460},
		{ID: "15", Name: "Video 15", Category: "Category 15", Date: parseDate("2024-01-15"), Duration: 480},
		{ID: "16", Name: "Video 16", Category: "Category 16", Date: parseDate("2024-01-16"), Duration: 500},
		{ID: "17", Name: "Video 17", Category: "Category 17", Date: parseDate("2024-01-17"), Duration: 520},
		{ID: "18", Name: "Video 18", Category: "Category 18", Date: parseDate("2024-01-18"), Duration: 540},
		{ID: "19", Name: "Video 19", Category: "Category 19", Date: parseDate("2024-01-19"), Duration: 560},
		{ID: "20", Name: "Video 20", Category: "Category 20", Date: parseDate("2024-01-20"), Duration: 580},
	}

	c.JSON(http.StatusOK, recommendedVideos)
}

func parseDate(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}
