/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 18:24
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"safe-community/core/controller/vo"
	"safe-community/core/dao/models"
	"safe-community/core/service"
)

func UserShow(c *gin.Context) {
	uid := c.Query("uid")
	log.Println("uid: ", uid)
	c.JSON(http.StatusOK, NewFrontData(OK, "hello user"))
}

func UserPersonalAccount(c *gin.Context) {
	alia, ok := c.Get("alia")
	if !ok {
		c.JSON(http.StatusBadRequest, NewFrontData(ErrorInvalidParamOfInterface, nil))
		return
	}

	info, err := service.UserInfoGet(alia.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewFrontData(ErrorDatabase, nil))
		return
	}

	c.JSON(http.StatusOK, NewFrontData(OK, *info))

}

func UserAccountSignUp(c *gin.Context) {
	var info vo.UserInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, NewFrontData(ErrorInvalidParamOfInterface, nil))
		return
	}

	if err := service.UserSignupPost(info); err != nil {
		c.JSON(http.StatusInternalServerError, NewFrontData(ErrorDatabase, nil))
		return
	}

	c.JSON(http.StatusOK, NewFrontData(OK, nil))
}

func UserTemperature(c *gin.Context) {
	var t vo.UserTemperature
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, NewFrontData(ErrorInvalidParamOfInterface, nil))
		return
	}

	if err := service.UserTemperaturePost(t); err != nil {
		c.JSON(http.StatusInternalServerError, NewFrontData(ErrorDatabase, nil))
		return
	}

	c.JSON(http.StatusOK, NewFrontData(OK, nil))
}

func UserPersonalInfo(c *gin.Context) {
	var info models.UserInfo
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, NewFrontData(ErrorInvalidParamOfInterface, nil))
		return
	}

	if err := service.UserInfoPost(&info); err != nil {
		c.JSON(http.StatusInternalServerError, NewFrontData(ErrorDatabase, nil))
		return
	}

	c.JSON(http.StatusOK, NewFrontData(OK, nil))
}

func UserPersonalTemperature(c *gin.Context) {
	alia, ok := c.Get("alia")
	if !ok {
		c.JSON(http.StatusBadRequest, NewFrontData(ErrorInvalidParamOfInterface, nil))
		return
	}

	temps, err := service.UserTemperatureGet(alia.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewFrontData(ErrorDatabase, nil))
		return
	}

	c.JSON(http.StatusOK, NewFrontData(OK, temps))
}