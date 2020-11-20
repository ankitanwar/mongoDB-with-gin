package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Person : Strcut of person
type Person struct {
	ID        primitive.ObjectID `json:"id" bson:"id"`
	Firstname string             `json:"first_name" bson:"first_name"`
	Lastname  string             `json:"last_name" bson:"last_name"`
}

var (
	client *mongo.Client
	router = gin.Default()
)

//CreateUser To create the new user
func CreateUser(c *gin.Context) {
	var newPerson = &Person{}
	if err := c.ShouldBind(newPerson); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	collection := client.Database("learning").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, _ := collection.InsertOne(ctx, newPerson)
	c.JSON(http.StatusOK, res)

}

//GetPeople : To reterive the Person from the database
func GetPeople(c *gin.Context) {
	findid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var findPerson = &Person{}
	findPerson.ID = findid
	

}

func mapURL() {
	router.POST("/users", CreateUser)
	router.GET("/users/:id", GetPeople)
}

func main() {
	fmt.Println("Starting the application.......")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	mapURL()
	router.Run(":8080")
}
