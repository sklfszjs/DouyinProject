// Code generated by hertz generator.

package douyin_core

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	// "io/ioutil"
	"io"

	"github.com/cloudwego/biz/utils"

	douyin_core "github.com/cloudwego/biz/model/douyin_core"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreatePublishActionResponse .
// @router /douyin/publish/action/ [POST]
func CreatePublishActionResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		print("bind and validate error ,", err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := PublishVideo(req)

	c.JSON(consts.StatusOK, resp)
}

func PublishVideo(req douyin_core.DouyinPublishActionRequest) douyin_core.DouyinPublishActionResponse {

	db := utils.GetDBConnPool().GetConn()
	defer utils.GetDBConnPool().ReturnConn(db)

	users := make([]*douyin_core.User, 0)
	tx := db.Begin()

	result := tx.Where("token = ?", req.Token).Find(&users)
	if result.RowsAffected == 1 {
		fmt.Println("legal user")
		fmt.Println(req.Title)
		filename := strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(time.Now().Nanosecond()) + strconv.Itoa(int(users[0].Id))
		picname := filename + ".jpg"
		filename = filename + ".mp4"
		fmt.Println(filename)
		path_filename, err := WriteVideo(req.Data, strconv.Itoa(int(users[0].Id)), filename)

		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return douyin_core.DouyinPublishActionResponse{
				StatusCode: 1,
				StatusMsg:  "write video error",
			}
		}
		path_picname, compressvideoname, err := SetCover(path_filename, picname)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return douyin_core.DouyinPublishActionResponse{
				StatusCode: 1,
				StatusMsg:  "form cover error",
			}
		}
		playurl := fmt.Sprintf("http://%s:%d", utils.GetConfigs().IP, utils.GetConfigs().Port) + compressvideoname[1:]
		coverurl := fmt.Sprintf("http://%s:%d", utils.GetConfigs().IP, utils.GetConfigs().Port) + path_picname[1:]
		fmt.Println(playurl, coverurl)
		tx.Create(&douyin_core.Video{
			UserId:   users[0].Id,
			Title:    req.Title,
			PlayUrl:  playurl,
			CoverUrl: coverurl,
			FileName: compressvideoname,
		})
		videos := make([]*douyin_core.Video, 0)
		result := tx.Where("play_url = ?", playurl).Find(&videos)
		if result.RowsAffected != 1 {
			fmt.Println(&result.Statement.SQL, result.RowsAffected)
			panic("where is the new data?")
		}
		tx.Model(&douyin_core.User{Id: users[0].Id}).Updates(douyin_core.User{WorkCount: users[0].WorkCount + 1})
		tx.Create(&douyin_core.UserVideos{UserId: users[0].Id, VideoId: videos[0].Id})
		tx.Commit()
		return douyin_core.DouyinPublishActionResponse{
			StatusCode: 0,
			StatusMsg:  "update success",
		}
		// users[0].Is_follow = true

	} else {
		tx.Rollback()
		fmt.Println("illegal user")
		return douyin_core.DouyinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  "illegal token",
		}
	}

}

// 创建路径
func CreateDataDir(basePath string, folderName string) (string, error) {
	// folderName := time.Now().Format("2006-01-02")

	// folderPath := filepath.Join(basePath, folderName)
	fmt.Println(basePath, folderName)
	folderPath := basePath + folderName
	fmt.Println(folderPath)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.MkdirAll(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath, nil
}

// 图片处理
func WriteVideo(file *multipart.FileHeader, dirname string, filename string) (string, error) {
	filepoint, err := file.Open() //打开文件
	if err != nil {
		fmt.Println("file open err is", err)
		return "", err
	}
	defer filepoint.Close()

	//创建新文件进行存储
	pathname, err := CreateDataDir("./statistic/", dirname)
	if err != nil {
		fmt.Println("total create data dir err is", err)
		return "", errors.New("createdatadir error")
	}
	filename = pathname + "/" + filename
	fmt.Println(filename)
	newfile, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer newfile.Close()

	//把旧文件的内容放入新文件
	var context []byte = make([]byte, 1024)
	for {
		n, err := filepoint.Read(context)
		newfile.Write(context[:n])
		if err != nil {
			if err == io.EOF {
				return filename, nil
			} else {
				return "", err
			}
		}
	}

}

func SetCover(videoName, picName string) (string, string, error) {
	picPath, err := CreateDataDir("./statistic/", "img/")
	if err != nil {
		fmt.Println("total create data dir err is", err)
		return "", "", errors.New("createdatadir error")
	}
	picName = picPath + "/" + picName
	fmt.Println(picName)
	videoname := utils.GetCoverImgAndCompress(videoName, picName)
	return picName, videoname, nil
}
