package database

import (
	"Communication/internal/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() {
	dsn := "user:password@tcp(127.0.0.1:3306)/communication?charset=utf8mb4&parseTime=True&loc=Local"
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
	err = DB.AutoMigrate(&model.Chat{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
}
