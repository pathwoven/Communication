package model

type MessageDeleted struct {
	UserID         int `gorm:"primaryKey;not null"`
	StartMessageID int `gorm:"primaryKey;not null"`
	EndMessageID   int `gorm:"primaryKey;not null"`
}
