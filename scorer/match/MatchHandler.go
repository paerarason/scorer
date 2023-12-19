package scorer
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"context"
	"github.com/paerarason/scorer/database"
)





func MatchCreate()gin.HandlerFunc {
    return func(c *gin.Context){
		var mth database.Match
        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&mth);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
		}
	
	client, err := database.ConnectToMongo()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("scorer").Collection("match")
	result, err := collection.InsertOne(context.Background(), mth)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
}