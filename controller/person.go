package controller

import (
	"context"
	"fmt"
	"log"
	"mongo/configs"
	"mongo/models"
	"mongo/mongolib"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

//var DB *mongo.Database

type PersonAge struct {
	Age int `json:"age"`
}

var DB *mongo.Database
var config *configs.Config

/*

	WE CANNOT CREATE GLOBAL CONTEXT !!!!

*/

func init() {
	//config, _ := configs.Configs()
	allConfig, _ := configs.Configs()
	config = allConfig
	mongoDataStore, err := mongolib.NewMongoDataStore(&config.Mongo)

	if err != nil {
		log.Fatal("Mongo connection error", err)
	} else {
		fmt.Println("successfully connected to mongoDB")
	}

	DB = mongoDataStore.DB

}

// get all people
func GetPeople(c *gin.Context) {
	fmt.Println(config.Mongo)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	people, err := mongolib.GetPeople(ctx, DB.Collection(config.Mongo.Collection))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code": -1,
			"error":      err.Error(),
			"sucess":     false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"error":      "",
		"success":    true,
		"people":     people,
	})
}

// get specific person by name
func GetPerson(c *gin.Context) {
	firstname := c.Param("firstname")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	person, err := mongolib.GetPerson(ctx, DB.Collection(config.Mongo.Collection), firstname)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code": -1,
			"error":      err.Error(),
			"sucess":     false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"error":      "",
		"success":    true,
		"person":     person,
	})
}

// createPerson
func CreatePerson(c *gin.Context) {
	var person models.Person

	err := c.ShouldBindJSON(&person)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": -1,
			"error":      err.Error(),
			"sucess":     false,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mongolib.CreatePerson(ctx, DB.Collection(config.Mongo.Collection), person)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code": -1,
			"error":      err.Error(),
			"success":    false,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"error_code": 0,
		"error":      "",
		"success":    true,
		"person":     result,
	})
}

func UpdateCustomerAge(c *gin.Context) {
	firstname := c.Param("firstname")
	var age PersonAge
	err := c.ShouldBindJSON(&age)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": -1,
			"error":      err.Error(),
			"success":    false,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mongolib.UpdatePerson(ctx, DB.Collection(config.Mongo.Collection), firstname, age.Age)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code": -1,
			"error":      err.Error(),
			"success":    false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"error":      "",
		"success":    true,
		"result":     result,
	})
}

// delete person

func DeletePerson(c *gin.Context) {
	firstname := c.Param("firstname")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mongolib.DeletePerson(ctx, DB.Collection(config.Mongo.Collection), firstname)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code": -1,
			"error":      err.Error(),
			"success":    false,
		})
	}

	c.JSON(http.StatusNoContent, gin.H{
		"success": true,
		"result":  result,
	})
}
