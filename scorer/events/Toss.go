package events
import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/scorer/database"
	"net/http"
)


func TOSS() gin.HandlerFunc {
    return func(c *gin.Context){

		client, err := database.connectToMongo()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer client.Disconnect(context.Background())
		collection := client.Database("scorer").Collection("match")
			
		result, err :=collection.findOneAndUpdate({_id:""},{$inc:{"toss":{
			"won":"TeamA",
			"elected":"Bat" }}})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}