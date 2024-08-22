package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alicse3/goundo/internal/util"
)

const (
	// goundoDirname is the name of the directory where the application's data is stored.
	goundoDirname = ".goundo"

	// configFilename is the name of the configuration file.
	configFilename = "config.json"

	// backupsBaseDirname is the name of the directory where backups are stored.
	backupsBaseDirname = "backups"

	// setupLogFilename is the name of the setup log file.
	setupLogFilename = "setup.log"
)

// Config holds application configuration settings.
type Config struct {
	// AppPath is the path where the application configuration is stored.
	AppPath string `json:"appPath"`

	// BackupsPath is the path where backups will be stored.
	BackupsPath string `json:"backupsPath"`
}

// GetConfig returns the app config.
func GetConfig() (*Config, error) {
	// Get app path
	appPath, err := getAppPath()
	if err != nil {
		return nil, &util.AppErr{Message: "error getting the app path", Err: err}
	}

	// Read config data
	configFilepath := appPath + string(filepath.Separator) + configFilename
	data, err := os.ReadFile(configFilepath)
	if err != nil {
		return nil, &util.AppErr{Message: "error reading the config file", Err: err}
	}

	// Unmarshal config data
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, &util.AppErr{Message: "error unmarshalling config data", Err: err}
	}

	return &config, nil
}

// InitSetup initializes the application setup.
// It creates the necessary directories and sets up the default configuration.
func InitSetup() {
	// Get app path
	appPath, err := getAppPath()
	if err != nil {
		fmt.Printf("error getting the app path: %v\n", err)
		return
	}

	// Create .goundo directory if it doesn't exist
	if err := createDir(appPath); err != nil {
		fmt.Printf("error creating the .goundo dir: %v\n", err)
		return
	}

	// Create logger
	setupLogPath := appPath + string(filepath.Separator) + setupLogFilename
	logger, err := util.NewLogger("DEBUG", setupLogPath)
	if err != nil {
		fmt.Printf("error initializing the logger: %v\n", err)
		return
	}

	logger.Debug("app path is: %s", appPath)

	// Load default config
	defCfg, err := loadDefaultConfig(logger)
	if err != nil {
		logger.Error("error loading default config: %v", err)
		return
	}
	logger.Info("default config is loaded")

	// Setup default config
	if err := setupDefaultConfig(logger, defCfg); err != nil {
		logger.Error("error setting up default config: %v", err)
		return
	}
	logger.Info("config file setup is done")

	// Create backups directory if it doesn't exist
	backupPath := appPath + string(filepath.Separator) + backupsBaseDirname
	if err := createDir(backupPath); err != nil {
		logger.Error("error creating the backups dir: %v", err)
		return
	}
}

// createDir creates a directory at the given path if it doesn't exist.
// It returns an error if the directory cannot be created or if there's an issue getting directory info.
func createDir(dirPath string) error {
	// Get directory info
	_, err := os.Stat(dirPath)
	if err != nil {
		// Create dir if it doesn't exist
		if os.IsNotExist(err) {
			if err := os.Mkdir(dirPath, 0755); err != nil {
				return &util.AppErr{Message: "error creating " + dirPath, Err: err}
			}
		} else {
			return &util.AppErr{Message: "error getting info for " + dirPath, Err: err}
		}
	}

	return nil
}

// setupDefaultConfig creates a new config file with default settings.
// It updates the BackupsPath field of the provided Config struct with the default backups directory path.
func setupDefaultConfig(logger *util.Logger, config *Config) error {
	// Get app path
	goundoPath, err := getAppPath()
	if err != nil {
		return &util.AppErr{Message: "error getting the app path", Err: err}
	}

	// Create config file
	configFilepath := goundoPath + string(filepath.Separator) + configFilename
	file, err := os.Create(configFilepath)
	if err != nil {
		return &util.AppErr{Message: "error creating the config file", Err: err}
	}
	logger.Debug("config file '%s' is created\n", file.Name())

	// Update default config
	updateDefaultConfig(goundoPath, config)

	// Marshal config data with indentation
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return &util.AppErr{Message: "error marshalling config data", Err: err}
	}

	logger.Debug("updated default config data: %v", string(data))

	// Write config data
	logger.Info("writing config data to file")
	_, err = file.Write(data)
	if err != nil {
		return &util.AppErr{Message: "error writing config data", Err: err}
	}
	defer file.Close()

	return nil
}

// updateDefaultConfig updates the fields of the provided Config struct with the default values.
func updateDefaultConfig(baseAppPath string, config *Config) {
	if config.AppPath == "" {
		config.AppPath = baseAppPath
	}
	if config.BackupsPath == "" {
		config.BackupsPath = baseAppPath + string(filepath.Separator) + backupsBaseDirname
	}
}

// loadDefaultConfig loads the default configuration from the config file.
// It returns a pointer to the Config struct and an error if any occurs.
func loadDefaultConfig(logger *util.Logger) (*Config, error) {
	// Get current workind dir
	wd, err := os.Getwd()
	if err != nil {
		return nil, &util.AppErr{Message: "error getting work directory", Err: err}
	}

	// Read default config data
	configPath := wd + string(filepath.Separator) + configFilename
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, &util.AppErr{Message: "error reading work directory", Err: err}
	}
	logger.Debug("default config data: %v", string(data))

	// Unmarshal config data
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, &util.AppErr{Message: "error unmarshalling config data", Err: err}
	}

	return &config, nil
}

// getAppPath retrieves the absolute path to the current application executable.
// It returns the path as a string and an error if any occurs.
func getAppPath() (string, error) {
	// Get home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", &util.AppErr{Message: "error getting the user home dir", Err: err}
	}

	// Construct and return app dir
	return homeDir + string(filepath.Separator) + goundoDirname, nil
}
