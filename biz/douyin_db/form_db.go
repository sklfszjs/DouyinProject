package douyin_db

import (
	"fmt"

	"github.com/cloudwego/biz/model/douyin_core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Form_tables() error {
	db, err := gorm.Open(
		mysql.Open("root:@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		fmt.Println("数据库链接错误", err)
	}
  db.AutoMigrate(douyin_core.DouyinUserLoginRequest{})
  db.AutoMigrate(douyin_core.Comment{})
	migErr := db.AutoMigrate(&douyin_core.User{VideoList: []*douyin_core.Video{}, Follows: []*douyin_core.User{},
		Comments:       []*douyin_core.Comment{},
		FavoriteVideos: []*douyin_core.Video{},
    UserLogin: &douyin_core.DouyinUserLoginRequest{} })
	return migErr
}
