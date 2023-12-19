package admin 
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/paerarason/scorer/database"
)

type Account struct {
    Email string               `json:"email"`
    Username string            `json:"username"`
    Password string            `json:"password"`
    ConfirmPassword string     `json:"confirmPassword"`
}


func CreateAccount() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.connectToMongo()      
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }

        defer db.Close()        
        var acc Account
        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&acc);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        
        var accid int
        bytes, err := bcrypt.GenerateFromPassword([]byte(acc.Password), 14)
        query := `INSERT INTO account (username, password, email, balance)
          VALUES ($1, $2, $3, $4) RETURNING id`
        
        dberr := db.QueryRow(query, acc.Username, bytes, acc.Email, 5000).Scan(&accid)
        //Handlinf while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"account ID ":accid}) 
    }
}