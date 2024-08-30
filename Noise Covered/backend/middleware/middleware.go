package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/noize_covered/database"
	"github.com/noize_covered/models"
)

func SessionMiddleware(c *gin.Context, user models.User) {
	db := db.GetDB()
	db.AutoMigrate(&models.Session{})

	token, exists := c.Get("session_token")
	if !exists {
		tokenStr := uuid.New().String()
		c.Set("session_token", tokenStr)
		c.Set("session_expire", time.Now().Add(time.Hour))

		session := models.Session{
			UserID:    user.ID,
			Token:     tokenStr,
			ExpiresAt: time.Now().Add(time.Hour),
		}

		db.Create(&session)
	} else {
		expireTime, ok := c.Get("session_expire")
		if !ok || time.Now().After(expireTime.(time.Time)) {
			c.Set("session_token", nil)
			c.Set("session_expire", nil)

			//clearBasket(token)

			db.Where("token = ?", token).Delete(&models.Session{})
			db.Where("id = ?", user.ID).Delete(&models.User{})
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			c.Abort()
			return
		}
	}

	c.Next()
}
