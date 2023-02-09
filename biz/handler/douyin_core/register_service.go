// Code generated by hertz generator.

package douyin_core

import (
	"context"

	douyin_core "github.com/cloudwego/biz/model/douyin_core"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateRegisterResponse .
// @router /douyin/user/register/ [POST]
func CreateRegisterResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_core.DouyinUserRegisterResponse)

	c.JSON(consts.StatusOK, resp)
}