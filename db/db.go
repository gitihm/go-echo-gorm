package db

import (
	"fmt"
	"main/config"
	"main/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_HOST, configuration.DB_PORT, configuration.DB_NAME)
	db, err = gorm.Open(mysql.Open(connect_string), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
}

func DbManager() *gorm.DB {
	return db
}

func Paginate(c echo.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page == 0 {
			page = 1
		}

		sort := c.QueryParam("sort")

		if sort == "" {
			sort = "desc"
		}

		order := c.QueryParam("order")
		if order == "" {
			order = "id"
		}

		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit).Order(fmt.Sprintf("%s %s", order, sort))
	}
}
