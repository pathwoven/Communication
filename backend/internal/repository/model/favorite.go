package model

type Favorite struct {
	ID       uint32 `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	UserID   uint32 `gorm:"type:int unsigned;not null;index:user_id_idx"`
	Type     uint8  `gorm:"type:tinyint;not null;comment:0:消息 1:多选消息 2:笔记 3:代码"`
	TargetID uint32 `gorm:"type:int unsigned;not null;comment:在对应表中的id"`
}
