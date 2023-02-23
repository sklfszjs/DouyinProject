// Code generated by hertz generator.

package douyin_extra_first

import (
	"context"
	"fmt"

	douyin_core "github.com/cloudwego/biz/model/douyin_core"
	douyin_extra_first "github.com/cloudwego/biz/model/douyin_extra_first"
	"github.com/cloudwego/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateCommentListResponse .
// @router /douyin/comment/list/ [GET]
func CreateCommentListResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_extra_first.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := CommentList(req)

	c.JSON(consts.StatusOK, resp)
}

func CommentList(req douyin_extra_first.DouyinCommentListRequest) douyin_extra_first.DouyinCommentListResponse {
	db := utils.GetDBConnPool().GetConn()
	defer utils.GetDBConnPool().ReturnConn(db)

	fmt.Printf("%#v", req)
	users := make([]*douyin_core.User, 0)
	tx := db.Begin()
	result := tx.Where("token = ?", req.Token).Find(&users)
	if result.RowsAffected > 1 {
		tx.Rollback()
		panic("repeat token")
	}
	comments := make([]*douyin_core.Comment, 0)
	tx.Where("user_id = ? and video_id = ?", users[0].Id, req.VideoId).Find(&comments)
	for i := 0; i < len(comments); i++ {
		users := make([]*douyin_core.User, 0)
		tx.Where("id = ?", comments[i].UserId).Find(&users)
		comments[i].User = users[0]
	}
	tx.Commit()
	return douyin_extra_first.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: comments,
	}

}