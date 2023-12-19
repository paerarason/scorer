package main
import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/paerarason/scorer/score/player"
	"github.com/paerarason/scorer/score/team"
	"github.com/paerarason/scorer/score/match"
	"github.com/paerarason/scorer/middleware"
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
	players:=router.Group("api/player")
	{
		players.POST("/",player.PlayerCreate())
	}

	matchs:=router.Group("api/match")
	{
		matchs.POST("/",match.MatchCreate())
	}
    
	teams:=router.Group("api/team")
	{
		teams.POST("/",team.TeamCreate())
	}
    

}