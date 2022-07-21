package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"user-api/controllers"
	"user-api/middlewares"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.POST("/login", controllers.Login)

    auth := r.Use(middlewares.JwtAuthMiddleware())
    // admin authorization
    auth.POST("/admin", controllers.RegistAdmin)
	auth.POST("/users", controllers.CreateUser)
    auth.PATCH("/users/:id", controllers.UpdateUser)
    auth.DELETE("users/:id", controllers.DeleteUser)

    // user & admin authorization
    auth.GET("/profile", controllers.GetProfile)
    auth.GET("/users", controllers.GetAllUser)
    auth.GET("/users/:id", controllers.GetUserById)

    return r
}