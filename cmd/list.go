package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/alicse3/goundo/internal/database"
)

// listHandler lists backups from database.
func listHandler() {
	// Get config
	cfg := getConfig()
	if cfg == nil {
		fmt.Println("error getting config")
		return
	}

	// Initialize the DB
	db, err := database.NewDBHandler(cfg.SqliteDBPath)
	if err != nil {
		fmt.Printf("error initializing the db: %v\n", err)
		return
	}

	// Get backups from DB
	backups, err := db.List()
	if err != nil {
		fmt.Printf("error getting backups from the db: %v\n", err)
		return
	}

	// Check if there are backups
	if len(backups) == 0 {
		fmt.Println("there are no backups to list")
		return
	}

	// Print backups
	for index := range backups {
		data, err := json.MarshalIndent(backups[index], "", "  ")
		if err != nil {
			fmt.Printf("error marshalling backup: %v\n", err)
			return
		}

		fmt.Println(string(data))
	}
}
