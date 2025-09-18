package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID      uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Title   string    `json:"title" form:"title" gorm:"type:varchar(255)"`
	Content string    `json:"content" form:"content" gorm:"type:text"`
}

func (u Post) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Title, validation.Required),
		validation.Field(&u.Content, validation.Required),
	)
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
