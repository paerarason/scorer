package player
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/paerarason/scorer/database"
)


func PlayerCreate() gin.HandlerFunc {
    return func(c *gin.Context){
		var ply player
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
		collection := client.Database("scorer").Collection("player")
			
		result, err := collection.InsertOne(context.Background(), ply)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
}

