package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint32 `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(64);" json:"name"`
	Username string `gorm:"type:varchar(64);not_null;unique:not null" json:"username"`
	Password string `gorm:"varhcar(255);" json:"-"`
	Verified bool   `gorm:"default:false;->:false;<-:create" json:"-"` // createonly (disable write permission unless it configured)
}

func generatepassword(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(bytes), err
}

func comparepassword(input_password string, real_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(real_password), []byte(input_password))
	return err == nil

}

func (u *User) ValidatePassword(value string) bool {
	return comparepassword(value, u.Password)
}

func (u *User) SetPassword() {
	password, err := generatepassword(u.Password)
	if err != nil {
		panic(err)
	}
	u.Password = password
}
