package model

import (
	"github.com/jinzhu/gorm"
	"viki/common"
)

// "github.com/lib/pq"

type User struct {
	gorm.Model
	Nickname string `gorm:"column: nickname"`
	Username string `gorm:"column: username"`
	Password string `gorm:"column: password"`

	Master bool `gorm:"DEFAULT: false"`
}

func FindUser(condition interface{}) (User, error) {
	db := common.GetDB()
	var user User
	err := db.Where(condition).First(&user).Error
	return user, err
}

// func (u *User) setPassword(password string) error {
// 	if len(password) == 0 {
// 		return errors.New("password should not be empty!")
// 	}
// 	bytePassword := []byte(password)
// 	// Make sure the second param `bcrypt generator cost` between [4, 32)
// 	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
// 	u.PasswordHash = string(passwordHash)
// 	return nil
// }

// func (u *User) checkPassword(password string) error {
// 	bytePassword := []byte(password)
// 	byteHashedPassword := []byte(u.PasswordHash)
// 	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
// }