package saver

import (
	"errors"
	"os"
	"web2kindle/logger"
	"web2kindle/parser"
)

const markdownFilePath = "../tmp/document.md"

func SaveParserEntityToMarkdownFile(parserEntity parser.EntityType) {
	if _, err := os.Stat(markdownFilePath); !errors.Is(err, os.ErrNotExist) {
		err = os.Remove(markdownFilePath)
		doError(err)
	}

	markdownFile, err := os.Create(markdownFilePath)

	if err != nil {
		doError(err)
	}

	defer func(filePointer *os.File) {
		err := filePointer.Close()

		if err != nil {
			doError(err)
		}
	}(markdownFile)

	writeLineToFile(markdownFile, "# "+parserEntity.Title)
	writeLineToFile(markdownFile, "\n\n")
	writeLineToFile(markdownFile, parserEntity.Text)
}

func writeLineToFile(filePointer *os.File, line string) {
	_, err := filePointer.WriteString(line)

	if err != nil {
		doError(err)
	}
}

func doError(errorEntity error) {
	logger.LogError("Save", errorEntity, true)
}
