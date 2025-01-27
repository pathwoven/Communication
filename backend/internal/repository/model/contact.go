package model

type Contact struct {
	UserID        int    `gorm:"type:int;primaryKey;not null"`
	TargetID      int    `gorm:"type:int;primaryKey;not null"`
	Remark        string `gorm:"type:varchar(255);default:null"`
	DivideID      int    `gorm:"type:int;not null"`
	IsBlacklisted bool   `gorm:"type:tinyint(1);not null;default:false"`
	BackgroundImg string `gorm:"type:varchar(255);default:null"`
}
