package model

type File struct {
	ID   uint32 `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255);not null"`
	Path string `gorm:"type:varchar(255);not null"`
	Size int64  `gorm:"type:int unsigned;not null"`
}
