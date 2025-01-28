package model

type DivideGroup struct {
	ID     uint32 `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	UserID uint32 `gorm:"type:int unsigned;not null;index:user_id_idx"`
	Name   string `gorm:"type:varchar(30);not null"`
}
