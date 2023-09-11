package controllers

import (
	"context"
	"fmt"
	"github.com/Badiru27/hngx-stage-2/configs"
	"github.com/Badiru27/hngx-stage-2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := c.ShouldBindJSON(&user); err != nil {

		fmt.Printf("Error: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request: Name is required",
			"error":   err.Error(),
		})

		return
	}

	if validateErr := validate.Struct(&user); validateErr != nil {

		fmt.Printf("Error validating: %v", validateErr)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error validating",
			"error":   validateErr.Error(),
		})

		return
	}

	filter := bson.M{"name": user.Name}
	existingUser := userCollection.FindOne(ctx, filter)

	if existingUser.Err() == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Name already exist",
		})

		return
	}

	newUser := models.User{
		Name: user.Name,
		Id: primitive.NewObjectID(),
	}

	_, err := userCollection.InsertOne(ctx, newUser)

	if err != nil {
		fmt.Printf("Error creating user %v\n", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error validating",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   newUser.Id,
		"name": newUser.Name,
	})
}

func GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	fmt.Printf("err: %v", err)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   user.Id,
		"name": user.Name,
	})

}

func UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"id": objId}
	existingUser := userCollection.FindOne(ctx, filter)

	if existingUser.Err() != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User does not exist",
		})

		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {

		fmt.Printf("Error: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request: Name is required",
			"error":   err.Error(),
		})

		return
	}

	if validateErr := validate.Struct(&user); validateErr != nil {

		fmt.Printf("Error validating: %v", validateErr)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error validating",
			"error":   validateErr.Error(),
		})

		return
	}

	update := bson.M{"name": user.Name}

	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		fmt.Printf("Error updating user: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error updating user",
			"error":   err,
		})

		return
	}

	var updateUser models.User

	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updateUser)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "User does not exist",
			})

			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"id":   updateUser.Id,
		"name": updateUser.Name,
	})

}

func DeleteUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"id": objId}
	existingUser := userCollection.FindOne(ctx, filter)

	if existingUser.Err() != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User does not exist",
		})
		return
	}

	_, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})

}
