package controllers

import (
	"net/http"
	"time"
	"user-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
    var input createUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{Name: input.Name, Email: input.Email}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&user)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUserById(c *gin.Context) {
    var user models.User

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    var input createUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.User
    updatedInput.Name = input.Name
    updatedInput.Email = input.Email
    updatedInput.UpdatedAt = time.Now()

    db.Model(&user).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var user models.User
    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&user)

    c.JSON(http.StatusOK, gin.H{"data": true})
}