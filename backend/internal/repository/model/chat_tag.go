package model

type ChatTag struct {
	UserID uint32 `gorm:"type:int;primaryKey;not null"`
	Name   string `gorm:"primaryKey;type:varchar(30);not null"`
}
