package model

type GroupMemberInfo struct {
	GroupID       uint32 `gorm:"type:int unsigned;primaryKey;not null"`
	UserID        uint32 `gorm:"type:int unsigned;primaryKey;not null;index:user_id_idx"`
	GroupNickname string `gorm:"type:varchar(30);default:null"`
	IsBanned      bool   `gorm:"type:tinyint(1);not null;default:false"`
	IsManager     bool   `gorm:"type:tinyint(1);not null"`
}

func (GroupMemberInfo) TableComment() string {
	return "记录群聊中的成员信息"
}
