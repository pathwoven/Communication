package model

import "time"

type UserSetting struct {
	UserID               uint32    `gorm:"type:int unsigned;primaryKey;not null"`
	Birthday             time.Time `gorm:"default:null"`
	Sex                  uint8     `gorm:"type:tinyint(3);not null;default:2;comment:0表示男，1表示女，2表示未知"`
	FontSize             uint8     `gorm:"type:tinyint;not null;default:16"`
	FontStyle            string    `gorm:"type:varchar(50);default:null"`
	Theme                uint8     `gorm:"type:tinyint;not null;default:0"`
	Background           string    `gorm:"type:varchar(255);default:null;comment:背景图片的存储地址"`
	ReceiveNotice        bool      `gorm:"type:tinyint(1);not null;default:true;comment:0表示没有新消息提醒，1表示有"`
	NoticeSound          string    `gorm:"type:varchar(255);default:null;comment:提示音的存储地址"`
	Password             string    `gorm:"type:varchar(255);not null"`
	Email                string    `gorm:"type:varchar(30);not null"`
	CreateTime           time.Time `gorm:"not null"`
	LastLoginTime        time.Time `gorm:"default:null"`
	IsDeleted            bool      `gorm:"type:tinyint(1);not null;default:false;comment:是否已注销"`
	FoundByAccountID     bool      `gorm:"type:tinyint(1);not null;default:true;comment:可以通过id被搜索"`
	FoundByNickname      bool      `gorm:"type:tinyint(1);not null;default:true;comment:可以通过昵称被搜索"`
	FoundByGroup         bool      `gorm:"type:tinyint(1);not null;default:true;comment:是否可通过群聊添加"`
	FoundByFriend        bool      `gorm:"type:tinyint(1);not null;default:true;comment:是否可被好友推荐"`
	AllowStrangerMessage bool      `gorm:"type:tinyint(1);not null;default:true;comment:允许陌生人发送消息"`
	IsInvisible          bool      `gorm:"type:tinyint(1);not null;default:false;comment:是否隐私登录"`
}

func (UserSetting) TableComment() string {
	return "记录用户设置"
}
