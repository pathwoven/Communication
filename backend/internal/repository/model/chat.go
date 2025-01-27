package model

import "time"

type Chat struct {
	ID          uint      `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	UserID      uint      `gorm:"type:int unsigned;not null;index:user_target_id_idx,priority:1"`
	TargetID    uint      `gorm:"type:int unsigned;not null;index:user_target_id_idx,priority:2;comment:对方的id"`
	UnreadCount uint8     `gorm:"type:tinyint;not null;default:0"`
	LastMessage string    `gorm:"type:varchar(255);default:null"`
	LastTime    time.Time `gorm:"default:null"`
	IsPinned    bool      `gorm:"type:tinyint(1);not null;default:false"`
	IsMute      bool      `gorm:"type:tinyint(1);not null;default:false"`
	IsBlocked   bool      `gorm:"type:tinyint(1);not null;default:false"`
	IsDeleted   bool      `gorm:"type:tinyint(1);not null;default:false"`
	Tag1ID      uint      `gorm:"type:int unsigned;default:null"`
	Tag2ID      uint      `gorm:"type:int unsigned;default:null"`
	Tag3ID      uint      `gorm:"type:int unsigned;default:null"`
}
