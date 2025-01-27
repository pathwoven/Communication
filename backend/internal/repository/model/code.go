package model

import (
	"time"
)

type Code struct {
	ID         uint32    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	UserID     uint32    `gorm:"type:int unsigned;not null;index:user_id_name_idx,priority:1"`
	Name       string    `gorm:"type:varchar(50);not null;index:user_id_name_idx,priority:2"`
	SavePath   string    `gorm:"type:varchar(255);not null"`
	Details    string    `gorm:"type:varchar(255);default:null"`
	CreateTime time.Time `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
	IsDeleted  bool      `gorm:"type:tinyint(1);not null;default:0"`
	Tag1ID     uint32    `gorm:"type:int unsigned;default:null"`
	Tag2ID     uint32    `gorm:"type:int unsigned;default:null"`
	Tag3ID     uint32    `gorm:"type:int unsigned;default:null"`
}
