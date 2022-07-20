package controllers

import (
	"net/http"
	"strconv"
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
	var usr []models.User
	
	limitPerPage := 5

	qN := ""
	qS := ""
	p := 0
	
	if qName := c.Query("name"); qName != "" {
		qN += "name LIKE " + "'%" + qName + "%'"
	}

	if qSort := c.Query("sort"); qSort != "" {
		qS += "name " + qSort
	}

	if page, _ := strconv.Atoi(c.Query("page")); page != 0 {
		p = (page-1)*limitPerPage
		db.Where(qN).Order(qS).Limit(5).Offset(p).Find(&usr)
	} else {
		db.Where(qN).Order(qS).Find(&usr)
	}

	c.JSON(http.StatusOK, gin.H{"data": usr})
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