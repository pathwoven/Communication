package model

import "time"

type Schedule struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	UserID     int       `gorm:"not null;index:user_id_idx"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Details    string    `gorm:"type:text"`
	StartTime  time.Time `gorm:"not null"`
	EndTime    time.Time `gorm:"default:null"`
	NeedRemind bool      `gorm:"not null;default:false"`
	Color      string    `gorm:"type:varchar(50);not null"`
}
