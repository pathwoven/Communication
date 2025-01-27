package model

import "time"

type Message struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	SenderID    int       `gorm:"not null;index:sender_id_idx"`
	ReceiverID  int       `gorm:"not null;index:receiver_id_idx"`
	Content     string    `gorm:"type:varchar(255);not null"`
	ContentType int       `gorm:"not null;comment:0表示文本，1表示图片，2表示文件，3表示群投票，4表示多选消息"`
	CreateTime  time.Time `gorm:"not null"`
	IsRecall    bool      `gorm:"not null;default:false"`
	Operation   int       `gorm:"default:null"`
	OpMessageID int       `gorm:"default:null"`
}
