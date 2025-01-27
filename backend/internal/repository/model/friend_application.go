package model

import "time"

type FriendApplication struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	SenderID   int       `gorm:"not null;index:sender_id_idx"`
	ReceiverID int       `gorm:"not null;index:receiver_id_idx"`
	Reason     string    `gorm:"type:varchar(255);default:null"`
	Time       time.Time `gorm:"not null"`
	Status     int       `gorm:"not null;comment:0:未处理 1:已同意 2:已拒绝 3:已忽略"`
}
