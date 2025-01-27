package model

type MultipleMessage struct {
	ID        uint32 `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	MessageID uint32 `gorm:"type:int unsigned;primaryKey;not null"`
}

func (MultipleMessage) TableComment() string {
	return "记录多选的消息"
}
