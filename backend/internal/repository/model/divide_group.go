package model

type DivideGroup struct {
	ID      uint32 `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	GroupID uint32 `gorm:"type:int unsigned;not null;index:group_id_idx"`
	Name    string `gorm:"type:varchar(30);not null"`
}
