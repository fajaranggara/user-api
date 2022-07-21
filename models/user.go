package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"user-api/utils/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name    	string    `json:"name" gorm:"not null"`
	Email		string    `json:"email" gorm:"not null; unique"`
	Password	string	  `json:"password" gorm:"not null"`
	Role		string	  `json:"role" gorm:"default:user"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (usr *User) HidePassword(){
	usr.Password = "******"
}

func GetCurrentUser(c *gin.Context) (User, error) {
    db := c.MustGet("db").(*gorm.DB)
	var usr User

	uid, err := token.ExtractTokenID(c)
	if err != nil {
		return usr, errors.New("Failed while extract token by id")
	}

    if err := db.Model(User{}).First(&usr, uid).Error; err != nil {
		return usr, errors.New("User not found")
	}
	
    return usr, err
}


func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email, password string, db *gorm.DB) (string, error) {
	var err error

	usr := User{}

	err = db.Model(User{}).Where("email = ?", email).Take(&usr).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, usr.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(usr.ID, usr.Role)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (usr *User) SaveUser(db *gorm.DB) (*User, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}

	usr.Password = string(hashedPassword)
	usr.Email = html.EscapeString(strings.TrimSpace(usr.Email))

	var err error = db.Create(&usr).Error
	if err != nil {
		return &User{}, err
	}

	return usr, nil
}


func (usr *User) SaveNewPassword(db *gorm.DB) (*User, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}

	usr.Password = string(hashedPassword)

	var err error = db.Save(&usr).Error
	if err != nil {
		return &User{}, err
	}

	return usr, nil
}
