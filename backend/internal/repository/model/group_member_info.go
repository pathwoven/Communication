package model

type GroupMemberInfo struct {
	GroupID       int    `gorm:"primaryKey;not null"`
	UserID        int    `gorm:"primaryKey;not null"`
	GroupNickname string `gorm:"type:varchar(30);default:null"`
	IsBanned      bool   `gorm:"not null;default:false"`
	IsManager     bool   `gorm:"not null"`
}
