package controllers

import (
	"fuel-api/database"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Register(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Username and password are required",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id
	`

	err = database.DB.QueryRow(
		query,
		user.Username,
		string(hashedPassword),
	).Scan(&user.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = ""

	c.JSON(200, gin.H{
		"message": "User registered successfully",
		"data":    user,
	})
}

func Login(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Username and password are required",
		})
		return
	}

	var storedUser User

	query := `
		SELECT id, username, password
		FROM users
		WHERE username = $1
	`

	err := database.DB.QueryRow(
		query,
		user.Username,
	).Scan(
		&storedUser.ID,
		&storedUser.Username,
		&storedUser.Password,
	)

	if err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(storedUser.Password),
		[]byte(user.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}