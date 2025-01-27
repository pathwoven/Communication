package model

import "time"

type Schedule struct {
	ID         uint32    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	UserID     uint32    `gorm:"type:int unsigned;not null;index:user_id_idx"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Details    string    `gorm:"type:varchar(255)"`
	StartTime  time.Time `gorm:"not null"`
	EndTime    time.Time `gorm:"default:null"`
	NeedRemind bool      `gorm:"type:tinyint(1);not null;default:false"`
	Color      string    `gorm:"type:varchar(50);not null"`
}
