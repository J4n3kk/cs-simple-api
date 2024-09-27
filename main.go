/*
I want to create simple api with post and get.
/players GET-> returns all players //func getPlayers
/players POST-> add player to db // func postPlayer
/players/:id GET-> return single player // func getPlayerByID

The data will be stored in mysql DB hltv, TABLE players
| playerID | nickname | kd |
*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var mydb *sql.DB

type playerModel struct {
	ID       int64   `json:"id"`
	Fullname string  `json:"fullname"`
	Nickname string  `json:"nickname"`
	Kd       float32 `json:"k/d"`
	Team     string  `json:"team"`
}

func getPlayersAll(c *gin.Context) {
	var plrs []playerModel

	rows, err := mydb.Query("SELECT * FROM players")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var plr playerModel
		if err := rows.Scan(&plr.ID, &plr.Nickname, &plr.Fullname, &plr.Kd, &plr.Team); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unable to map db row"})
			return
		}
		plrs = append(plrs, plr)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("rows error: %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, plrs)
}

func postPlayer(c *gin.Context) {
	var p playerModel
	if err := c.BindJSON(&p); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "bind failed", "err": err})
		return
	}
	if res, err := mydb.Exec("INSERT INTO players (fullname, nickname, kd, team) VALUES (?, ?, ?, ?)", p.Fullname, p.Nickname, p.Kd, p.Team); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": "Insert failed", "err": err})
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBADDR"), //"127.0.0.1:3306"
		DBName: "playerdata",
	}
	var err error
	mydb, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/players", getPlayersAll)
	router.POST("/players", postPlayer)
	router.Run("0.0.0.0:8081")
}
