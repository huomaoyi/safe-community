/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 18:17
 */

package controller

import (
	"github.com/gin-gonic/gin"
	swag "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginglog "github.com/szuecs/gin-glog"
	"net/http"
	"time"
)

//CORSAllow 跨域允许设置
func CORSAllow() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization,Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}


func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", swag.WrapHandler(swaggerFiles.Handler))
	router.Use(ginglog.Logger(3*time.Second), CORSAllow())

	//前台公共接口
	frontPublic := router.Group("/user")
	{
		frontPublic.POST("/signup", UserAccountSignUp)       					// 用户账户注册
		frontPublic.POST("/temperature", UserTemperature)       				// 用户温度上报
		frontPublic.GET("personal/account", UserPersonalAccount)				// 用户个人信息查看
		frontPublic.GET("/personal/temperature", UserPersonalTemperature)		// 用户历史温度列表
	}

	return router
}