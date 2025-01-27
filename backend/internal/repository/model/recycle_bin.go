package model

type RecycleBin struct {
	UserID   int `gorm:"primaryKey;not null"`
	Type     int `gorm:"primaryKey;not null;comment:0：笔记，1：代码"`
	TargetID int `gorm:"primaryKey;not null;comment:在对应表中的id"`
}
