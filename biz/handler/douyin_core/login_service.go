// Code generated by hertz generator.

package douyin_core

import (
	"context"

	douyin_core "github.com/cloudwego/biz/model/douyin_core"
	"github.com/cloudwego/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"fmt"
)

// CreateLoginResponse .
// @router /douyin/user/login/ [POST]
func CreateLoginResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := UserLogin(req)

	c.JSON(consts.StatusOK, resp)
}

func UserLogin(req douyin_core.DouyinUserLoginRequest) douyin_core.DouyinUserLoginResponse {
	db := utils.GetDBConnPool().GetConn()
	defer utils.GetDBConnPool().ReturnConn(db)

	fmt.Printf("%#v", req)
	username := req.Username
	password := req.Password
	password = utils.AesEncrypt(username, password)
	users := make([]*douyin_core.User, 0)
	tx := db.Begin()
	// t:=douyin_core.DouyinUserLoginRequest{Username: username, Password: password}

	result := tx.Table("douyin_user_login_requests").Joins("join users on users.name=douyin_user_login_requests.username").Where("users.name=? and douyin_user_login_requests.password=?", username, password).Select("id", "token").Find(&users)
	// result := tx.(&douyin_core.DouyinUserLoginRequest{Username: username, Password: password}).Joins("join users on users.name=douyin_user_login_requests.username").Find(&users)
	if result.RowsAffected > 0 {
		if result.RowsAffected > 1 {
			tx.Rollback()
			panic("same user in db")
		}
		fmt.Println("login success")
		fmt.Printf("%#v", users[0])
		fmt.Println("token", users[0].Token)
		tx.Commit()
		return douyin_core.DouyinUserLoginResponse{
			StatusCode: 0,
			StatusMsg:  "login success",
			UserId:     users[0].Id,
			Token:      users[0].Token,
		}
	} else {
		tx.Rollback()
		return douyin_core.DouyinUserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "username/password error",
		}
	}

}
