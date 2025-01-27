package model

type Contact struct {
	UserID        uint32 `gorm:"type:int unsigned;primaryKey;not null"`
	TargetID      uint32 `gorm:"type:int unsigned;primaryKey;not null"`
	Remark        string `gorm:"type:varchar(255);default:null"`
	DivideID      uint32 `gorm:"type:int unsigned;not null"`
	IsBlacklisted bool   `gorm:"type:tinyint(1);not null;default:false"`
	BackgroundImg string `gorm:"type:varchar(255);default:null"`
}

func (Contact) TableComment() string {
	return "记录好友与群组的关系"
}
