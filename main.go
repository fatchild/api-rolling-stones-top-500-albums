package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strconv"
)

func main() {
    router := gin.Default()

    router.LoadHTMLGlob("templates/*")
    router.GET("/getAlbumList", getAlbums)
	router.GET("/getAlbumAtPosition/:position", getAlbumByID)
    router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

    router.Run("localhost:8080")
}

type albums struct {
    Position int `json:"position"`
    Artist string `json:"artist"`
    Album string `json:"album"`
}

type albumsList []albums

func parseAlbumsJSON() (obj albumsList) {
    data, err := ioutil.ReadFile("./albums_2024.json")
    if err != nil {
      fmt.Print(err)
    }

    err = json.Unmarshal(data, &obj)
    if err != nil {
        fmt.Println("error:", err)
    }

    return obj
}

func getAlbums(c *gin.Context) {
    obj := parseAlbumsJSON()
    c.IndentedJSON(http.StatusOK, obj)
}

func getAlbumByID(c *gin.Context) {
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