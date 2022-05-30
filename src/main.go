package main

import (
	"flag"
	"fmt"
	"net/url"
	"web2kindle/config"
	"web2kindle/logger"
	"web2kindle/parser"
	"web2kindle/remote"
	"web2kindle/saver"
	"web2kindle/script"
	"web2kindle/sender"
	"web2kindle/telegram"
)

const (
	runModeDefault = "default"
	runModeTest    = "test"
)

func main() {
	var urlPointer *string

	logger.LogMessage("app started", "app")

	runMode := getRumMode()

	logger.LogMessage(fmt.Sprintf("run mode: %s", runMode), "app")

	if runMode == runModeTest {
		doHandleUrl(getTestUrl())
	}

	if runMode == runModeDefault {
		logger.LogMessage("start retrieving new messages from telegram", "app")

		for {
			urlPointer = getUrlFromTelegram()

			if urlPointer == nil {
				break
			}

			doHandleUrl(*urlPointer)
		}
	}

	logger.LogMessage("app closed", "app")
}

func doHandleUrl(url string) {
	logger.LogMessage(fmt.Sprintf("start handling URL: %s", url), "app")

	parserEntity := getParserEntityByUrl(url)

	logger.LogMessage("start saving content to markdown file", "app")

	saver.SaveParserEntityToMarkdownFile(parserEntity)

	logger.LogMessage("start converting markdown to mobi", "app")

	script.RunConvert()

	logger.LogMessage("start sending file to email", "app")

	sender.Send()

	logger.LogMessage(fmt.Sprintf("finished handling URL: %s", url), "app")
}

func getParserEntityByUrl(url string) parser.EntityType {
	logger.LogMessage("start retrieving HTML from remote", "app")

	html := remote.GetHtml(url)

	return parser.DoParseHtml(html)
}

func getRumMode() string {
	var runMode string

	flag.StringVar(&runMode, "mode", runModeDefault, "running mode")

	flag.Parse()

	switch runMode {
	case runModeDefault:
		runMode = runModeDefault
	case runModeTest:
		runMode = runModeTest
	default:
		doError(fmt.Errorf("unsupported run mode: %s", runMode))
	}

	return runMode
}

func getTestUrl() string {
	return config.GetTestUrl()
}

func getUrlFromTelegram() *string {
	messagePointer := telegram.GetMessage()

	if messagePointer == nil {
		return nil
	}

	logger.LogMessage(fmt.Sprintf("got message from telegram: %s", *messagePointer), "app")

	_, err := url.ParseRequestURI(*messagePointer)

	if err != nil {
		return nil
	}

	return messagePointer
}

func doError(errorEntity error) {
	logger.LogError("App", errorEntity, true)
}
