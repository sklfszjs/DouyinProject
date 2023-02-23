// Code generated by hertz generator. DO NOT EDIT.

package DouyinCore

import (
	douyin_core "github.com/cloudwego/biz/handler/douyin_core"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {
	
	root := r.Group("/", rootMw()...)
	{
		root.GET("/statistic/*path",append(_sendvideoMw(),douyin_core.SendFile)...)
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_createfeedresponseMw(), douyin_core.CreateFeedResponse)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action := _publish.Group("/action", _actionMw()...)
				_action.POST("/", append(_createpublish_ctionresponseMw(), douyin_core.CreatePublishActionResponse)...)
			}
			{
				_list := _publish.Group("/list", _listMw()...)
				_list.GET("/", append(_createpublishlistresponseMw(), douyin_core.CreatePublishListResponse)...)
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_createuserresponseMw(), douyin_core.CreateUserResponse)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_createloginresponseMw(), douyin_core.CreateLoginResponse)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_createregisterresponseMw(), douyin_core.CreateRegisterResponse)...)
			}
		}
	}
}
