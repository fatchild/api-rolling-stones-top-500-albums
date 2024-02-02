package getAlbums

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type albums struct {
	Position int    `json:"position"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
}

type albumsList []albums

func parseAlbumsJSON() (obj albumsList) {
	data, err := ioutil.ReadFile("./database/albums_2024.json")
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func GetAlbums(c *gin.Context) {
	obj := parseAlbumsJSON()
	c.IndentedJSON(http.StatusOK, obj)
}
