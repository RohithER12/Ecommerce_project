package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			tokenCookie, err := c.Cookie("token")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header or token cookie missing"})
				return
			}
			tokenString = tokenCookie
		} else {
			tokenString = tokenString[len("Bearer "):]
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("email", claims.Email)
		c.SetCookie("token", tokenString, 3600, "/", "70off.online", true, true)

		userIDString := strconv.Itoa(int(claims.UserID))
		adminIDString := strconv.Itoa(int(claims.AdminID))

		c.SetCookie("userID", userIDString, 3600, "/", "70off.online", true, false)
		c.SetCookie("adminID", adminIDString, 3600, "/", "70off.online", true, false)

		c.Next()
	}
}
