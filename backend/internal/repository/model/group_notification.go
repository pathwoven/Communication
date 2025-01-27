package model

import "time"

type GroupNotification struct {
	ID         uint32    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	GroupID    uint32    `gorm:"type:int unsigned;not null"`
	SenderID   uint32    `gorm:"type:int unsigned;not null;index:sender_id_idx"`
	ReceiverID uint32    `gorm:"type:int unsigned;not null;index:receiver_id_idx"`
	Type       uint8     `gorm:"type:tinyint;not null;comment:0:申请 1:邀请 2:邀请后的申请 3:通知"`
	Reason     string    `gorm:"type:varchar(255);default:null"`
	Time       time.Time `gorm:"not null"`
	Status     uint8     `gorm:"type:tinyint;not null;comment:0:未处理 1:已同意 2:已拒绝 3:已忽略 4:已撤销"`
	PendStatus uint8     `gorm:"type:tinyint;default:null;comment:（用于邀请后的申请）0:未处理 1:已同意 2:已拒绝 3:已忽略"`
}
