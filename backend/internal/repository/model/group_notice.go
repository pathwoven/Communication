package model

import "time"

type GroupNotice struct {
	ID          int       `gorm:"type:int;primaryKey;autoIncrement"`
	GroupID     int       `gorm:"type:int;not null;index:group_id_idx"`
	SenderID    int       `gorm:"type:int;not null"`
	Content     string    `gorm:"type:varchar(255);not null"`
	ContentType int       `gorm:"type:tinyint;not null"`
	CreateTime  time.Time `gorm:"not null"`
}
