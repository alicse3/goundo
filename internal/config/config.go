package config

import (
	"encoding/json"
	"log"
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
)

// Config holds application configuration settings.
type Config struct {
	// BackupsPath is the path where backups will be stored.
	BackupsPath string `json:"backupsPath"`
}

// InitSetup initializes the application setup.
// It creates the necessary directories and sets up the default configuration.
func InitSetup() {
	// Get app path
	appPath, err := getAppPath()
	if err != nil {
		log.Fatal("error getting the app path", err)
	}
	log.Println("app path is:", appPath)

	// Create .goundo directory if it doesn't exist
	if err := createDir(appPath); err != nil {
		log.Fatal("error creating the .goundo dir: ", err)
	} else {
		// Load default config
		defCfg, err := loadDefaultConfig()
		if err != nil {
			log.Fatal("error loading default config: ", err)
		}
		log.Println("default config is loaded")

		// Setup default config
		if err := setupDefaultConfig(defCfg); err != nil {
			log.Fatal("error setting up default config: ", err)
		}
		log.Println("config file setup is done")
	}

	// Create backups directory if it doesn't exist
	backupPath := appPath + string(filepath.Separator) + backupsBaseDirname
	if err := createDir(backupPath); err != nil {
		log.Fatal("error creating the backups dir: ", err)
	}
}

// createDir creates a directory at the given path if it doesn't exist.
// It returns an error if the directory cannot be created or if there's an issue getting directory info.
func createDir(dirPath string) error {
	// Get directory info
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		// Create dir if it doesn't exist
		if os.IsNotExist(err) {
			if err := os.Mkdir(dirPath, 0755); err != nil {
				return &util.AppErr{Message: "error creating " + dirPath, Err: err}
			}
			log.Printf("directory '%s' is created\n", dirPath)
		} else {
			return &util.AppErr{Message: "error getting info for " + dirPath, Err: err}
		}
	} else {
		log.Printf("%s directory already exists\n", fileInfo.Name())
	}

	return nil
}

// setupDefaultConfig creates a new config file with default settings.
// It updates the BackupsPath field of the provided Config struct with the default backups directory path.
func setupDefaultConfig(config *Config) error {
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
	log.Printf("config file '%s' is created\n", file.Name())

	// Update default config
	updateDefaultConfig(goundoPath, config)

	// Marshal config data with indentation
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return &util.AppErr{Message: "error marshalling config data", Err: err}
	}

	// Write config data
	log.Println("writing config data to file")
	_, err = file.Write(data)
	if err != nil {
		return &util.AppErr{Message: "error writing config data", Err: err}
	}
	defer file.Close()

	return nil
}

// updateDefaultConfig updates the fields of the provided Config struct with the default values.
func updateDefaultConfig(baseAppPath string, config *Config) {
	if config.BackupsPath == "" {
		config.BackupsPath = baseAppPath + string(filepath.Separator) + backupsBaseDirname
	}
}

// loadDefaultConfig loads the default configuration from the config file.
// It returns a pointer to the Config struct and an error if any occurs.
func loadDefaultConfig() (*Config, error) {
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
	log.Println("default config data:", string(data))

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
