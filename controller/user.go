package controller

import (
	"net/http"

	"github.com/TesyarRAz/sa-api-lks-jabar-2021/model"
	"github.com/TesyarRAz/sa-api-lks-jabar-2021/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var v struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	if err := c.ShouldBind(&v); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	var user model.User

	if err := db.Where("username", v.Username).Where("password", v.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	if err := db.Model(&user).Update("token", util.RandomText(18)).Error; err != nil {
		c.JSON(http.StatusForbidden, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	user.Token = util.RandomText(18)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func InfoUser(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
