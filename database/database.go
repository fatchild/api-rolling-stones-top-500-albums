package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type albums struct {
	Position int    `json:"position"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
}

type albumsList []albums

func ParseAlbumsJSON() (obj albumsList) {
	databaseFile := "albums_2024.json"

	data, err := os.ReadFile(databaseFile)
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}
