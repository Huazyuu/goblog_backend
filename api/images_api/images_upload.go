package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/utils"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`
}

// ImageUploadView 上传图片,返回图片的url
func (imagesAps *ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	// 判断path是否存在
	basePath := global.Config.Upload.Path
	if _, err = os.ReadDir(basePath); err != nil {
		// 不存在就创建,存在返回nil
		if err = os.MkdirAll(basePath, fs.ModePerm); err != nil {
			global.Logger.Error(err)
		}
	}

	// response data
	var resList []FileUploadResponse

	for _, file := range fileList {
		extension := filepath.Ext(file.Filename)
		//fmt.Println(extension)
		if !utils.InStringList(strings.ToLower(extension), global.ImageTypeList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprint("上传图片格式不符合,请上传以下格式: ", global.ImageTypeList),
			})
			continue
		}

		filePath := path.Join(basePath, file.Filename)
		// 判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小为 %.2f MB,请上传: %d MB大小图片", size, global.Config.Upload.Size),
			})
			continue
		}

		if err = c.SaveUploadedFile(file, filePath); err != nil {
			global.Logger.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}

		resList = append(resList, FileUploadResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}
	res.OkWithData(resList, c)
}
