package entity

import (
	"github.com/adamnasrudin03/my-gram/pkg/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null;uniqueIndex" json:"username" `
	Email    string `gorm:"not null;uniqueIndex" json:"email" `
	Password string `gorm:"not null" json:"password,omitempty"`
	Age      uint64 `gorm:"not null" json:"age"`
	GORMModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass

	return
}
