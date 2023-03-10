// Code generated by hertz generator. DO NOT EDIT.

package DouyinExtraFirst

import (
	// "go/doc/comment"

	douyin_extra_first "github.com/cloudwego/biz/handler/douyin_extra_first"
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
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_comment := _douyin.Group("/comment", _commentMw()...)
			{
				_action := _comment.Group("/action", _actionMw()...)
				_action.POST("/", append(_createcomment_ctionresponseMw(), douyin_extra_first.CreateCommentActionResponse)...)
			}
			{
				_list := _comment.Group("/list", _listMw()...)
				_list.GET("/", append(_createcommentlistresponseMw(), douyin_extra_first.CreateCommentListResponse)...)
			}
		}
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			{
				_action0 := _favorite.Group("/action", _action0Mw()...)
				_action0.POST("/", append(_createfavorite_ctionresponseMw(), douyin_extra_first.CreateFavoriteActionResponse)...)
			}
			{
				_list0 := _favorite.Group("/list", _list0Mw()...)
				_list0.GET("/", append(_createfavoritelistresponseMw(), douyin_extra_first.CreateFavoriteListResponse)...)
			}
		}
	}
}
