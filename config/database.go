package config

import (
	"github.com/TesyarRAz/sa-api-lks-jabar-2021/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/sa_api_lks_jabar_2021?parseTime=true"))

	if err != nil {
		panic("Cant connect to database")
	}

	db.AutoMigrate(model.User{})
	db.AutoMigrate(model.Menu{})

	return &Database{
		db: db,
	}
}

func (db *Database) Middleware(c *gin.Context) {
	c.Set("db", db.db)
}
