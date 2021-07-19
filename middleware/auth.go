package middleware

import (
	"net/http"
	"strings"

	"github.com/TesyarRAz/sa-api-lks-jabar-2021/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorizedUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	token := strings.Split(c.GetHeader("Authorization"), "Bearer ")[1]

	var user model.User

	if err := db.First(&user, "token = ?", token).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})

		return
	}

	c.Set("user", &user)
}
