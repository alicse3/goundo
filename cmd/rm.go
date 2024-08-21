package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/alicse3/goundo/internal/config"
	"github.com/alicse3/goundo/internal/util"
)

// rmHandler handles the rm commands.
func rmHandler() {
	// Get rm args
	args := os.Args[2:]

	// Process if there are args
	if len(args) > 0 {
		// TODO: Parse rm args

		// Check file type
		fi, err := os.Stat(args[0])
		if err != nil {
			fmt.Printf("error getting the file stat: %v\n", err)
			return
		}

		// Check and back file or directories
		if fi.IsDir() {
			// TODO: Backup dir
		} else {
			// Get config
			cfg, err := config.GetConfig()
			if err != nil {
				fmt.Printf("error getting the config data: %v\n", err)
				return
			}

			// Create a dir with current timestamp
			dirToMove := cfg.BackupsPath + string(filepath.Separator) + generateUniqueTimestamp()
			if err := os.Mkdir(dirToMove, os.ModeDir|0755); err != nil {
				fmt.Printf("error getting the config data: %v\n", err)
				return
			}

			// Move file to the timestamp dir
			dstPath := dirToMove + string(filepath.Separator) + args[0]
			if err := util.MoveFile(args[0], dstPath); err != nil {
				fmt.Printf("error moving %s to backups dir: %v\n", args[0], err)
				return
			}

			// TODO: Undo backup in case of failures
			// TODO: Track info in DB in case of success
		}
	} else {
		println("rm: missing operand")
	}
}

var mu sync.Mutex

// generateUniqueTimestamp generates the unique timestamp
func generateUniqueTimestamp() string {
	mu.Lock()
	defer mu.Unlock()

	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
