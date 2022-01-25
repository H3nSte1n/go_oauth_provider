package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"net/http"
	"oauth_provider/models"
	"oauth_provider/db"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(user)
		fmt.Println("USERASDASD")
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	currentTime := time.Now()
	user.CreatedAt = &currentTime
	id, err := db.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	
	c.JSON(200, gin.H{"id": id})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	currentTime := time.Now()
	user.UpdatedAt = &currentTime
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	id, err := db.UpdateUser(docID, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func DeleteUser(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user_id, err := db.DeleteUser(docID)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": user_id,
	})
}

func GetUser(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user, err := db.GetUser(docID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUsers(c *gin.Context) {
	users, err := db.GetUsers()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}