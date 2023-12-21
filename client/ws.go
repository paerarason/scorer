package client
import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
    // Check origin for production environments
    CheckOrigin: func(r *http.Request) bool { return true },
}

func ScoreHandler(conn *websocket.Conn){
	for {
        select {
        case matchID <-match:
			score, userDetails := fetchUpdatedUserDetails(userID)
        }
    }
}