package models

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Tracking struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Judul     string     `json:"judul" valid:"required~Username is required" example:"Johndee" gorm:"not null;uniqueIndex;"`
	Gmaps     string     `json:"gmaps" valid:"required~Email is required,email~Invalid email format" example:"johndee@gmail.com" gorm:"not null;uniqueIndex;"`
	Latitude  string     `json:"latitude,omitempty" valid:"required~Password is required,minstringlength(6)~Your password must be at least 6 characters long" example:"12345678" gorm:"not null"`
	Longitude int        `json:"longitude,omitempty" valid:"required~Age is required,range(8|100)~Your age must be at least greater than 8 years old" example:"8" gorm:"not null"`
	Catatan   string     `json:"catatan" valid:"required~Username is required" example:"Johndee" gorm:"not null;uniqueIndex;"`
	Photo     string     `json:"photo,omitempty" example:"https://avatars.dicebear.com/api/identicon/your-custom-seed.svg"`
	UsersId   int        `json:"users_id" valid:"required~Username is required" example:"Johndee" gorm:"not null;uniqueIndex;"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *Tracking) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}

	return
}

func (u *Tracking) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}
	return
}

type TrackingUsecase interface {
	Fetch(context.Context, *[]Tracking, uint) error
	Store(context.Context, *Tracking) error
	GetByUserID(context.Context, *Tracking, uint) error
	Update(context.Context, Tracking, uint) (Tracking, error)
	Delete(context.Context, uint) error
}

type TrackingRepo interface {
	Fetch(context.Context, *[]Tracking, uint) error
	Store(context.Context, *Tracking) error
	GetByUserID(context.Context, *Tracking, uint) error
	Update(context.Context, Tracking, uint) (Tracking, error)
	Delete(context.Context, uint) error
}
