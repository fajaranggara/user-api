package controllers

import (
	"net/http"
	"time"
	"user-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type adminRegisterInput struct {
    Name string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type createUserInput struct {
	Name string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type updateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string  `json:"password"`
}


func RegistAdmin(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input adminRegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    usr := models.User{}

    usr.Name = input.Name
    usr.Email = input.Email
    usr.Password = input.Password
    usr.Role = "admin"

    _, err := usr.SaveUser(db)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := map[string]string{
        "name": usr.Name,
        "email": usr.Email,
        "role": usr.Role,
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registration success. You are Admin now", "user": user})
}


func CreateUser(c *gin.Context) {
    currentUser, errAuth := models.GetCurrentUser(c)
    // authentication check
	if errAuth != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}
    // authorization check
    if currentUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "You are not Admin"})
        return
	}

    db := c.MustGet("db").(*gorm.DB)

    var input createUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    usr := models.User{}

    usr.Name = input.Name
    usr.Email = input.Email
    usr.Password = input.Password
    usr.Role = "user"

    _, err := usr.SaveUser(db)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := map[string]string{
        "name": input.Name,
        "email": input.Email,
    }

    c.JSON(http.StatusOK, gin.H{"message": "Create user success", "user": user})
}


func UpdateUser(c *gin.Context) {
    currentUser, errAuth := models.GetCurrentUser(c)
    // authentication check
	if errAuth != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}
    // authorization check
    if currentUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "You are not Admin"})
        return
	}

    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    var input updateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedUser models.User
    updatedUser.Name = input.Name
    updatedUser.Email = input.Email
    updatedUser.Password = input.Password
    updatedUser.UpdatedAt = time.Now()

    db.Model(&user).Updates(updatedUser)

    c.JSON(http.StatusOK, gin.H{"user": user})
}


func DeleteUser(c *gin.Context) {
    currentUser, errAuth := models.GetCurrentUser(c)
    // authentication check
	if errAuth != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}
    // autorization check
    if currentUser.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"forbidden": "You are not Admin"})
        return
	}

    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&user)

    c.JSON(http.StatusOK, gin.H{"message": user.Name + " deleted"})
}

