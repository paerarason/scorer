package database

type Player struct {
	Name  string `json:"name" binding:"required"`
	role  string `json:"role" binding:"required" "oneof=batsman bowler allrounder"`
	Age   int    `json:"age" binding:"required"`
}

type Team struct {
	Name  string     `json:"name" binding:"required"`
	players []Player `json:"players`
}

type Match struct {
	teamA    string  `json:"TeamA" binding:"required"`
	teamB    string  `json:"TeamB" binding:"required"`
	overs    int     `json:"overs" binding:"required"`
	players  int     `json:"players" binding:"required"`
}





 


