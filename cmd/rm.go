package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/alicse3/goundo/internal/config"
	"github.com/alicse3/goundo/internal/database"
	"github.com/alicse3/goundo/internal/util"
)

// rmHandler handles the rm commands.
func rmHandler() {
	// Get rm args
	args := os.Args[2:]

	// Process if there are args
	if len(args) > 0 {
		// Get config
		cfg, err := config.GetConfig()
		if err != nil {
			fmt.Printf("error getting the config data: %v\n", err)
			return
		}

		// Initialize the DB
		dbPath := cfg.AppPath + string(filepath.Separator) + "backups.db"
		db, err := database.NewDBHandler(dbPath)
		if err != nil {
			fmt.Printf("error initializing the db: %v\n", err)
			return
		}

		// Parse rm args
		for ind := range args {
			// TODO: Parse rm flags
			if strings.HasPrefix(args[ind], "-") {
				fmt.Printf("rm: unrecognized option '%s'\n", args[ind])
				return
			}

			// Check file type
			fi, err := os.Stat(args[ind])
			if err != nil {
				fmt.Printf("error getting the file stat: %v\n", err)
				return
			}

			// Create a dir with current timestamp
			dirToMove, err := createDirWithTimestamp(cfg.BackupsPath)
			if err != nil {
				fmt.Printf("error creating dir with timestamp: %v\n", err)
				return
			}

			// Check and backup files or directories
			if fi.IsDir() {
				// TODO: Backup dir
			} else {
				// Move file to the timestamp dir
				dstPath := dirToMove + string(filepath.Separator) + args[ind]
				if err := util.MoveFile(args[ind], dstPath); err != nil {
					fmt.Printf("error moving %s to backups dir: %v\n", args[ind], err)
					return
				}

				absPath, err := filepath.Abs(args[ind])
				if err != nil {
					fmt.Printf("error getting the absolute file path: %v\n", err)
					return
				}

				// Track info in DB
				if err := db.Insert(absPath, dstPath); err != nil {
					fmt.Printf("error inserting file info in db: %v\n", err)
					return
				}
			}
		}
	} else {
		println("rm: missing operand")
	}
}

// createDirWithTimestamp creates a directory with current timestamp
func createDirWithTimestamp(backupsPath string) (dirToMove string, err error) {
	dirToMove = backupsPath + string(filepath.Separator) + generateUniqueTimestamp()
	if err = os.Mkdir(dirToMove, os.ModeDir|0755); err != nil {
		return "", err
	}

	return dirToMove, nil
}

var mu sync.Mutex

// generateUniqueTimestamp generates the unique timestamp
func generateUniqueTimestamp() string {
	mu.Lock()
	defer mu.Unlock()

	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
