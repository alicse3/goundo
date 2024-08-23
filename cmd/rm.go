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

const (
	TypeFile      = "FILE"
	TypeDirectory = "DIRECTORY"
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
			srcPath := args[ind]

			// TODO: Parse rm flags
			if strings.HasPrefix(srcPath, "-") {
				fmt.Printf("rm: unrecognized option '%s'\n", srcPath)
				return
			}

			// Check file type
			fi, err := os.Stat(srcPath)
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

			// Construct destination path
			dstPath := filepath.Join(dirToMove, filepath.Base(srcPath))

			fileType := TypeFile

			// Check and backup files or directories
			if fi.IsDir() {
				fileType = TypeDirectory

				// Move directory to the timestamp dir
				if err := util.MoveDirectory(srcPath, dstPath); err != nil {
					fmt.Printf("error moving %s directory to backups dir: %v\n", srcPath, err)
					return
				}
			} else {
				// Move file to the timestamp dir
				if err := util.MoveFile(srcPath, dstPath); err != nil {
					fmt.Printf("error moving %s file to backups dir: %v\n", srcPath, err)
					return
				}
			}

			// Get absolute path
			absPath, err := filepath.Abs(srcPath)
			if err != nil {
				fmt.Printf("error getting the absolute path: %v\n", err)
				return
			}

			// Track info in DB
			if err := db.Insert(absPath, dstPath, fileType); err != nil {
				fmt.Printf("error inserting info in db: %v\n", err)
				return
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
