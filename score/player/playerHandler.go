package player
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"context"
	"github.com/paerarason/scorer/database"
)


func PlayerCreate() gin.HandlerFunc {
    return func(c *gin.Context){
		var ply database.Player
        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&ply);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
		}
		client, err := database.ConnectToMongo()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer client.Disconnect(context.Background())
		collection := client.Database("scorer").Collection("player")
			
		_, err = collection.InsertOne(context.Background(), ply)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message":"DOcument Saved!!"})
			return
	}


}