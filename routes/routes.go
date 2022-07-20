package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"user-api/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })
    r.GET("/users", controllers.GetAllUser)
    r.POST("/users", controllers.CreateUser)
    r.GET("/users/:id", controllers.GetUserById)
    r.PATCH("/users/:id", controllers.UpdateUser)
    r.DELETE("users/:id", controllers.DeleteUser)

    return r
}