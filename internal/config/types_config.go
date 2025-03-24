package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	newConfig := Config{}
	file_loc, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(file_loc)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &newConfig)
	if err != nil {
		return Config{}, err
	}

	return newConfig, nil

}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fileLocation := fmt.Sprintf("%v/%v", homeDir, configFileName)

	return fileLocation, nil
}

func (c *Config) SetUser(newUserName string) error {
	c.CurrentUserName = newUserName

	return write(c)
}

func write(config *Config) error {
	fileLoc, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileLoc, data, 0777)
	if err != nil {
		return err
	}

	return nil
}
