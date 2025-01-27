package model

type DivideGroup struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	GroupID int    `gorm:"not null;index:group_id_idx"`
	Name    string `gorm:"type:varchar(30);not null"`
}
