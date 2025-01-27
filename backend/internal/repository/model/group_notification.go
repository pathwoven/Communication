package model

import "time"

type GroupNotification struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	GroupID    int       `gorm:"not null"`
	SenderID   int       `gorm:"not null;index:sender_id_idx"`
	ReceiverID int       `gorm:"not null;index:receiver_id_idx"`
	Type       int       `gorm:"not null;comment:0:申请 1:邀请 2:邀请后的申请 3:通知"`
	Reason     string    `gorm:"type:varchar(255);default:null"`
	Time       time.Time `gorm:"not null"`
	Status     int       `gorm:"not null;comment:0:未处理 1:已同意 2:已拒绝 3:已忽略 4:已撤销"`
	PendStatus int       `gorm:"default:null;comment:（用于邀请后的申请）0:未处理 1:已同意 2:已拒绝 3:已忽略"`
}
