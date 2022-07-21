package middlewares

import (
	"net/http"
	"user-api/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		err := token.TokenValid(c)
        if err != nil {
            c.String(http.StatusUnauthorized, err.Error() + " (You need to sign in)")
            c.Abort()
            return
        }
        c.Next()
	}
}