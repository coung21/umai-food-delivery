package model

import (
	"umai-auth-service/common"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	common.SqlModel
	Name     string `json:"name" gorm:"column:name;"`
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password,omitempty" gorm:"column:password;"`
	Phone    string `json:"phone" gorm:"column:phone;"`
	Role     Role   `json:"-" gorm:"column:role;"`
}

func (User) TableName() string {
	return "users"
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

type Role int

const (
	RoleUser Role = iota
	RoleRestaurant
	RoleShipper
	RoleAdmin
)

func (role Role) String() string {
	switch role {
	case RoleUser:
		return "user"
	case RoleRestaurant:
		return "restaurant"
	case RoleShipper:
		return "shipper"
	case RoleAdmin:
		return "admin"
	default:
		return "user"
	}
}
