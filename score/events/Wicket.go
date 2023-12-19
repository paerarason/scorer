package events
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required"`
    Age   int    `json:"age" binding:"required"`
    Choice string `json:"choice" binding:"required,eq=1|eq=2"`
}

type Wicket struct {
    batter  string `json:"batter_id" binding:"required"`
    mode string `json:"mode" binding:"required "oneof=runout hitwicket caught&bowled caught bowled lbw stumped"`
    bowler   int    `json:"bowler_id" binding:"required"`
    fielder string `json:"fielder_id"`
}


func WicketHandler() gin.HandlerFunc {
    return func(c *gin.Context){
		var wic Wicket
        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&wic);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
	    

		
		
		if wic.mode=="bowled"{

		}else if wic.mode=="lbw"{

		}else if wic.mode=="hitwicket"{
			
		}else if wic.mode=="caught"{
			
		}else if wic.mode=="caught&bowled"{
		 
		}else if wic.mode=="stumped"{
			
		}else if wic.mode=="runout"{
			
		}else if wic.mode=="man"{
			
		}
	}
}




func WicketTransaction(){

} 