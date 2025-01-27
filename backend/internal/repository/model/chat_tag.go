package model

type ChatTag struct {
	UserID int    `gorm:"type:int;primaryKey;not null"`
	Name   string `gorm:"primaryKey;type:varchar(30);not null"`
}
