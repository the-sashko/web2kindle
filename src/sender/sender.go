package sender

import (
	"github.com/go-gomail/gomail"
	"io"
	"os"
	"web2kindle/config"
	"web2kindle/logger"
)

const (
	documentFileName = "document.mobi"
	documentFilePath = "../tmp/document.mobi"
)

func Send() {
	credentials := config.GetSmtpCredentials()

	message := gomail.NewMessage()
	message.SetHeader("From", credentials.Sender)
	message.SetHeader("To", config.GetEmail())
	message.SetHeader("Cc", credentials.Sender)
	message.SetHeader("Subject", config.GetProjectName())
	message.SetBody("text/plain", "File sent by web2kindle app")
	message.Attach(
		documentFileName,
		gomail.SetCopyFunc(func(w io.Writer) error {
			mobiFileContent, err := os.ReadFile(documentFilePath)

			if err != nil {
				doError(err)
			}

			_, err = w.Write(mobiFileContent)

			if err != nil {
				doError(err)
			}

			return err
		}),
		gomail.SetHeader(map[string][]string{"Content-Type": {"application/x-mobipocket-ebook"}}),
	)

	goMailDialer := gomail.NewDialer(credentials.Host, credentials.Port, credentials.Username, credentials.Password)

	err := goMailDialer.DialAndSend(message)

	if err != nil {
		doError(err)
	}
}

func doError(errorEntity error) {
	logger.LogError("Sender", errorEntity, true)
}
