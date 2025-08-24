package models

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID       string `json:"id" gorm:"type:varchar(36);primaryKey"`
	Name     string `json:"name" form:"name" gorm:"type:varchar(100)"`
	Username string `json:"username" form:"username" gorm:"type:varchar(100)"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(255);uniqueIndex"`
	Password string `json:"password" form:"password" gorm:"type:varchar(255)"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 20)),
	)
}
