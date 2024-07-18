package controllers

import (
	"fmt"
	"goblogart/inits"
	"goblogart/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}

	result := inits.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}

	var user models.User

	result := inits.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(401, gin.H{
			"error": "unauthorised",
		})
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(500, gin.H{
			"error": "error signing token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorisation", tokenString, 3600*24*30, "", "localhost", false, true)
}

func GetUsers(c *gin.Context) {
	var users []models.User

	err := inits.DB.Model(&models.User{}).Preload("Posts").Find(&users).Error

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": "error getting users",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": users,
	})
}

func Validate(c *gin.Context) {
	user, err := c.Get("user")

	if err {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": "You are logged in!",
		"user": user,
	})
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorisation", "", -1, "", "localhost", false, true)
	c.JSON(200, gin.H{
		"data": "You are logged out!",
	})
}
