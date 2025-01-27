package model

type Contact struct {
	UserID        int    `gorm:"primaryKey;not null"`
	TargetID      int    `gorm:"primaryKey;not null"`
	Remark        string `gorm:"type:varchar(255);default:null"`
	DivideID      int    `gorm:"not null"`
	IsBlacklisted bool   `gorm:"not null;default:false"`
	BackgroundImg string `gorm:"type:varchar(255);default:null"`
}
