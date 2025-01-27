package model

import "time"

type Chat struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null;index:user_target_id_idx,priority:1"`
	TargetID    uint      `gorm:"not null;index:user_target_id_idx,priority:2;comment:对方的id"`
	UnreadCount uint8     `gorm:"not null;default:0"`
	LastMessage string    `gorm:"type:varchar(255);default:null"`
	LastTime    time.Time `gorm:"default:null"`
	IsPinned    bool      `gorm:"not null;default:false"`
	IsMute      bool      `gorm:"not null;default:false"`
	IsBlocked   bool      `gorm:"not null;default:false"`
	IsDeleted   bool      `gorm:"not null;default:false"`
	Tag1ID      uint      `gorm:"default:null"`
	Tag2ID      uint      `gorm:"default:null"`
	Tag3ID      uint      `gorm:"default:null"`
}
