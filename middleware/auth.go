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
	authorization := c.GetHeader("Authorization")

	if len(authorization) > 0 {
		splitToken := strings.Split(authorization, "Bearer ")

		if len(splitToken) > 0 {
			token := splitToken[1]
			var user model.User

			if err := db.First(&user, "token = ?", token).Error; err == nil {
				c.Set("user", &user)

				return
			}
		}
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "unauthorized",
	})
}
