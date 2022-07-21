package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"user-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr := models.User{}

	usr.Email = input.Email
	usr.Password = input.Password

	token, err := models.LoginCheck(usr.Email, usr.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	user := map[string]string{
		"email": usr.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})
}

func GetProfile(c *gin.Context){
	// get current user
	usr, err := models.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}

	usr.HidePassword()

	c.JSON(http.StatusOK, gin.H{"data": usr})
}

func GetAllUser(c *gin.Context) {
	_, errAuth := models.GetCurrentUser(c)
	if errAuth != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	
	qN := ""
	qS := ""
	limitPerPage := 5
	
	if qName := c.Query("name"); qName != "" {
		qN += "name LIKE " + "'%" + qName + "%'"
	}

	if qSort := c.Query("sort"); qSort != "" {
		qS += "name " + qSort
	}

	if page, _ := strconv.Atoi(c.Query("page")); page != 0 {
		p := (page-1)*limitPerPage
		db.Where(qN).Order(qS).Limit(limitPerPage).Offset(p).Find(&user)
	} else {
		db.Where(qN).Order(qS).Find(&user)
	}

	listUser := []map[string]string{}

	for _, usr := range(user) {
		userData := map[string]string{
			"id": strconv.Itoa(int(usr.ID)),
			"name": usr.Name,
			"email": usr.Email,
			"password": "******",
			"role": usr.Role,
		}
		listUser = append(listUser, userData)
	}

	c.JSON(http.StatusOK, gin.H{"data": listUser})
}

func GetUserById(c *gin.Context) {
	_, errAuth := models.GetCurrentUser(c)
	if errAuth != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You need to sign in"})
		return
	}

    db := c.MustGet("db").(*gorm.DB)
    var user models.User

    if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}
