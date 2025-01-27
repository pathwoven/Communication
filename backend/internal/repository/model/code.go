package model

import "time"

type Code struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	UserID          int       `gorm:"not null"`
	Name            string    `gorm:"type:varchar(50);not null"`
	SavePath        string    `gorm:"type:varchar(255);not null"`
	Details         string    `gorm:"type:text"`
	CreateTime      time.Time `gorm:"not null"`
	UpdateTime      time.Time `gorm:"not null"`
	IsDeleted       bool      `gorm:"not null;default:false"`
	Tag1ID          int       `gorm:"default:null"`
	Tag2ID          int       `gorm:"default:null"`
	Tag3ID          int       `gorm:"default:null"`
	UserIDNameIndex struct{}  `gorm:"index:user_id_name_idx,unique"`
}
