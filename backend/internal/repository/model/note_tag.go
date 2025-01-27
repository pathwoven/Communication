package model

type NoteTag struct {
	ID     int    `gorm:"primaryKey;autoIncrement"`
	UserID int    `gorm:"not null;index:user_id_idx"`
	Name   string `gorm:"type:varchar(30);not null"`
}
