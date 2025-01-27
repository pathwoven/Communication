package model

type MultipleMessage struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	MessageID int `gorm:"primaryKey;not null"`
}
