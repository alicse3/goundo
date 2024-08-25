package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/alicse3/goundo/internal/database"
	"github.com/alicse3/goundo/internal/util"
)

// restoreHandler handles the restore command.
func restoreHandler(backupIds string) {
	// Check if there are any backup ids provided
	if backupIds == "" {
		fmt.Println("please specify backup ids to restore")
		return
	}

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

	// Restore from backups using the backup ids
	ids := strings.Split(backupIds, ",")
	for index := range ids {
		// Get backup by id
		backup, err := db.GetById(ids[index])
		if err != nil {
			fmt.Printf("error getting the backup data for id %s, err: %v\n", ids[index], err)
			return
		}

		// Check if the status is backed up
		if backup.Status != database.StatusBackedUp {
			fmt.Printf("backup with id %d status is not backed up\n", backup.ID)
			return
		}

		// Restore backup
		if err := restore(backup); err != nil {
			fmt.Printf("error restoring backup with id %d, err: %v\n", backup.ID, err)
			return
		}
		fmt.Printf("restored backup with id %d to %s location\n", backup.ID, backup.SrcPath)

		// Update status
		if err := db.UpdateStatus(backup.ID, database.StatusRestored); err != nil {
			fmt.Printf("error updating status for backup with id %d, err: %v\n", backup.ID, err)
			return
		}
	}
}

// restore restores backups to their original source directory.
func restore(backup *database.Backup) error {
	if backup.Type == TypeFile {
		if err := util.MoveFile(backup.DstPath, backup.SrcPath); err != nil {
			return err
		}
	} else if backup.Type == TypeDirectory {
		if err := util.MoveDirectory(backup.DstPath, backup.SrcPath); err != nil {
			return err
		}
	} else {
		return errors.New("unknown type")
	}

	// Remove the backup(empty timestamp directory) path after restore
	if err := os.RemoveAll(backup.BackupPath); err != nil {
		return err
	}

	return nil
}
