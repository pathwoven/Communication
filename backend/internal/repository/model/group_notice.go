package model

import "time"

type GroupNotice struct {
	ID          uint32    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	GroupID     uint32    `gorm:"type:int unsigned;not null;index:group_id_idx"`
	SenderID    uint32    `gorm:"type:int unsigned;not null"`
	Content     string    `gorm:"type:varchar(255);not null"`
	ContentType uint8     `gorm:"type:tinyint;not null"`
	CreateTime  time.Time `gorm:"not null"`
}

func (GroupNotice) TableComment() string {
	return "群公告"
}
