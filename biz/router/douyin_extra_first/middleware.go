// Code generated by hertz generator.

package DouyinExtraFirst

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		func(c context.Context, ctx *app.RequestContext) {
			fmt.Println("host:", string(ctx.Request.Header.Host()), "uri:", string(ctx.Request.Header.RequestURI()))
		},
	}
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createcomment_ctionresponseMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createfavorite_ctionresponseMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list0Mw() []app.HandlerFunc {
	// your code...
	return nil
}
func _createcommentlistresponseMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createfavoritelistresponseMw() []app.HandlerFunc {
	// your code...
	return nil
}
