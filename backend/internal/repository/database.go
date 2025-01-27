package repository

import (
	"Communication/config"
	"Communication/internal/repository/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
		config.AppConfig.DBCharset,
	)
	log.Println(dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// 自动迁移
	err = DB.AutoMigrate(
		&model.Chat{},
		&model.ChatTag{},
		&model.Code{},
		&model.CodeTag{},
		&model.Contact{},
		&model.DivideFriend{},
		&model.DivideGroup{},
		&model.Entity{},
		&model.Favorite{},
		&model.FavoriteTag{},
		&model.FriendApplication{},
		&model.GroupMemberInfo{},
		&model.GroupNotice{},
		&model.GroupNotification{},
		&model.GroupSetting{},
		&model.Message{},
		&model.MessageDeleted{},
		&model.MultipleMessage{},
		&model.Note{},
		&model.NoteTag{},
		&model.RecycleBin{},
		&model.Schedule{},
		&model.UserSetting{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
}
