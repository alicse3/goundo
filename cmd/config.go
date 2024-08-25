package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alicse3/goundo/internal/util"
)

const (
	// defaultAppDir is the name of the directory where the application's data is stored.
	defaultAppDir = ".goundo"

	// defaultBackupsDir is the name of the directory where the default backups are stored.
	defaultBackupsDir = "backups"

	// defaultConfigFile is the name of the default configuration file.
	defaultConfigFile = ".goundo_config.json"

	// defaultSqliteDBFile is the name of the default Sqlite db file.
	defaultSqliteDBFile = "backups.db"
)

// configHandler handle's the configuration initialization.
func configHandler() {
	cfg := getConfig()
	cfg.promptForUpdates()
}

// InitSetup sets up the default configuration if not exists.
func InitSetup() {
	if !isDefaultAppConfigExists() {
		cfg := initDefault()
		applyConfig(cfg)
	}
}

// isDefaultAppConfigExists reports whether the default app config exists or not.
func isDefaultAppConfigExists() bool {
	path := getDefaultConfigPath()
	_, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

// getDefaultConfigPath retrieves the default app's config path
func getDefaultConfigPath() string {
	// Construct and return default app config path
	return filepath.Join(getHomeDir(), defaultConfigFile)
}

// getDefaultAppPath retrieves the app's default path
func getDefaultAppPath() string {
	// Construct and return default app path
	return filepath.Join(getHomeDir(), defaultAppDir)
}

// getHomeDir returns the user home directory
func getHomeDir() string {
	// Get home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		util.ExitOnError(err, "error getting the user home dir")
	}
	return homeDir
}

// configuration holds the application configuration settings.
type configuration struct {
	// AppPath is the path where the application configuration is stored.
	AppPath string `json:"appPath"`

	// BackupsPath is the path where the removed files/directories will be stored.
	BackupsPath string `json:"backupsPath"`

	// ConfigFilepath is the path where the configuration file is stored.
	ConfigFilepath string `json:"configFilepath"`

	// SqliteDBPath is where the Sqlite database file is stored.
	SqliteDBPath string `json:"sqliteDBPath"`
}

// initDefault initilizes and returns the default configuration.
func initDefault() *configuration {
	return &configuration{
		AppPath:        getDefaultAppPath(),
		BackupsPath:    filepath.Join(getDefaultAppPath(), defaultBackupsDir),
		ConfigFilepath: getDefaultConfigPath(),
		SqliteDBPath:   filepath.Join(getDefaultAppPath(), defaultSqliteDBFile),
	}
}

// updateAppPathForAll updates the AppPath for all the paths that require it.
func (cfg *configuration) updateAppPathForAll() {
	cfg.BackupsPath = filepath.Join(cfg.AppPath, defaultBackupsDir)
	cfg.SqliteDBPath = filepath.Join(cfg.AppPath, defaultSqliteDBFile)
}

// promptForUpdates prompts for config values update.
func (cfg *configuration) promptForUpdates() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please enter the app configuration path [%s]: ", cfg.AppPath)
	scanner.Scan()
	appPath := strings.TrimSpace(scanner.Text())
	if appPath != "" {
		cfg.AppPath = appPath
		cfg.updateAppPathForAll()
	}

	fmt.Printf("Please enter the backups path [%s]: ", cfg.BackupsPath)
	scanner.Scan()
	backupsPath := strings.TrimSpace(scanner.Text())
	if backupsPath != "" {
		cfg.BackupsPath = backupsPath
	}

	fmt.Print("Save this configuration? (yes/no) [yes]: ")
	scanner.Scan()
	saveConfig := strings.TrimSpace(scanner.Text())
	if saveConfig == "" || saveConfig == "yes" {
		// If yes, apply the custom configuration
		applyConfig(cfg)
		fmt.Println("Configuration saved successfully.")
	} else if saveConfig == "no" {
		// If no, don't do anything
		fmt.Println("Configuration not saved.")
	} else {
		fmt.Println("Invalid input. Configuration failed.")
	}
}

// applyConfig applies config changes by creating necessary files/directories.
func applyConfig(cfg *configuration) {
	util.CreateDir(cfg.AppPath)
	util.CreateDir(cfg.BackupsPath)

	// Marshal the configuration to JSON
	data, err := json.Marshal(cfg)
	if err != nil {
		util.ExitOnError(err, "error marshaling the configuration")
	}

	util.WriteToFile(cfg.ConfigFilepath, string(data))
}

// getConfig reads data from the config file and returns it.
func getConfig() *configuration {
	data, err := os.ReadFile(getDefaultConfigPath())
	if err != nil {
		util.ExitOnError(err, "error reading the configuration file")
	}

	var cfg configuration
	if err := json.Unmarshal(data, &cfg); err != nil {
		util.ExitOnError(err, "error unmarshalling the config data")
	}

	return &cfg
}
