package intercept

import (
	"fmt"
	"gin-template/conf"
	"gin-template/global"
	"gin-template/model"
	"gin-template/utils"
	"github.com/gin-gonic/gin"
)

// AuthApp 是一个中间件，用于保护路由
func AuthApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userId int64
			token  = c.GetHeader(conf.LocalToken)
			db     = global.DB
			err    error
		)

		// 打印请求地址
		global.Log.Info("Request URL: " + c.Request.URL.Path)
		if token == "" || len(token) < 10 {
			utils.MessageResponse(c, conf.TokenFail, "token is null", "令牌为null")
			c.Abort()
			return
		}

		user, err := model.UserSelectIdByToken(db, token)
		if err != nil {
			if err.Error() != "record not found" {
				fmt.Println(err.Error())
			}
			utils.MessageResponse(c, conf.TokenFail, "token invalid or network error", "token失效或网络故障")
			c.Abort()
			return
		}
		userId = int64(user.ID)

		if !utils.CheckSpecialCharacters(&token) {
			utils.MessageResponse(c, conf.TokenFail, "token is invalid", "令牌无效")
			c.Abort()
			return
		}
		//检查token 有效时间
		if !utils.CheckTokenValidityTime(&user.Token) {
			utils.MessageResponse(c, conf.TokenFail, "token is exceed", "令牌超过")
			c.Abort()
			return
		}

		//刷新token有效时间
		if err = model.UserRefreshToken(db, userId, user.Token); err != nil {
			utils.MessageResponse(c, conf.TokenFail, "db UserRefreshAppToken err", "刷新失败")
			c.Abort()
			return
		}

		c.Set(conf.LocalUseridUint, uint(userId))
		c.Set(conf.LocalUseridInt64, userId)
		c.Next()
	}
}

// AuthWebOperationProtected 用于保护网页操作的路由
func AuthWebOperationProtected(rights string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authority = c.GetHeader(conf.LocalAuthority)
		)

		//刷新token有效时间
		if authority != "all" {
			if authority != rights {
				utils.MessageResponse(c, conf.TokenFail, "un authorized operation", "无权限操作")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// ManagerAuthProtected 是用于保护管理员操作的路由
//func ManagerAuthProtected() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var (
//			token = c.GetHeader(conf.LocalToken)
//			db    = global.DB
//			err   error
//		)
//
//		// 打印请求地址
//		global.Log.Info("Request URL: ", zap.Field{Key: c.Request.URL.Path})
//		if token == "" || len(token) < 10 {
//			c.JSON(utils.MessageResponse(conf.TokenFail, "token is null", ""))
//			c.Abort()
//			return
//		}
//
//		user, err := model.UserSelectIdByManagerToken(db, token)
//		if err != nil {
//			c.JSON(utils.MessageResponse(conf.TokenFail, "token is invalid", ""))
//			c.Abort()
//			return
//		}
//
//		if !utils.CheckSpecialCharacters(&token) {
//			c.JSON(utils.MessageResponse(conf.TokenFail, "token is invalid", ""))
//			c.Abort()
//			return
//		}
//		//检查token 有效时间
//		if !utils.CheckTokenValidityTime(&user.ManagerToken) {
//			c.JSON(utils.MessageResponse(conf.TokenFail, "token is exceed", ""))
//			c.Abort()
//			return
//		}
//
//		//刷新manager_token有效时间 要加一个时间戳
//		if err = model.UserRefreshManagerToken(db, int64(user.ID), user.ManagerToken); err != nil {
//			c.JSON(utils.MessageResponse(conf.TokenFail, "db UserRefreshManagerToken ", err.Error()))
//			c.Abort()
//			return
//		}
//		c.Set(conf.AdminUseridInt64, int64(user.ID))
//		//c.Set(conf.AdminUsername, user.Name)
//		//c.Set(conf.ManageRole, user.Role)
//		c.Set(conf.ManageUser, user)
//		c.Next()
//	}
//}
