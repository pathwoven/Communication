package model

import "time"

type Note struct {
	ID         int       `gorm:"type:int;primaryKey;autoIncrement"`
	UserID     int       `gorm:"type:int;not null;index:user_id_idx"`
	Name       string    `gorm:"type:varchar(255);not null"`
	SavePath   string    `gorm:"type:varchar(255);not null"`
	Details    string    `gorm:"type:varchar(255)"`
	CreateTime time.Time `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
	IsDeleted  bool      `gorm:"type:tinyint(1);not null;default:false"`
	Tag1ID     int       `gorm:"type:int;default:null"`
	Tag2ID     int       `gorm:"type:int;default:null"`
	Tag3ID     int       `gorm:"type:int;default:null"`
}
