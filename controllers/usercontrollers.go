package controllers

import (
	"backend-golang/config"
	"backend-golang/models"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	lock = sync.Mutex{}
)

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "This is Widura Hasta")
}

func CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	userModel := new(models.User)

	if err := c.Bind(userModel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	userModel.Password = string(hashedPassword)
	userModel.ID = uuid.New().String()

	if err := c.Validate(userModel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := config.DB.Create(userModel).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, userModel)
}
