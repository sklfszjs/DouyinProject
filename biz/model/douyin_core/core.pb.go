// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: core.proto

package douyin_core

import (
	"mime/multipart"
	"time"

	_ "github.com/cloudwego/biz/model/api"
)


type Comment struct {
	UserId int64 `json:"-"`
	VideoId int64 `json:"-"`
	CreatedAt time.Time `json:"-"`
	Id         int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                    // 视频评论id
	Content    string `protobuf:"bytes,3,req,name=content" json:"content,required" form:"content,required" query:"content,required"`                                 // 评论内容
	// CreateDate string `protobuf:"bytes,4,req,name=create_date,json=createDate" json:"create_date,required" form:"create_date,required" query:"create_date,required"` // 评论发布日期，格式 mm-dd
	User       *User   `protobuf:"bytes,2,req,name=user" json:"user,required" form:"user,required" query:"user,required" gorm:"-"`                                             // 评论用户信息
}


type DouyinUserRequest struct {

	UserId int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                      // 用户鉴权token
}


type DouyinUserResponse struct {

	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	User       *User   `protobuf:"bytes,3,req,name=user" json:"user,required" form:"user,required" query:"user,required"`                                              // 用户信息
}


type UserFollowers struct{
	UserId int64
	FollowerId int64
}

type UserFavVideos struct{
	UserId int64
	VideoId int64 
}

type UserVideos struct{
	UserId int64
	VideoId int64 
}


type User struct {
  Token         string `json:"-"`
	Id            int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                   // 用户id
	Name          string `protobuf:"bytes,2,req,name=name" json:"name,required" form:"name,required" query:"name,required" `                                            // 用户名称
	FollowCount   int64         // 关注总数
	FollowerCount int64   // 粉丝总数
	// IsFollow      bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow,required" form:"is_follow,required" query:"is_follow,required"`         // true-已关注，false-未关注
	UserLogin     *DouyinUserLoginRequest `json:"-" gorm:"-"`// gorm:"foreignKey:UserId;references:Id"`//用户与自己的登录信息，一对一的关系
	VideoList []*Video `json:"-" gorm:"-"`//gorm:"foreignKey:UserId;references:Id"`               //一名用户，多个视频，一对多的关系
	Follows []*User `json:"-" gorm:"-"`//gorm:"many2many:user_followers"`                  //用户之间，多对多的关系
	FavoriteVideos []*Video `json:"-" gorm:"-"`//gorm:"many2many:user_fav_videos"`             //用户与喜欢的视频之间多对多的关系
	Comments []*Comment `json:"-" gorm:"-"`// gorm:"foreignKey:UserId;references:Id"`  //一名用户，多条评论
	Is_follow bool 
	Avatar string 
	Background_image string 
	Signature string 
	Total_favorited int64 
	Work_count int64 
	Favorite_count int64 
}



type DouyinUserRegisterRequest struct {
  Username string `protobuf:"bytes,1,req,name=username" json:"username,required" form:"username,required" query:"username,required" ` // 注册用户名，最长32个字符
  Password string `protobuf:"bytes,2,req,name=password" json:"password,required" form:"password,required" query:"password,required" ` // 密码，最长32个字符
}



type DouyinUserRegisterResponse struct {

	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	UserId     int64  `protobuf:"varint,3,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"`                     // 用户id
	Token      string `protobuf:"bytes,4,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                                          // 用户鉴权token
}



type DouyinUserLoginRequest struct {

	UserId int64 `json:"-"`
	Username string `protobuf:"bytes,1,req,name=username" json:"username,required" form:"username,required" query:"username,required"` // 登录用户名
	Password string `protobuf:"bytes,2,req,name=password" json:"password,required" form:"password,required" query:"password,required"` // 登录密码
}



type DouyinUserLoginResponse struct {

	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	UserId     int64  `protobuf:"varint,3,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"`                     // 用户id
	Token      string `protobuf:"bytes,4,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                                          // 用户鉴权token
}



type DouyinPublishListRequest struct {

	UserId int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Token  string `protobuf:"bytes,2,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                      // 用户鉴权token
}



type DouyinPublishListResponse struct {

	StatusCode int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list" form:"video_list" query:"video_list"`                                  // 用户发布的视频列表
}



type Video struct {

	Id            int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                                    // 视频唯一标识
	Author        *User   `protobuf:"bytes,2,req,name=author" json:"author,required" form:"author,required" query:"author,required" gorm:"-"`//这里不能存作者，因为是一对多的关系。                                                        // 视频作者信息
	PlayUrl       string `protobuf:"bytes,3,req,name=play_url,json=playUrl" json:"play_url,required" form:"play_url,required" query:"play_url,required"`                                // 视频播放地址
	CoverUrl      string `protobuf:"bytes,4,req,name=cover_url,json=coverUrl" json:"cover_url,required" form:"cover_url,required" query:"cover_url,required"`                           // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,required" form:"favorite_count,required" query:"favorite_count,required"` // 视频的点赞总数
	CommentCount  int64  `protobuf:"varint,6,req,name=comment_count,json=commentCount" json:"comment_count,required" form:"comment_count,required" query:"comment_count,required"`      // 视频的评论总数
	IsFavorite    bool   `protobuf:"varint,7,req,name=is_favorite,json=isFavorite" json:"is_favorite,required" form:"is_favorite,required" query:"is_favorite,required"`                // true-已点赞，false-未点赞
	Title         string `protobuf:"bytes,8,req,name=title" json:"title,required" form:"title,required" query:"title,required"`                                                         // 视频标题
	Users         *User ` gorm:"-"`//gorm:"many2many:user_fav_videos;"`
    Comments      []*Comment  `json:"-" gorm:"-"`//gorm:"foreignKey:UserId;references:Id"`
    CreatedAt     time.Time   `json:"-"`
    UpdatedAt     time.Time   `json:"-"`
	UserId        int64 `json:"-"`
	FileName string `json:"-"`
}


type DouyinFeedRequest struct {

	LatestTime int64  `protobuf:"varint,1,opt,name=latest_time,json=latestTime" json:"latest_time,omitempty" form:"latest_time" query:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty" form:"token" query:"token"`                                          // 可选参数，登录用户设置
}



type DouyinFeedResponse struct {

	StatusCode int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list" form:"video_list" query:"video_list"`                                  // 视频列表
	NextTime   int64   `protobuf:"varint,4,opt,name=next_time,json=nextTime" json:"next_time,omitempty" form:"next_time" query:"next_time"`                            // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}



type DouyinPublishActionRequest struct {

	Token string `protobuf:"bytes,1,req,name=token" json:"token,required" form:"token,required" query:"token,required"` // 用户鉴权token
	Data  *multipart.FileHeader  `protobuf:"bytes,2,req,name=data" json:"data,required" form:"data,required" query:"data,required"`     // 视频数据
	Title string `protobuf:"bytes,3,req,name=title" json:"title,required" form:"title,required" query:"title,required"` // 视频标题
}

type DouyinPublishActionResponse struct {

	StatusCode int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
}

