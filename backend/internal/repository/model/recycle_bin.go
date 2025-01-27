package model

type RecycleBin struct {
	UserID   uint32 `gorm:"type:int unsigned;primaryKey;not null"`
	Type     uint8  `gorm:"type:tinyint;primaryKey;not null;comment:0：笔记，1：代码"`
	TargetID uint32 `gorm:"type:int unsigned;primaryKey;not null;comment:在对应表中的id"`
}
