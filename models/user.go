package models

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/erzaffasya/Go-Tracking-Mafindo/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Username        string `json:"username" valid:"required~Username is required" example:"Johndee" gorm:"not null;uniqueIndex;"`
	Email           string `json:"email" valid:"required~Email is required,email~Invalid email format" example:"johndee@gmail.com" gorm:"not null;uniqueIndex;"`
	Password        string `json:"password,omitempty" valid:"required~Password is required,minstringlength(6)~Your password must be at least 6 characters long" example:"12345678" gorm:"not null"`
	NamaLengkap     string `json:"nama_lengkap" valid:"required~Nama Lengkap is required" gorm:"not null"`
	Hp              string `json:"hp" valid:"required~Hp is required" example:"08133233215" `
	ProfileImageUrl string `json:"profile_image_url,omitempty" example:"https://avatars.dicebear.com/api/identicon/your-custom-seed.svg"`
	Alamat          string `json:"alamat" valid:"required~Alamat is required" example:"Jl. MT Haryono" `

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}

	u.Password = helpers.Hash(u.Password)
	return
}

func (u *User) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}
	return
}

type UserUsecase interface {
	Register(context.Context, *User) error
	Login(context.Context, *User) error
	Update(context.Context, User, uint) (User, error)
	Delete(context.Context, uint) error
}

type UserRepo interface {
	Register(context.Context, *User) error
	Login(context.Context, *User) error
	Update(context.Context, User, uint) (User, error)
	Delete(context.Context, uint) error
}
