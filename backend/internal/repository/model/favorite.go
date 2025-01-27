package model

type Favorite struct {
	ID       int `gorm:"type:int;primaryKey;autoIncrement"`
	UserID   int `gorm:"type:int;not null;index:user_id_idx"`
	Type     int `gorm:"type:tinyint;not null;comment:0:消息 1:多选消息 2:笔记 3:代码"`
	TargetID int `gorm:"type:int;not null;comment:在对应表中的id"`
}
