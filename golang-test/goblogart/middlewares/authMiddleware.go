package middlewares

import (
	"fmt"
	"goblogart/inits"
	"goblogart/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorisation")

	if err != nil {
		c.JSON(401, gin.H{
			"error": "unauthorised",
		})

		c.AbortWithStatus(http.StatusUnauthorized)
		// -> ensures that no further processing or middleware execution occurs

		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(401, gin.H{
				"error": "unauthorised",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		inits.DB.First(&user, int(claims["id"].(float64)))
		if user.ID == 0 {
			c.JSON(401, gin.H{
				"error": "unauthorised",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
