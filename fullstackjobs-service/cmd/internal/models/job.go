package models

import (
	"time"
)

// Job represents a job listing
type Job struct {
	ID          string    `gorm:"primaryKey" json:"id"` // UUID
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	LogoURL     string    `json:"logoUrl"`
	PostedAt    time.Time `json:"postedAt"`
	CreatedAt   time.Time `json:"createdAt"`
}
