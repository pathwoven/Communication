package model

type MessageDeleted struct {
	UserID         uint32 `gorm:"type:int unsigned;primaryKey;not null"`
	StartMessageID uint32 `gorm:"type:int unsigned;primaryKey;not null"`
	EndMessageID   uint32 `gorm:"type:int unsigned;primaryKey;not null"`
}
