package getAlbumByID

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
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

func GetAlbumByID(c *gin.Context) {
	obj := parseAlbumsJSON()

	idString := c.Param("position")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	for _, a := range obj {
		if a.Position == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
