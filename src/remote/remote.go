package remote

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"web2kindle/logger"
)

const httpStatusOk = 200

func Send(url string) (*http.Response, error) {
	response, err := http.Get(url)

	if err != nil {
		return response, err
	}

	if response.StatusCode != httpStatusOk {
		return response, fmt.Errorf("HTTP Response Code %d", response.StatusCode)
	}

	return response, nil
}

func GetHtml(url string) string {
	response, err := Send(url)

	if err != nil {
		doError(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()

		if err != nil {
			doError(err)
		}
	}(response.Body)

	html, err := ioutil.ReadAll(response.Body)

	if err != nil {
		doError(err)
	}

	return string(html)
}

func doError(errorEntity error) {
	logger.LogError("Remote", errorEntity, true)
}
