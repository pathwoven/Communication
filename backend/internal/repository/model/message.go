package model

import "time"

type Message struct {
	ID          uint32    `gorm:"type:int unsigned;primaryKey;autoIncrement" json:"id"`
	SenderID    uint32    `gorm:"type:int unsigned;not null;index:sender_id_idx" json:"sender_id"`
	ReceiverID  uint32    `gorm:"type:int unsigned;not null;index:receiver_id_idx" json:"receiver_id"`
	Content     string    `gorm:"type:varchar(255);not null" json:"content"`
	ContentType uint8     `gorm:"type:tinyint;not null;comment:0表示文本，1表示图片，2表示文件，3表示群投票，4表示多选消息" json:"content_type"`
	CreateTime  time.Time `gorm:"not null" json:"create_time"`
	Operation   uint8     `gorm:"type:tinyint;default:null;comment:0撤回" json:"operation"`
	OpMessageID uint32    `gorm:"type:int unsigned;default:null" json:"op_message_id"`
}
