package model

type Favorite struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	UserID   int `gorm:"not null;index:user_id_idx"`
	Type     int `gorm:"not null;comment:0:消息 1:多选消息 2:笔记 3:代码"`
	TargetID int `gorm:"not null;comment:在对应表中的id"`
}
