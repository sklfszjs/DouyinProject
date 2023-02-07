// Code generated by hertz generator.

package douyin_extra_second

import (
	"context"

	douyin_extra_second "github.com/cloudwego/biz/model/douyin_extra_second"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateMessageActionResponse .
// @router /douyin/message/action/ [POST]
func CreateMessageActionResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_extra_second.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(douyin_extra_second.DouyinRelationActionResponse)

	c.JSON(consts.StatusOK, resp)
}
