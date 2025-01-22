package imgServer

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
)

var WhiteImageList = []string{
	".bmp", ".jpg", ".png",
	".tif", ".gif", ".pcx",
	".tga", ".exif", ".fpx",
	".svg", ".psd", ".cdr",
	".pcd", ".dxf", ".ufo",
	".eps", ".ai", ".raw",
	".WMF", ".webp", ".avif",
	".apng",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`
}

// ImageUploadService image upload service
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	filePath := path.Join(global.Config.Upload.Path, file.Filename)
	var fileType ctype.ImageType
	res.FileName = filePath

	// 白名单
	extension := filepath.Ext(fileName)
	if !utils.InStringList(strings.ToLower(extension), global.ImageTypeList) {
		res.Msg = "非法文件"
		return
	}

	// 判断大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小为 %.2f MB,请上传: %d MB大小图片", size, global.Config.Upload.Size)
		return
	}

	// md5 加密hash
	fileObj, err := file.Open()
	if err != nil {
		global.Logger.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	if err != nil {
		global.Logger.Error(err)
	}
	imageHash := utils.Md5(byteData)

	// 去数据库查询文件是否存在
	var bannerModel models.BannerModel

	// gorm query 自定义session不输出log
	err = global.DB.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}).Take(&bannerModel, "hash= ?", imageHash).Error
	if err == nil {
		// 图片重复 不需要入库
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	fileType = ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true

	// qiniu
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Logger.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu
	}

	// 入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
