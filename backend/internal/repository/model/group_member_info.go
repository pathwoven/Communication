package model

type GroupMemberInfo struct {
	GroupID       int    `gorm:"type:int;primaryKey;not null"`
	UserID        int    `gorm:"type:int;primaryKey;not null"`
	GroupNickname string `gorm:"type:varchar(30);default:null"`
	IsBanned      bool   `gorm:"type:tinyint(1);not null;default:false"`
	IsManager     bool   `gorm:"type:tinyint(1);not null"`
}
