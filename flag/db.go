package flag

import (
	"gvb_server/global"
	"gvb_server/models"
)

func MakeMigration() {
	var err error
	// _ = global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	_ = global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.AdvertModel{},
		// &models.ArticleModel{},
		&models.BannerModel{},
		&models.CommentModel{},
		&models.FadeBackModel{},
		&models.LoginDataModel{},
		&models.MenuBannerModel{},
		&models.MenuModel{},
		&models.MessageModel{},
		&models.TagModel{},
		&models.UserModel{},
	)
	if err != nil {
		global.Logger.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Logger.Info("[ success ] 生成数据库表结构成功！")
}
