package model

type NoteTag struct {
	ID     int    `gorm:"type:int;primaryKey;autoIncrement"`
	UserID int    `gorm:"type:int;not null;index:user_id_idx"`
	Name   string `gorm:"type:varchar(30);not null"`
}
