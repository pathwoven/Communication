package model

type CodeTag struct {
	ID     uint32 `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	UserID uint32 `gorm:"type:int unsigned;not null;index:user_id_idx"`
	Name   string `gorm:"type:varchar(50);not null"`
}
