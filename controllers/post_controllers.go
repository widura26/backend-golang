package controllers

import (
	"backend-golang/config"
	"backend-golang/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	lock.Lock()
	post := new(models.Post)

	if err := c.Bind(post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(post); err != nil {
		return err
	}

	if err := config.DB.Create(post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
	})

}

func GetPosts(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var posts []models.Post
	getAllData := config.DB.Find(&posts)

	if getAllData.Error != nil {
		return c.JSON(http.StatusBadRequest, getAllData.Error.Error())
	}

	return c.JSON(http.StatusOK, posts)
}
