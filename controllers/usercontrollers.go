package controllers

import (
	"backend-golang/config"
	"backend-golang/dto"
	"backend-golang/models"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusAccepted, "Hello World")
}

func Signup(c echo.Context) error {
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

func Signin(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	req := new(dto.LoginRequest)
	user := new(models.User)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := config.DB.Where("username = ?", req.Username).First(user).Error; err != nil {
		return echo.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.ErrUnauthorized
	}

	claims := &dto.JwtCustomClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
