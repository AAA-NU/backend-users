package models

type User struct {
	TelegramID string `json:"tgID" gorm:"primaryKey"`
	Role       string `json:"role" gorm:"default:student"`
	Language   string `json:"language" gorm:"default:ru"`
}
