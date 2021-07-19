package controller

import (
	"net/http"

	"github.com/TesyarRAz/sa-api-lks-jabar-2021/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func IndexMenu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(*model.User)

	var menus []model.Menu

	if err := db.Where("user_id = ?", user.ID).Find(&menus).Error; err != nil {
		c.JSON(http.StatusForbidden, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": menus,
	})
}

func ShowMenu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(*model.User)
	id := c.Param("id")

	if len(id) == 0 {
		c.Status(http.StatusNotFound)

		return
	}

	var menu model.Menu

	if err := db.Where("user_id = ?", user.ID).First(&menu, id).Error; err != nil {
		c.JSON(http.StatusForbidden, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": menu,
	})
}

func StoreMenu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(*model.User)

	var menu model.Menu

	if err := c.ShouldBind(&menu); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	menu.UserID = user.ID

	if err := db.Omit(clause.Associations).Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil input data",
	})
}

func UpdateMenu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(*model.User)
	id := c.Param("id")

	if len(id) == 0 {
		c.Status(http.StatusNotFound)

		return
	}

	var (
		menu    model.Menu
		newMenu model.Menu
	)

	if err := db.Where("user_id = ?", user.ID).First(&menu, id).Error; err != nil {
		c.JSON(http.StatusForbidden, err.Error())

		return
	}

	if err := c.ShouldBind(&newMenu); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}

	newMenu.UserID = user.ID

	if err := db.Model(&menu).Updates(&newMenu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil rubah data",
	})
}

func DestroyMenu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(*model.User)
	id := c.Param("id")

	if len(id) == 0 {
		c.Status(http.StatusNotFound)

		return
	}

	if err := db.Where("user_id", user.ID).Delete(&model.Menu{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil hapus data",
	})
}
