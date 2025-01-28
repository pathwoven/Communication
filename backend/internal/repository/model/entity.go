package model

type Entity struct {
	ID           uint32 `gorm:"type:int unsigned;primaryKey"`
	IsGroup      bool   `gorm:"type:tinyint(1);not null"`
	DisplayID    string `gorm:"type:varchar(15);unique;not null;comment:展示给用户的id（账号或群号）"`
	Name         string `gorm:"type:varchar(50);not null"`
	Avatar       string `gorm:"type:varchar(255);not null"`
	OnlineStatus uint8  `gorm:"type:tinyint;not null;comment:0为离线，1为在线，2为隐身"`
	Status       string `gorm:"type:varchar(20);default:null"`
	Introduction string `gorm:"type:varchar(255);default:null"`
	GroupOwnerID uint32 `gorm:"type:int unsigned;default:null"`
}

func (Entity) TableComment() string {
	return "记录个人账户及群聊账户"
}
