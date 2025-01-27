package model

import "time"

type Message struct {
	ID          int       `gorm:"type:int;primaryKey;autoIncrement"`
	SenderID    int       `gorm:"type:int;not null;index:sender_id_idx"`
	ReceiverID  int       `gorm:"type:int;not null;index:receiver_id_idx"`
	Content     string    `gorm:"type:varchar(255);not null"`
	ContentType int       `gorm:"type:tinyint;not null;comment:0表示文本，1表示图片，2表示文件，3表示群投票，4表示多选消息"`
	CreateTime  time.Time `gorm:"not null"`
	IsRecall    bool      `gorm:"type:tinyint(1);not null;default:false"`
	Operation   int       `gorm:"type:tinyint;default:null"`
	OpMessageID int       `gorm:"type:int;default:null"`
}
