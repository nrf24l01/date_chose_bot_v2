package models

type UserChoice struct {
	BaseModel
	UserID        int      `json:"user_id" gorm:"not null;uniqueIndex"`
	SelectedDates []string `json:"selected_dates" gorm:"type:text[];not null"`
}