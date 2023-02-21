package utils

import (
	"fmt"

	"github.com/cloudwego/biz/model/douyin_core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Form_tables() {
	db, err := gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			GetConfigs().DBUsrName, GetConfigs().DBPassword, GetConfigs().DBName)),
		&gorm.Config{})
	if err != nil {
		fmt.Println("数据库链接错误", err)
	}
	db.AutoMigrate(douyin_core.DouyinUserLoginRequest{})
	db.AutoMigrate(douyin_core.UserFavVideos{})
	db.AutoMigrate(douyin_core.Comment{})
	db.AutoMigrate(douyin_core.User{})
	db.AutoMigrate(douyin_core.Video{})
	db.AutoMigrate(douyin_core.UserVideos{})
	db.AutoMigrate(douyin_core.UserFollowers{})

	return
}
