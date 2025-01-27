package model

import "time"

type Note struct {
	ID         uint32    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	UserID     uint32    `gorm:"type:int unsigned;not null;index:user_id_idx"`
	Name       string    `gorm:"type:varchar(255);not null"`
	SavePath   string    `gorm:"type:varchar(255);not null"`
	Details    string    `gorm:"type:varchar(255)"`
	CreateTime time.Time `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
	IsDeleted  bool      `gorm:"type:tinyint(1);not null;default:false"`
	Tag1ID     uint32    `gorm:"type:int unsigned;default:null"`
	Tag2ID     uint32    `gorm:"type:int unsigned;default:null"`
	Tag3ID     uint32    `gorm:"type:int unsigned;default:null"`
}
