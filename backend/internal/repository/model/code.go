package model

import (
	"time"
)

type Code struct {
	ID         int       `gorm:"type:int;primaryKey;autoIncrement"`
	UserID     int       `gorm:"type:int;not null;index:user_id_name_idx,priority:1"`
	Name       string    `gorm:"type:varchar(50);not null;index:user_id_name_idx,priority:2"`
	SavePath   string    `gorm:"type:varchar(255);not null"`
	Details    string    `gorm:"type:varchar(255);default:null"`
	CreateTime time.Time `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
	IsDeleted  bool      `gorm:"type:tinyint(1);not null;default:0"`
	Tag1ID     int       `gorm:"type:int;default:null"`
	Tag2ID     int       `gorm:"type:int;default:null"`
	Tag3ID     int       `gorm:"type:int;default:null"`
}
