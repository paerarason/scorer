package database
import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-gonic/gin"
)



func connectToMongo() (*mongo.Client, error) {
	// Replace with your actual connection details
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getUsers(c *gin.Context) {
	client, err := connectToMongo()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("users").Collection("users")
	users := []User{}
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &users)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func createUser(c *gin.Context) {
	client, err := connectToMongo()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("users").Collection("users")
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"id": result.InsertedID.(string)})
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	client, err := connectToMongo()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("users").Collection("users")
	var user User
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	client, err := connectToMongo()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	
