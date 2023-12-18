package team
import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/scorer/database"
	"net/http"
)


func TeamCreate() gin.HandlerFunc {
    return func(c *gin.Context){
		var ply Team
        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&ply);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
		}

		client, err := database.connectToMongo()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer client.Disconnect(context.Background())
		collection := client.Database("scorer").Collection("Team")
			
		result, err := collection.InsertOne(context.Background(), ply)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}