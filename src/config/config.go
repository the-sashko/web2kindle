package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	TelegramBotTypeLog  = "log"
	TelegramBotTypeMain = "main"
	configFilePath      = "../config/config.json"
	credentialsFilePath = "../config/credentials.json"
)

var config configType

var credentials credentialsType

func setConfig() bool {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		doError(errors.New("config file not exists"))
	}

	configFile, err := os.Open(configFilePath)

	if err != nil {
		doError(err)
	}

	defer func() {
		if err := configFile.Close(); err != nil {
			doError(err)
		}
	}()

	jsonData, err := ioutil.ReadAll(configFile)

	if err != nil {
		doError(err)
	}

	err = json.Unmarshal(jsonData, &config)

	if err != nil {
		doError(err)
	}

	if !isValidConfig() {
		doError(errors.New("config has bad format"))
	}

	return true
}

func setCredentials() bool {
	if _, err := os.Stat(credentialsFilePath); os.IsNotExist(err) {
		doError(errors.New("credentials file not exists"))
	}

	credentialsFile, err := os.Open(credentialsFilePath)

	if err != nil {
		doError(err)
	}

	defer func() {
		if err := credentialsFile.Close(); err != nil {
			doError(err)
		}
	}()

	jsonData, err := ioutil.ReadAll(credentialsFile)

	if err != nil {
		doError(err)
	}

	err = json.Unmarshal(jsonData, &credentials)

	if err != nil {
		doError(err)
	}

	if !isValidCredentials() {
		doError(errors.New("credentials has bad format"))
	}

	return true
}

func GetProjectName() string {
	if !isValidConfig() {
		setConfig()
	}

	return config.ProjectName
}

func GetEmail() string {
	if !isValidConfig() {
		setConfig()
	}

	return config.Email
}

func GetTestUrl() string {
	if !isValidConfig() {
		setConfig()
	}

	return config.TestUrl
}

func GetTelegramBotCredentials(botType string) TelegramType {
	var telegramBotCredentials TelegramType

	if !isValidCredentials() {
		setCredentials()
	}

	switch botType {
	case TelegramBotTypeLog:
		telegramBotCredentials = credentials.TelegramLogBot
	case TelegramBotTypeMain:
		telegramBotCredentials = credentials.TelegramMainBot
	default:
		doError(fmt.Errorf("invalid telegram bot type: %s", botType))
	}

	return telegramBotCredentials
}

func GetSmtpCredentials() SmtpType {
	if !isValidCredentials() {
		setCredentials()
	}

	return credentials.SMTP
}

func isValidConfig() bool {
	if len(config.Email) < 1 {
		return false
	}

	if len(config.ProjectName) < 1 {
		return false
	}

	if len(config.TestUrl) < 1 {
		return false
	}

	return true
}

func isValidCredentials() bool {
	if len(credentials.TelegramMainBot.Token) < 1 {
		return false
	}

	if credentials.TelegramMainBot.ChatID == 0 {
		return false
	}

	if len(credentials.TelegramLogBot.Token) < 1 {
		return false
	}

	if credentials.TelegramLogBot.ChatID == 0 {
		return false
	}

	return true
}

func doError(errorEntity error) {
	fmt.Println(fmt.Sprintf("Config Error: %s", errorEntity.Error()))
	os.Exit(0)
}
