package middleware

import (
	"fmt"
	"goauth/initializers"
	"goauth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context){
	// Fetching the cookie
	tokenString, err := c.Cookie("token")
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Login to access this route",
		})
	}

	// Parsing and validating token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("SECRET")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Checking the expiry
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token expired. Relog to authorize.",
			})
		}

		// Fetching the user
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "User not found.",
			})
		}

		c.Set("user", user)

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid access token",
		})
	}

	// Continuing
	c.Next()
}