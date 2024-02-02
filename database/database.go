package database

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/utils"
)

type albums struct {
	Position int    `json:"position"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
}

type albumsList []albums

func ParseAlbumsJSON() (obj albumsList) {
	projectRoot := utils.RootDir()
	databaseFilePath := "/database/albums_2024.json"
	data, err := os.ReadFile(projectRoot + databaseFilePath)
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}
