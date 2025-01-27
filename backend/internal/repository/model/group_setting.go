package model

type GroupSetting struct {
	GroupID           int    `gorm:"primaryKey;not null"`
	AllowFoundByID    bool   `gorm:"not null;default:true;comment:允许被id搜索"`
	AllowFoundByName  bool   `gorm:"not null;default:true;comment:允许被群名搜索"`
	AllowMemberInvite bool   `gorm:"not null;default:true;comment:允许成员邀请"`
	JoinAuthMethod    int    `gorm:"not null;default:1;comment:0表示无需验证，1表示需要发送验证消息，2表示设置了问题"`
	AllowDirectInvite bool   `gorm:"not null;default:true;comment:成员邀请无需验证"`
	AuthQuestion      string `gorm:"type:varchar(255);default:null;comment:验证问题"`
	AuthAnswer        string `gorm:"type:varchar(255);default:null;comment:答案"`
}
