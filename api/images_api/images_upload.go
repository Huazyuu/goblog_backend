package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/img_service"
	"io/fs"
	"os"
)

// ImageUploadView 上传图片,返回图片的url
func (imagesApi *ImagesApi) ImageUploadView(c *gin.Context) {
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
	var resList []img_service.FileUploadResponse

	for _, file := range fileList {
		// 上传文件 调用服务层
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		// false
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// true and check qiniu enable
		if !global.Config.QiNiu.Enable {
			// save on local
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Logger.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
