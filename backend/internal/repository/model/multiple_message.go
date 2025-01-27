package model

type MultipleMessage struct {
	ID        int `gorm:"type:int;primaryKey;autoIncrement"`
	MessageID int `gorm:"type:int;primaryKey;not null"`
}
