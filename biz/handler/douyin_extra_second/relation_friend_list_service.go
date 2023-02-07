// Code generated by hertz generator.

package douyin_extra_second

import (
	"context"

	douyin_extra_second "github.com/cloudwego/biz/model/douyin_extra_second"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateRelationFriendListResponse .
// @router /douyin/relatioin/friend/list/ [GET]
func CreateRelationFriendListResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_extra_second.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_extra_second.DouyinRelationFriendListResponse)

	c.JSON(consts.StatusOK, resp)
}
