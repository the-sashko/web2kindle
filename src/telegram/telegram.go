package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"web2kindle/config"
)

const (
	lastUpdateIdFilePath    = "../tmp/telegram_last_update_id.txt"
	apiSendingUrlPattern    = "https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s"
	apiRetrievingUrlPattern = "https://api.telegram.org/bot%s/getUpdates?limit=1&offset=%d"
	httpStatusOk            = 200
)

func SendError(text string) {
	text = fmt.Sprintf("#%s %s", config.GetProjectName(), text)

	credentials := config.GetTelegramBotCredentials(config.TelegramBotTypeLog)

	sendToTelegramAPI(text, credentials.ChatID, credentials.Token)
}

func GetMessage() *string {
	credentials := config.GetTelegramBotCredentials(config.TelegramBotTypeMain)
	lastUpdateId := getLastUpdateId()

	return getFromTelegramAPI(credentials.ChatID, credentials.Token, lastUpdateId)
}

func getLastUpdateId() int {
	if _, err := os.Stat(lastUpdateIdFilePath); errors.Is(err, os.ErrNotExist) {
		return 0
	}

	lastUpdateId, err := os.ReadFile(lastUpdateIdFilePath)

	if err != nil {
		doError(err)
	}

	if len(lastUpdateId) < 1 {
		return 0
	}

	lastUpdateIdInt, err := strconv.Atoi(string(lastUpdateId))

	if err != nil {
		doError(err)
	}

	return lastUpdateIdInt
}

func saveLastUpdateId(lastUpdateId int) {
	lastUpdateIdFile, err := os.OpenFile(lastUpdateIdFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		doError(err)
	}

	defer func(filePointer *os.File) {
		err := filePointer.Close()

		if err != nil {
			doError(err)
		}
	}(lastUpdateIdFile)

	_, err = lastUpdateIdFile.WriteString(strconv.Itoa(lastUpdateId))

	if err != nil {
		doError(err)
	}
}

func sendToTelegramAPI(text string, chatID int, token string) {
	var response apiSendResponse

	text = url.QueryEscape(text)

	apiUrl := getApiUrlForSending(text, chatID, token)

	response = sendToRemote(apiUrl)

	if len(response.ErrorDescription) < 1 {
		response.ErrorDescription = "unknown telegram api error"
	}

	if !response.Status {
		doError(errors.New(response.ErrorDescription))
	}
}

func getFromTelegramAPI(chatID int, token string, lastUpdateId int) *string {
	var response apiResponse

	apiUrl := getApiUrlForRetrieving(token, lastUpdateId)

	response = getFromRemote(apiUrl)

	if len(response.ErrorDescription) < 1 {
		response.ErrorDescription = "unknown telegram api error"
	}

	if !response.Status {
		doError(errors.New(response.ErrorDescription))
	}

	if len(response.Result) < 1 {
		return nil
	}

	responseResult := response.Result[0]

	saveLastUpdateId(responseResult.UpdateId)

	MessageFromId := responseResult.Message.From.Id
	MessageChatId := responseResult.Message.Chat.Id
	IsBot := responseResult.Message.From.IsBot

	if MessageFromId != chatID || MessageChatId != chatID || IsBot {
		return nil
	}

	return &responseResult.Message.Text
}

func getApiUrlForSending(text string, chatID int, token string) string {
	apiUrl := fmt.Sprintf(apiSendingUrlPattern, token, chatID, text)

	if len(apiUrl) > 2048 {
		apiUrl = apiUrl[:2048]
	}

	return apiUrl
}

func getApiUrlForRetrieving(token string, lastUpdateId int) string {
	offset := lastUpdateId + 1

	return fmt.Sprintf(apiRetrievingUrlPattern, token, offset)
}

func sendToRemote(url string) apiSendResponse {
	var response apiSendResponse

	remoteResponse, err := http.Get(url)

	if err != nil {
		doError(err)
	}

	if remoteResponse.StatusCode != httpStatusOk {
		doError(fmt.Errorf("HTTP Response Code %d", remoteResponse.StatusCode))
	}

	jsonDecoder := json.NewDecoder(remoteResponse.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()

		if err != nil {
			doError(err)
		}
	}(remoteResponse.Body)

	err = jsonDecoder.Decode(&response)

	if err != nil {
		doError(err)
	}

	return response
}

func getFromRemote(url string) apiResponse {
	var response apiResponse

	remoteResponse, err := http.Get(url)

	if err != nil {
		doError(err)
	}

	if remoteResponse.StatusCode != httpStatusOk {
		doError(fmt.Errorf("HTTP Response Code %d", remoteResponse.StatusCode))
	}

	jsonDecoder := json.NewDecoder(remoteResponse.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()

		if err != nil {
			doError(err)
		}
	}(remoteResponse.Body)

	err = jsonDecoder.Decode(&response)

	if err != nil {
		doError(err)
	}

	return response
}

func doError(errorEntity error) {
	fmt.Println(fmt.Sprintf("Telegram error: %s", errorEntity.Error()))
	os.Exit(0)
}
