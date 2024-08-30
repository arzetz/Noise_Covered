package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/noize_covered/database"
	md "github.com/noize_covered/middleware"
	"github.com/noize_covered/models"
)

func main() {
	router := gin.Default()
	db := db.GetDB()
	db.AutoMigrate(&models.User{}, &models.Composition{}, &models.Basket{}, &models.Order{}, &models.Session{})

	router.POST("/", func(c *gin.Context) {
		md.CreateNewSession(c)
		md.CheckSession(c)
	})

	router.GET("/", func(c *gin.Context) {
		token, _ := c.Cookie("sessionCookie")
		c.JSON(http.StatusOK, token)
	})

	router.GET("/people", func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			// Обработка ошибки, если элемент не найден
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		// Ответ в формате JSON
		c.JSON(http.StatusOK, users)
	})

	router.POST("/order", func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
		result := db.Create(&user)
		if result.Error != nil {
			fmt.Print("Всё ок")
		}
	})

	router.GET("/basket", func(c *gin.Context) {
		var basket []models.Basket
		cookie, _ := c.Cookie("sessionCookie")
		db.Where("token = ?", cookie).Find(&basket)
		c.JSON(http.StatusOK, basket)
	})

	router.GET("/compositions", func(c *gin.Context) {
		var items []models.Composition

		if err := db.Find(&items).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		c.JSON(http.StatusOK, items)
	})

	router.GET("/compositions/:compositionID", func(c *gin.Context) {
		itemID := c.Param("compositionID")
		var item models.Composition

		// Запрос к базе данных по itemId
		if err := db.First(item, itemID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		c.JSON(http.StatusOK, item)
	})

	router.POST("/compositions", func(c *gin.Context) {
		var composition models.Composition
		if err := c.BindJSON(&composition); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		db.Create(&composition)
	})

	router.POST("/compositions/:compositionID", func(c *gin.Context) {
		compositionID := c.Param("compositionID")
		var user models.User
		var composition models.Composition
		cookie, _ := c.Cookie("sessionCookie")
		if err := db.First(&composition, compositionID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Composition not found"})
			return
		}
		if err := db.Where("token = ?", cookie).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Composition not found"})
			return
		}
		basket := models.Basket{
			CompositionID: composition.ID, //  Заполни  поле  CompositionID
			Name:          composition.Name,
			Price:         composition.Price,
			Quantity:      1,
			Token:         cookie,
			UserID:        user.ID,
		}
		db.Create(&basket)
	})

	router.Run(":8080")
}
