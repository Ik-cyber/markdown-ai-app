package user

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ik-cyber/markdown-ai-app/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Hashed
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Register a new user
func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	// Insert user into the database
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err = database.DB.QueryRow(context.Background(), query, user.Username, string(hashedPassword)).Scan(&user.ID)
	if err != nil {
		log.Println("Error inserting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user_id": user.ID})
}

// Login a user
func Login(c *gin.Context) {
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var storedUser User
	query := "SELECT id, username, password FROM users WHERE username=$1"
	err := database.DB.QueryRow(context.Background(), query, input.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
