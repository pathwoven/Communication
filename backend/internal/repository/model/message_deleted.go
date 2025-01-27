package model

type MessageDeleted struct {
	UserID         int `gorm:"type:int;primaryKey;not null"`
	StartMessageID int `gorm:"type:int;primaryKey;not null"`
	EndMessageID   int `gorm:"type:int;primaryKey;not null"`
}
