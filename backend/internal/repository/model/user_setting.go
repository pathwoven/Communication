package model

import "time"

type UserSetting struct {
	UserID               int       `gorm:"primaryKey;not null"`
	Birthday             time.Time `gorm:"default:null"`
	Sex                  int       `gorm:"not null;default:2;comment:0表示男，1表示女，2表示未知"`
	FontSize             int       `gorm:"not null;default:16"`
	FontStyle            string    `gorm:"type:varchar(50);default:null"`
	Theme                int       `gorm:"not null;default:0"`
	Background           string    `gorm:"type:varchar(255);default:null;comment:背景图片的存储地址"`
	ReceiveNotice        bool      `gorm:"not null;default:true;comment:0表示没有新消息提醒，1表示有"`
	NoticeSound          string    `gorm:"type:varchar(255);default:null;comment:提示音的存储地址"`
	Password             string    `gorm:"type:varchar(255);not null"`
	Email                string    `gorm:"type:varchar(30);not null"`
	CreateTime           time.Time `gorm:"not null"`
	LastLoginTime        time.Time `gorm:"default:null"`
	IsDeleted            bool      `gorm:"not null;default:false;comment:是否已注销"`
	FoundByAccountID     bool      `gorm:"not null;default:true;comment:可以通过id被搜索"`
	FoundByNickname      bool      `gorm:"not null;default:true;comment:可以通过昵称被搜索"`
	FoundByGroup         bool      `gorm:"not null;default:true;comment:是否可通过群聊添加"`
	FoundByFriend        bool      `gorm:"not null;default:true;comment:是否可被好友推荐"`
	AllowStrangerMessage bool      `gorm:"not null;default:true;comment:允许陌生人发送消息"`
	IsInvisible          bool      `gorm:"not null;default:false;comment:是否隐私登录"`
}
