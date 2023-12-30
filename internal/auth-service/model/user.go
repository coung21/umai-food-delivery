package model

import (
	"common"
	"time"
	jwt "umai-auth-service/component"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	common.SqlModel
	Name     string        `json:"name" gorm:"column:name;not null;type:varchar(255);"`
	Email    string        `json:"email" gorm:"column:email;not null;unique;type:varchar(255);"`
	Avatar   *common.Image `json:"avatar" gorm:"column:avatar;type:json;"`
	Password string        `json:"password,omitempty" gorm:"column:password;not null;type:text;"`
	Phone    string        `json:"phone" gorm:"column:phone;not null;unique;type:varchar(255);"`
	Role     string        `json:"-" gorm:"column:role;type:ENUM('customer', 'restaurant', 'shipper', 'admin');default:'customer';not null"`
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
	Name     *string       `json:"name" gorm:"column:name;"`
	Avatar   *common.Image `json:"avatar" gorm:"column:avatar"`
	Phone    *string       `json:"phone" gorm:"column:phone;"`
	UpdateAt time.Time     `json:"updated_at" gorm:"column:updated_at"`
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
