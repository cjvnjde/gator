// Package config provides a shared app configuration
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}

	fileContent, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var configData Config
	if err := json.Unmarshal(fileContent, &configData); err != nil {
		return Config{}, err
	}

	return configData, nil
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}

func write(config Config) error {
	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, jsonData, 0o600)
	if err != nil {
		return err
	}
	return nil
}
