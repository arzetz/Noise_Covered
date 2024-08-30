package middleware

import (
	"net/http"
	"time"
	"github.com/noize_covered/models"
	"github.com/noize_covered/main"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var db *gorm.DB
db.AutoMigrate(&models.Session{})

func sessionMiddleware(c *gin.Context) {
	token, exists := c.Get("session_token")
	if !exists {
		token = uuid.New().String()
		c.Set("session_token", token)
		c.Set("session_expire", time.Now().Add(time.Hour))

		session := Session{
			SessionID: token,
			UserID:    c.ClientIP(),
			Token:     token,
			ExpiresAt: time.Now().Add(time.Hour),
		}

		db.Create(&session)
	} else {
		expireTime, ok := c.Get("session_expire")
		if !ok || time.Now().After(expireTime.(time.Time)) {
			c.Set("session_token", nil)
			c.Set("session_expire", nil)

			clearBasket(token)

			db.Where("token = ?", token).Delete(&Session{})

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			c.Abort()
			return
		}
	}

	c.Next()
}
