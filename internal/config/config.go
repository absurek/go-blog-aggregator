package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DataBaseURL     string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homedir, configFileName), nil
}

func write(config Config) error {
	dat, err := json.Marshal(config)
	if err != nil {
		return err
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, dat, 0664)
}

func Read() (*Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("get config path: %w", err)
	}

	dat, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read config at %s: %w", configPath, err)
	}

	var config Config
	err = json.Unmarshal(dat, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config at %s: %w", configPath, err)
	}

	return &config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username

	err := write(*c)
	if err != nil {
		return fmt.Errorf("write user to file: %w", err)
	}

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("Config{ database_url: %s, current_user_name: %s }\n", c.DataBaseURL, c.CurrentUsername)
}
