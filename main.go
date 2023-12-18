package main
import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	
	router := gin.Default()
	err := godotenv.Load(".env")	
	if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
	defer router.Run(":"+os.Getenv("PORT"))
    fmt.Println(os.Getenv("PORT"))
    
	router.POST("/login",middleware.GenerateToken())
	router.POST("api/account",account.CreateAccount())
    router.Use(middleware.JWTokenMiddlerware)
	//Bunch of APIS for account management 
	accounts:=router.Group("api/account")
	{
		
	}

    

}