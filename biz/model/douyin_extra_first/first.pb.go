// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: first.proto

package douyin_extra_first

import (
	_ "github.com/cloudwego/biz/model/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                   // 用户id
	Name          *string `protobuf:"bytes,2,req,name=name" json:"name,required" form:"name,required" query:"name,required"`                                            // 用户名称
	FollowCount   *int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count,omitempty" form:"follow_count" query:"follow_count"`           // 关注总数
	FollowerCount *int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count,omitempty" form:"follower_count" query:"follower_count"` // 粉丝总数
	IsFollow      *bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow,required" form:"is_follow,required" query:"is_follow,required"`         // true-已关注，false-未关注
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil && x.FollowCount != nil {
		return *x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil && x.FollowerCount != nil {
		return *x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil && x.IsFollow != nil {
		return *x.IsFollow
	}
	return false
}

type DouyinCommentActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       *string `protobuf:"bytes,1,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                                          // 用户鉴权token
	VideoId     *int64  `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,required" form:"video_id,required" query:"video_id,required"`                // 视频id
	ActionType  *int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,required" form:"action_type,required" query:"action_type,required"` // 1-发布评论，2-删除评论
	CommentText *string `protobuf:"bytes,4,opt,name=comment_text,json=commentText" json:"comment_text,omitempty" form:"comment_text" query:"comment_text"`              // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `protobuf:"varint,5,opt,name=comment_id,json=commentId" json:"comment_id,omitempty" form:"comment_id" query:"comment_id"`                       // 要删除的评论id，在action_type=2的时候使用
}

func (x *DouyinCommentActionRequest) Reset() {
	*x = DouyinCommentActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinCommentActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinCommentActionRequest) ProtoMessage() {}

func (x *DouyinCommentActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinCommentActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinCommentActionRequest) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinCommentActionRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *DouyinCommentActionRequest) GetVideoId() int64 {
	if x != nil && x.VideoId != nil {
		return *x.VideoId
	}
	return 0
}

func (x *DouyinCommentActionRequest) GetActionType() int32 {
	if x != nil && x.ActionType != nil {
		return *x.ActionType
	}
	return 0
}

func (x *DouyinCommentActionRequest) GetCommentText() string {
	if x != nil && x.CommentText != nil {
		return *x.CommentText
	}
	return ""
}

func (x *DouyinCommentActionRequest) GetCommentId() int64 {
	if x != nil && x.CommentId != nil {
		return *x.CommentId
	}
	return 0
}

type DouyinCommentActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	Comment    *Comment `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty" form:"comment" query:"comment"`                                                   // 评论成功返回评论内容，不需要重新拉取整个列表
}

func (x *DouyinCommentActionResponse) Reset() {
	*x = DouyinCommentActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinCommentActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinCommentActionResponse) ProtoMessage() {}

func (x *DouyinCommentActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinCommentActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinCommentActionResponse) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{2}
}

func (x *DouyinCommentActionResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinCommentActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *DouyinCommentActionResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                    // 视频评论id
	User       *User   `protobuf:"bytes,2,req,name=user" json:"user,required" form:"user,required" query:"user,required"`                                             // 评论用户信息
	Content    *string `protobuf:"bytes,3,req,name=content" json:"content,required" form:"content,required" query:"content,required"`                                 // 评论内容
	CreateDate *string `protobuf:"bytes,4,req,name=create_date,json=createDate" json:"create_date,required" form:"create_date,required" query:"create_date,required"` // 评论发布日期，格式 mm-dd
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{3}
}

func (x *Comment) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Comment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Comment) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *Comment) GetCreateDate() string {
	if x != nil && x.CreateDate != nil {
		return *x.CreateDate
	}
	return ""
}

type DouyinFavoriteActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *string `protobuf:"bytes,1,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                                          // 用户鉴权token
	VideoId    *int64  `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,required" form:"video_id,required" query:"video_id,required"`                // 视频id
	ActionType *int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,required" form:"action_type,required" query:"action_type,required"` // 1-点赞，2-取消点赞
}

func (x *DouyinFavoriteActionRequest) Reset() {
	*x = DouyinFavoriteActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteActionRequest) ProtoMessage() {}

func (x *DouyinFavoriteActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteActionRequest) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{4}
}

func (x *DouyinFavoriteActionRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *DouyinFavoriteActionRequest) GetVideoId() int64 {
	if x != nil && x.VideoId != nil {
		return *x.VideoId
	}
	return 0
}

func (x *DouyinFavoriteActionRequest) GetActionType() int32 {
	if x != nil && x.ActionType != nil {
		return *x.ActionType
	}
	return 0
}

type DouyinFavoriteActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
}

func (x *DouyinFavoriteActionResponse) Reset() {
	*x = DouyinFavoriteActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteActionResponse) ProtoMessage() {}

func (x *DouyinFavoriteActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteActionResponse) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{5}
}

func (x *DouyinFavoriteActionResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinFavoriteActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type DouyinFavoriteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,required" form:"user_id,required" query:"user_id,required"` // 用户id
	Token  *string `protobuf:"bytes,2,req,name=token" json:"token,required" form:"token,required" query:"token,required"`                      // 用户鉴权token
}

func (x *DouyinFavoriteListRequest) Reset() {
	*x = DouyinFavoriteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteListRequest) ProtoMessage() {}

func (x *DouyinFavoriteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteListRequest.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteListRequest) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{6}
}

func (x *DouyinFavoriteListRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *DouyinFavoriteListRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type DouyinFavoriteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,required" form:"status_code,required" query:"status_code,required"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`                        // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list" form:"video_list" query:"video_list"`                                  // 用户点赞视频列表
}

func (x *DouyinFavoriteListResponse) Reset() {
	*x = DouyinFavoriteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFavoriteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFavoriteListResponse) ProtoMessage() {}

func (x *DouyinFavoriteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFavoriteListResponse.ProtoReflect.Descriptor instead.
func (*DouyinFavoriteListResponse) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{7}
}

func (x *DouyinFavoriteListResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinFavoriteListResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *DouyinFavoriteListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *int64  `protobuf:"varint,1,req,name=id" json:"id,required" form:"id,required" query:"id,required"`                                                                    // 视频唯一标识
	Author        *User   `protobuf:"bytes,2,req,name=author" json:"author,required" form:"author,required" query:"author,required"`                                                     // 视频作者信息
	PlayUrl       *string `protobuf:"bytes,3,req,name=play_url,json=playUrl" json:"play_url,required" form:"play_url,required" query:"play_url,required"`                                // 视频播放地址
	CoverUrl      *string `protobuf:"bytes,4,req,name=cover_url,json=coverUrl" json:"cover_url,required" form:"cover_url,required" query:"cover_url,required"`                           // 视频封面地址
	FavoriteCount *int64  `protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,required" form:"favorite_count,required" query:"favorite_count,required"` // 视频的点赞总数
	CommentCount  *int64  `protobuf:"varint,6,req,name=comment_count,json=commentCount" json:"comment_count,required" form:"comment_count,required" query:"comment_count,required"`      // 视频的评论总数
	IsFavorite    *bool   `protobuf:"varint,7,req,name=is_favorite,json=isFavorite" json:"is_favorite,required" form:"is_favorite,required" query:"is_favorite,required"`                // true-已点赞，false-未点赞
	Title         *string `protobuf:"bytes,8,req,name=title" json:"title,required" form:"title,required" query:"title,required"`                                                         // 视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_first_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_first_proto_rawDescGZIP(), []int{8}
}

func (x *Video) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil && x.PlayUrl != nil {
		return *x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil && x.CoverUrl != nil {
		return *x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil && x.FavoriteCount != nil {
		return *x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil && x.CommentCount != nil {
		return *x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil && x.IsFavorite != nil {
		return *x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

var File_first_proto protoreflect.FileDescriptor

var file_first_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x18, 0x05, 0x20, 0x02, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x22, 0xb3, 0x01, 0x0a, 0x1d, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x82, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x72, 0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x61, 0x0a, 0x1f, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x4d, 0x0a, 0x1c,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x1d,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x38, 0x0a,
	0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x84, 0x02, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x30, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x02, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x08, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0xba,
	0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0xd2, 0xc1, 0x18, 0x17, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x32, 0xbf, 0x01, 0x0a, 0x15,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa5, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0xd2, 0xc1, 0x18, 0x18, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x32, 0xb5, 0x01,
	0x0a, 0x13, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0xca, 0xc1, 0x18, 0x16, 0x2f,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x65, 0x67, 0x6f, 0x2f, 0x62, 0x69,
	0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74,
}

var (
	file_first_proto_rawDescOnce sync.Once
	file_first_proto_rawDescData = file_first_proto_rawDesc
)

func file_first_proto_rawDescGZIP() []byte {
	file_first_proto_rawDescOnce.Do(func() {
		file_first_proto_rawDescData = protoimpl.X.CompressGZIP(file_first_proto_rawDescData)
	})
	return file_first_proto_rawDescData
}

var file_first_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_first_proto_goTypes = []interface{}{
	(*User)(nil),                         // 0: douyin.extra.first.User
	(*DouyinCommentActionRequest)(nil),   // 1: douyin.extra.first.douyin_comment_action_request
	(*DouyinCommentActionResponse)(nil),  // 2: douyin.extra.first.douyin_comment_action_response
	(*Comment)(nil),                      // 3: douyin.extra.first.Comment
	(*DouyinFavoriteActionRequest)(nil),  // 4: douyin.extra.first.douyin_favorite_action_request
	(*DouyinFavoriteActionResponse)(nil), // 5: douyin.extra.first.douyin_favorite_action_response
	(*DouyinFavoriteListRequest)(nil),    // 6: douyin.extra.first.douyin_favorite_list_request
	(*DouyinFavoriteListResponse)(nil),   // 7: douyin.extra.first.douyin_favorite_list_response
	(*Video)(nil),                        // 8: douyin.extra.first.Video
}
var file_first_proto_depIdxs = []int32{
	3, // 0: douyin.extra.first.douyin_comment_action_response.comment:type_name -> douyin.extra.first.Comment
	0, // 1: douyin.extra.first.Comment.user:type_name -> douyin.extra.first.User
	8, // 2: douyin.extra.first.douyin_favorite_list_response.video_list:type_name -> douyin.extra.first.Video
	0, // 3: douyin.extra.first.Video.author:type_name -> douyin.extra.first.User
	1, // 4: douyin.extra.first.CommentActionService.CreateCommentActionResponse:input_type -> douyin.extra.first.douyin_comment_action_request
	4, // 5: douyin.extra.first.FavoriteActionService.CreateFavoriteActionResponse:input_type -> douyin.extra.first.douyin_favorite_action_request
	6, // 6: douyin.extra.first.FavoriteListService.CreateFavoriteListResponse:input_type -> douyin.extra.first.douyin_favorite_list_request
	2, // 7: douyin.extra.first.CommentActionService.CreateCommentActionResponse:output_type -> douyin.extra.first.douyin_comment_action_response
	5, // 8: douyin.extra.first.FavoriteActionService.CreateFavoriteActionResponse:output_type -> douyin.extra.first.douyin_favorite_action_response
	7, // 9: douyin.extra.first.FavoriteListService.CreateFavoriteListResponse:output_type -> douyin.extra.first.douyin_favorite_list_response
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_first_proto_init() }
func file_first_proto_init() {
	if File_first_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_first_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_first_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinCommentActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_first_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinCommentActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_first_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_first_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_first_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_first_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_first_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFavoriteListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_first_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_first_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_first_proto_goTypes,
		DependencyIndexes: file_first_proto_depIdxs,
		MessageInfos:      file_first_proto_msgTypes,
	}.Build()
	File_first_proto = out.File
	file_first_proto_rawDesc = nil
	file_first_proto_goTypes = nil
	file_first_proto_depIdxs = nil
}
