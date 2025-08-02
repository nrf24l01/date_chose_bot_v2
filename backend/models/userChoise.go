package models

import (
	"github.com/lib/pq"
)

type UserChoice struct {
	BaseModel
	UserID        int64         `json:"user_id" gorm:"not null;uniqueIndex"`
	UserName      string        `json:"user_name" gorm:"default:''"`
	SelectedDates pq.StringArray `json:"selected_dates" gorm:"type:text[];not null"`
}