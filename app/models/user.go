package models

import (
	"be-medela-potentia/app/helpers"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	ID       string   `gorm:"primaryKey" json:"id" form:"id" alias:"id" swaggerignore:"true"`
	Name     *string  `gorm:"size:255" json:"name" validate:"required"`
	Username *string  `gorm:"size:255" json:"username"`
	Password string   `gorm:"column:password" json:"password,omitempty" validate:"required"`
	RoleId   uint     `gorm:"size:255" json:"role_id,omitempty" validate:"required"`
	UserRole UserRole `json:"user_roles,omitempty" gorm:"foreignKey:RoleId" swaggerignore:"true"`

	// gorm.Model
	CreatedAt time.Time      `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updated_at" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

func (User) TableName() string {
	return "tbl_users"
}

func (d *User) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = helpers.GenerateID()
	return
}

type DataLogin struct {
	UserContact string `json:"username" form:"username" alias:"username" validate:"required,min=3"`
	Password    string `json:"password" form:"password" alias:"password" validate:"required,min=3"`
}

type JwtUser struct {
	ID       string   ` json:"id" form:"id" alias:"id"`
	Name     *string  `json:"name" `
	Username *string  `json:"username" `
	RoleId   uint     `json:"role_id" `
	UserRole UserRole `json:"user_roles" gorm:"foreignKey:RoleId"`
	jwt.RegisteredClaims
}
