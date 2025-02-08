package model

import "time"

type Chat struct {
	ID          uint32     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	UserID      uint32     `gorm:"type:int unsigned;not null;index:user_target_id_idx,priority:1"`
	TargetID    uint32     `gorm:"type:int unsigned;not null;index:user_target_id_idx,priority:2;comment:对方的id"`
	UnreadCount uint8      `gorm:"type:tinyint;not null;default:0"`
	LastMessage string     `gorm:"type:varchar(255);default:null"`
	LastTime    *time.Time `gorm:"default:null"`
	LastPerson  string     `gorm:"type:varchar(20);default:null"`
	IsPinned    bool       `gorm:"type:tinyint(1);not null;default:false"`
	IsMuted     bool       `gorm:"type:tinyint(1);not null;default:false"`
	IsBlocked   bool       `gorm:"type:tinyint(1);not null;default:false"`
	IsDeleted   bool       `gorm:"type:tinyint(1);not null;default:false"`
	Tag1        string     `gorm:"type:varchar(20);default:null"`
	Tag2        string     `gorm:"type:varchar(20);default:null"`
	Tag3        string     `gorm:"type:varchar(20);default:null"`
}

func (Chat) TableComment() string {
	return "聊天列表"
}
