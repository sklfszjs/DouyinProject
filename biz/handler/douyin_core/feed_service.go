// Code generated by hertz generator.

package douyin_core

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cloudwego/biz/utils"

	douyin_core "github.com/cloudwego/biz/model/douyin_core"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateFeedResponse .
// @router /douyin/feed/ [GET]
func CreateFeedResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := FeedService(req)
	c.JSON(consts.StatusOK, resp)
}

func SendFile(ctx context.Context, c *app.RequestContext) {
	fileName := c.Param("path")
	fmt.Println("file name is ", fileName)
	content, err := os.ReadFile("./statistic/" + fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	words := strings.Split(fileName, ".")
	suffix := words[len(words)-1]
	if suffix == "mp4" {
		c.Data(200, "video/mp4", content)
	}
	if suffix == "jpg" {
		c.Data(200, "image/jpeg", content)
	}
}

func SendVideo(fileName string, c *app.RequestContext) {

	fmt.Println("file name is ", fileName)
	content, err := os.ReadFile("./statistic/" + fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	c.Data(200, "video/mp4", content)
}

func FeedService(req douyin_core.DouyinFeedRequest) douyin_core.DouyinFeedResponse {
	db := utils.GetDBConnPool().GetConn()
	defer utils.GetDBConnPool().ReturnConn(db)

	token := req.Token
	users := make([]*douyin_core.User, 0)
	tx := db.Begin()

	result := tx.Where("token = ? ", token).Find(&users)
	if result.RowsAffected == 0 {
		fmt.Println("does not login in visit")
	} else {
		fmt.Println("user login visit")
	}
	videos := make([]*douyin_core.Video, 0)
	tx.Where("UNIX_TIMESTAMP(created_at) < ?", req.LatestTime).Order("created_at").Limit(30).Find(&videos)
	for i := 0; i < len(videos); i++ {
		users := make([]*douyin_core.User, 0)
		tx.Where("id = ?", videos[i].UserId).Find(&users)
		userfavvideo := make([]*douyin_core.UserFavVideos, 0)
		tx.Where("userid = ? and videoid = ", videos[i].UserId, videos[i].Id).Find(&userfavvideo)
		if len(userfavvideo) != 0 {
			videos[i].IsFavorite = true
		}
		videos[i].Author = users[0]
	}
	tx.Commit()
	var nexttime int64 = 1 << 62
	for _, v := range videos {
		if nexttime > v.CreatedAt.Unix() {
			nexttime = v.CreatedAt.Unix()
		}
	}

	if nexttime > time.Now().Unix() {
		nexttime = time.Now().Unix()
	}
	//TODO这里逻辑得重新理一下

	return douyin_core.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "operation success",
		VideoList:  videos,
		NextTime:   nexttime,
	}

}
