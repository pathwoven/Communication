package model

import "time"

type GroupNotice struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	GroupID     int       `gorm:"not null;index:group_id_idx"`
	SenderID    int       `gorm:"not null"`
	Content     string    `gorm:"type:text;not null"`
	ContentType int       `gorm:"not null"`
	CreateTime  time.Time `gorm:"not null"`
}
