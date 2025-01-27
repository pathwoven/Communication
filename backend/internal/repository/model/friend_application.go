package model

import "time"

type FriendApplication struct {
	ID         uint32    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	SenderID   uint32    `gorm:"type:int unsigned;not null;index:sender_id_idx"`
	ReceiverID uint32    `gorm:"type:int unsigned;not null;index:receiver_id_idx"`
	Reason     string    `gorm:"type:varchar(255);default:null"`
	Time       time.Time `gorm:"not null"`
	Status     uint8     `gorm:"type:tinyint;not null;comment:0:未处理 1:已同意 2:已拒绝 3:已忽略"`
}
