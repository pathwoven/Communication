package model

type Entity struct {
	ID           int    `gorm:"primaryKey"`
	IsGroup      bool   `gorm:"not null"`
	DisplayID    string `gorm:"type:varchar(15);not null;comment:展示给用户的id（账号或群号）"`
	Name         string `gorm:"type:varchar(50);not null"`
	Avatar       string `gorm:"type:varchar(255);not null"`
	OnlineStatus int    `gorm:"not null;comment:0为离线，1为在线，2为隐身"`
	Status       string `gorm:"type:varchar(20);default:null"`
	Introduction string `gorm:"type:varchar(255);default:null"`
	GroupOwnerID int    `gorm:"default:null"`
}
