package markdown

import (
	"context"
	"net/http"
	"time"

	"github.com/Ik-cyber/markdown-ai-app/internal/database"
	"github.com/gin-gonic/gin"
)

type Markdown struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Struct for input
type CreateMarkdownInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Create a markdown
func CreateMarkdown(c *gin.Context) {
	var input CreateMarkdownInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetInt("userID")

	query := `INSERT INTO markdowns (user_id, title, content) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	var markdown Markdown

	err := database.DB.QueryRow(context.Background(), query, userID, input.Title, input.Content).
		Scan(&markdown.ID, &markdown.CreatedAt, &markdown.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create markdown"})
		return
	}

	markdown.UserID = userID
	markdown.Title = input.Title
	markdown.Content = input.Content

	c.JSON(http.StatusCreated, markdown)
}

// List markdowns for the logged-in user
func ListMarkdowns(c *gin.Context) {
	userID := c.GetInt("userID")

	query := `SELECT id, user_id, title, content, created_at, updated_at FROM markdowns WHERE user_id=$1`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch markdowns"})
		return
	}
	defer rows.Close()

	var markdowns []Markdown
	for rows.Next() {
		var m Markdown
		if err := rows.Scan(&m.ID, &m.UserID, &m.Title, &m.Content, &m.CreatedAt, &m.UpdatedAt); err == nil {
			markdowns = append(markdowns, m)
		}
	}

	c.JSON(http.StatusOK, markdowns)
}

// Get a specific markdown by ID
func GetMarkdown(c *gin.Context) {
	userID := c.GetInt("userID")
	id := c.Param("id")

	query := `SELECT id, user_id, title, content, created_at, updated_at FROM markdowns WHERE id=$1 AND user_id=$2`
	var m Markdown

	err := database.DB.QueryRow(context.Background(), query, id, userID).
		Scan(&m.ID, &m.UserID, &m.Title, &m.Content, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Markdown not found"})
		return
	}

	c.JSON(http.StatusOK, m)
}

// Update a markdown
func UpdateMarkdown(c *gin.Context) {
	userID := c.GetInt("userID")
	id := c.Param("id")

	var input CreateMarkdownInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `UPDATE markdowns SET title=$1, content=$2, updated_at=$3 WHERE id=$4 AND user_id=$5`
	result, err := database.DB.Exec(context.Background(), query, input.Title, input.Content, time.Now(), id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update markdown"})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Markdown not found or not authorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Markdown updated successfully"})
}

// Delete a markdown
func DeleteMarkdown(c *gin.Context) {
	userID := c.GetInt("userID")
	id := c.Param("id")

	query := `DELETE FROM markdowns WHERE id=$1 AND user_id=$2`
	result, err := database.DB.Exec(context.Background(), query, id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete markdown"})
		return
	}

	if result.RowsAffected() == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Markdown not found or not authorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Markdown deleted successfully"})
}
