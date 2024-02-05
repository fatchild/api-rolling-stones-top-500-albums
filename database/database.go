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
	databaseFile := "albums_2024.json"

	fmt.Println(utils.RootDir() + "/" + databaseFile)
	data, err := os.ReadFile(databaseFile)
	if err != nil {
		fmt.Print(err)
		fmt.Print("Check for the file in the root (with respect to utils")
		data, err = os.ReadFile(utils.RootDir() + "/" + databaseFile)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}
