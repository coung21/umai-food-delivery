package model

import (
	"common"
	jwt "umai-auth-service/component"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	common.SqlModel
	Name     string        `json:"name" gorm:"column:name;"`
	Email    string        `json:"email" gorm:"column:email;"`
	Avatar   *common.Image `json:"avatar" gorm:"column:avatar"`
	Password string        `json:"password,omitempty" gorm:"column:password;"`
	Phone    string        `json:"phone" gorm:"column:phone;"`
	Role     string        `json:"-" gorm:"column:role;"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) DefaultRole() {
	u.Role = RoleCustomer
}
func (u *User) SetRestaurantRole() {
	u.Role = RoleRestaurant
}

func (u *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (u *User) SanitizePassword() {
	u.Password = ""
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserWithToken struct {
	User
	Token jwt.Token `json:"token"`
}

type UserUpdate struct {
	Name   *string       `json:"name" gorm:"column:name;"`
	Avatar *common.Image `json:"avatar" gorm:"column:avatar"`
	Phone  *string       `json:"phone" gorm:"column:phone;"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

const (
	RoleCustomer   = "customer"
	RoleRestaurant = "restaurant"
	RoleShipper    = "shipper"
	RoleAdmin      = "admin"
)
