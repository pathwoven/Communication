package model

import "time"

type Message struct {
	ID          uint32    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	SenderID    uint32    `gorm:"type:int unsigned;not null;index:sender_id_idx"`
	ReceiverID  uint32    `gorm:"type:int unsigned;not null;index:receiver_id_idx"`
	Content     string    `gorm:"type:varchar(255);not null"`
	ContentType uint8     `gorm:"type:tinyint;not null;comment:0表示文本，1表示图片，2表示文件，3表示群投票，4表示多选消息"`
	CreateTime  time.Time `gorm:"not null"`
	IsRecall    bool      `gorm:"type:tinyint(1);not null;default:false"`
	Operation   uint8     `gorm:"type:tinyint;default:null"`
	OpMessageID uint32    `gorm:"type:int unsigned;default:null"`
}
