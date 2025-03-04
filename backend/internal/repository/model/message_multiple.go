package model

type MessageMultiple struct {
	ID        uint32 `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	MessageID uint32 `gorm:"type:int unsigned;primaryKey;not null"`
}

func (MessageMultiple) TableComment() string {
	return "记录多选的消息"
}
