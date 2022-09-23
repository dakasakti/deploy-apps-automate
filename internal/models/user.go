package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-" bson:"-"`
	ID         uint   `json:"id" bson:"id"`
	Fullname   string `json:"fullname" bson:"fullname,omitempty"`
	Username   string `json:"username" bson:"username,omitempty"`
	Password   string `json:"-" bson:"password,omitempty"`
}

type UserRequest struct {
	Fullname string `json:"fullname" bson:"fullname" validate:"required,min=3,max=25"`
	Username string `json:"username" bson:"username" validate:"required,lowercase,min=5,max=25"`
	Password string `json:"password" bson:"password" validate:"required,min=10,max=255"`
}

type UserUpdateRequest struct {
	Fullname string `json:"fullname,omitempty" bson:"fullname,omitempty" validate:"omitempty,min=3,max=25"`
	Username string `json:"username,omitempty" bson:"username,omitempty" validate:"omitempty,lowercase,min=5,max=25"`
	Password string `json:"password,omitempty" bson:"password,omitempty" validate:"omitempty,min=10,max=255"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required,lowercase,min=5,max=25"`
	Password string `json:"password" validate:"required,min=10,max=255"`
}
