package model

import "time"

type FriendApplication struct {
	ID         int       `gorm:"type:int;primaryKey;autoIncrement"`
	SenderID   int       `gorm:"type:int;not null;index:sender_id_idx"`
	ReceiverID int       `gorm:"type:int;not null;index:receiver_id_idx"`
	Reason     string    `gorm:"type:varchar(255);default:null"`
	Time       time.Time `gorm:"not null"`
	Status     int       `gorm:"type:tinyint;not null;comment:0:未处理 1:已同意 2:已拒绝 3:已忽略"`
}
