package model

type DivideGroup struct {
	ID      int    `gorm:"type:int;primaryKey;autoIncrement"`
	GroupID int    `gorm:"type:int;not null;index:group_id_idx"`
	Name    string `gorm:"type:varchar(30);not null"`
}
