package dto

import "time"

type Video struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
	Duration int       `json:"duration"`
}

type VideoReview struct {
	ID         string `json:"id"`
	Evaluation int    `json:"evaluation"`
}
