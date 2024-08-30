package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/noize_covered/database"
	"github.com/noize_covered/models"
)

func CheckSession(c *gin.Context) {
	db := db.GetDB()
	db.AutoMigrate(&models.Session{})

	token, err := c.Cookie("sessionCookie")
	if err != nil {
		CreateNewSession(c)
	} else {
		if !validateSession(token) {
			var session models.Session
			if err := db.Where("token = ?", token).First(&session).Error; err != nil {
				// Обработать ошибку, если сессия не найдена
				c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
				return
			}
			c.JSON(http.StatusOK, &models.Session{})
			clearSession(token)
			CreateNewSession(c)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired, new one created"})
			c.Abort()
			return
		}
	}

	c.Next()
}

func GenerateToken() (string, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func validateSession(token string) bool {
	db := db.GetDB()
	var session models.Session
	if err := db.Where("token = ?", token).First(&session).Error; err != nil {
		return false
	}
	if time.Now().After(session.ExpiresAt) {
		return false
	}

	return true
}

func clearSession(token string) {
	// TODO: Очистить корзину, если нужно.
	db := db.GetDB()
	db.Where("token = ?", token).Delete(&models.Session{})
	db.Where("token = ?", token).Delete(&models.User{})
	db.Where("token = ?", token).Delete(&models.Basket{})
}

func CreateNewSession(c *gin.Context) {
	db := db.GetDB()
	token, _ := GenerateToken()
	c.SetCookie("sessionCookie", token, 3600, "/", "", false, true)
	user := models.User{
		Token: token,
	}
	db.Create(&user)
	session := models.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour),
	}
	db.Create(&session)
}
