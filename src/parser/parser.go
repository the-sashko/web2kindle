package parser

import (
	"web2kindle/logger"
)

func DoParseHtml(html string) EntityType {
	return parseHtml(html)
}

func doError(errorEntity error) {
	logger.LogError("Parser", errorEntity, true)
}

func removeSpaces(text string) string {
	regexpObject := getRegexp(` `)

	text = regexpObject.ReplaceAllString(text, " ")

	regexpObject = getRegexp(`&nbsp;`)

	text = regexpObject.ReplaceAllString(text, " ")

	regexpObject = getRegexp(`(\r?\n)`)

	text = regexpObject.ReplaceAllString(text, "<br>")

	regexpObject = getRegexp(`\s+`)

	text = regexpObject.ReplaceAllString(text, " ")

	regexpObject = getRegexp(`(\s*<br>\s*)+`)

	text = regexpObject.ReplaceAllString(text, "\n")

	regexpObject = getRegexp(`((^\s)|(\s$))`)

	return regexpObject.ReplaceAllString(text, "")
}
