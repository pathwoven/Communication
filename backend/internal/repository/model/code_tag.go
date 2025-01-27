package model

type CodeTag struct {
	ID     int    `gorm:"primaryKey;autoIncrement"`
	UserID int    `gorm:"not null;index:user_id_idx"`
	Name   string `gorm:"type:varchar(50);not null"`
}
