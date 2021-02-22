package controllers

import (
	"log"
	"main/db"
	"main/models"

	"github.com/labstack/echo/v4"
)

type PageginateUser struct{}

func GetUsers(c echo.Context) error {
	var users []models.User
	model := db.DbManager()
	result := model.Scopes(db.Paginate(c)).Find(&users)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return c.JSON(200, users)
}
