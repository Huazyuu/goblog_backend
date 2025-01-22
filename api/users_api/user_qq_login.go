package users_api

import (
	"github.com/gin-gonic/gin"
)

// todo qq login view
// https://www.wolai.com/fengfeng/mCyj3r81gxyax3T3nuQycV

// QQLoginView qq登录，返回token，用户信息需要从token中解码
// @Tags 用户管理
// @Summary qq登录
// @Description qq登录，返回token，用户信息需要从token中解码
// @Param code query string  true  "qq登录的code"
// @Router /api/login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UsersApi) QQLoginView(c *gin.Context) {

}
