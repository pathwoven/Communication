package model

type RecycleBin struct {
	UserID   int `gorm:"type:int;primaryKey;not null"`
	Type     int `gorm:"type:tinyint(1);primaryKey;not null;comment:0：笔记，1：代码"`
	TargetID int `gorm:"type:int;primaryKey;not null;comment:在对应表中的id"`
}
