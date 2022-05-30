package parser

import (
	"html"
	"strings"
	"web2kindle/logger"
)

func parseHtml(text string) EntityType {
	var entity EntityType

	logger.LogMessage("start parsing html", "app")

	text = normalizeHtml(text)

	entity.Title = retrieveTitleFromHtml(text)
	entity.Text = retrieveTextFromHtml(text)
	entity.Summary = retrieveSummaryFromHtml(text)

	return entity
}

func retrieveTitleFromHtml(text string) string {
	var title string

	regexpObject := getRegexp(`(?is)^.*?<h1.*?>(.*?)<\/h1>.*?$`)

	if len(regexpObject.FindAllString(text, 1)) > 0 {
		title = regexpObject.ReplaceAllString(text, "$1")
	}

	regexpObject = getRegexp(`(?is)^.*?<title.*?>(.*?)<\/title>.*?$`)

	if len(regexpObject.FindAllString(text, 1)) > 0 {
		title = regexpObject.ReplaceAllString(text, "$1")
	}

	regexpObject = getRegexp(`\s*=\s*`)

	text = regexpObject.ReplaceAllString(text, "=")

	regexpObject = getRegexp(`(?isU)(<meta.*(property=("|'|)(twitter|og):title)("|'|).*>)`)

	if len(regexpObject.FindAllString(text, 1)) < 1 {
		return title
	}

	metaTagChunks := regexpObject.FindAllString(text, 1)
	metaTagChunks = strings.Split(metaTagChunks[0], "<meta")

	metaTag := metaTagChunks[len(metaTagChunks)-1]

	regexpObject = getRegexp(`(?isU)^.*content=("|'|)(.*?)$`)

	title = regexpObject.ReplaceAllString(metaTag, "$2")

	regexpObject = getRegexp(`(?isU)^(.*?)>$`)

	title = regexpObject.ReplaceAllString(title, "$1")

	regexpObject = getRegexp(`(?isU)^(.*)("|')$`)

	title = regexpObject.ReplaceAllString(title, "$1")

	regexpObject = getRegexp(`(?isU)^("|')(.*)("|')$`)

	title = regexpObject.ReplaceAllString(title, "$2")

	title = removeTags(title)

	regexpObject = getRegexp(`\s+`)

	title = regexpObject.ReplaceAllString(title, " ")

	title = removeSpaces(title)

	return normalizeText(title)
}

func retrieveTextFromHtml(text string) string {
	text = html.UnescapeString(text)

	regexpObject := getRegexp(`(?is)^.*?<body.*?>(.*?)<\/body>.*?$`)

	text = regexpObject.ReplaceAllString(text, "$1")

	regexpObject = getRegexp(`(?is)<header((.*?)|)>(.*?)<\/header>`)

	text = regexpObject.ReplaceAllString(text, " ")

	regexpObject = getRegexp(`(?is)^.*?<main.*?>(.*?)<\/main>.*?$`)

	text = regexpObject.ReplaceAllString(text, "$1")

	regexpObject = getRegexp(`(?is)^.*?<article.*?>(.*?)<\/article>.*?$`)

	text = regexpObject.ReplaceAllString(text, "$1")

	regexpObject = getRegexp(`(?isU)^("|')(.*)("|')$`)

	text = regexpObject.ReplaceAllString(text, "$2")

	text = normalizeText(text)

	return convertHtml2Markdown(text)
}

func retrieveSummaryFromHtml(text string) string {
	summary := retrieveTextFromHtml(text)
	summary = removeMarkdown(summary)

	summaryChunks := strings.Split(summary, "\n")

	summary = summaryChunks[len(summaryChunks)-1]

	summary = removeSpaces(summary)

	regexpObject := getRegexp(`\s*=\s*`)

	text = regexpObject.ReplaceAllString(text, "=")

	regexpObject = getRegexp(`(?isU)(<meta.*((property=("|'|)(twitter|og):description("|'|))|(name=("|'|)description("|'|))).*>)`)

	if len(regexpObject.FindAllString(text, 1)) < 1 {
		return normalizeText(summary)
	}

	metaTagChunks := regexpObject.FindAllString(text, 1)
	metaTagChunks = strings.Split(metaTagChunks[0], "<meta")

	metaTag := metaTagChunks[len(metaTagChunks)-1]

	regexpObject = getRegexp(`(?isU)^.*content=(.*?)$`)

	summary = regexpObject.ReplaceAllString(metaTag, "$1")

	regexpObject = getRegexp(`(?isU)^("|'|)(.*?)$`)

	summary = regexpObject.ReplaceAllString(summary, "$2")

	regexpObject = getRegexp(`(?isU)^(.*?)>$`)

	summary = regexpObject.ReplaceAllString(summary, "$1")

	regexpObject = getRegexp(`(?isU)^(.*)("|'|)$`)

	summary = regexpObject.ReplaceAllString(summary, "$1")

	regexpObject = getRegexp(`(?isU)^("|')(.*)("|')$`)

	summary = regexpObject.ReplaceAllString(summary, "$2")

	summary = removeTags(summary)
	summary = removeSpaces(summary)

	return normalizeText(summary)
}

func removeTag(html string, tag string) string {
	regexpObject := getRegexp(`(?is)<` + tag + `(.*?)>.*?<\/` + tag + `>`)

	html = regexpObject.ReplaceAllString(html, " ")

	regexpObject = getRegexp(`(?is)<` + tag + `(.*?)>`)

	html = regexpObject.ReplaceAllString(html, " ")

	regexpObject = getRegexp(`(?is)<\/` + tag + `(.*?)>`)

	return regexpObject.ReplaceAllString(html, " ")
}

func removeTags(text string) string {
	text = html.UnescapeString(text)

	regexpObject := getRegexp(`<.*?>`)

	text = regexpObject.ReplaceAllString(text, " ")

	return removeSpaces(text)
}

func convertHtml2Markdown(text string) string {
	text = html.UnescapeString(text)

	regexpObject := getRegexp(`<([^>]+)\s*>`)

	text = regexpObject.ReplaceAllString(text, "<$1>")

	regexpObject = getRegexp(`<\s*([^>]+)>`)

	text = regexpObject.ReplaceAllString(text, "<$1>")

	regexpObject = getRegexp(`<([^>\s]+)\s([^>]+)>`)

	text = regexpObject.ReplaceAllString(text, "<$1>")

	regexpObject = getRegexp(`<([^>\s]+)/>`)

	text = regexpObject.ReplaceAllString(text, "<$1>")

	regexpObject = getRegexp(`<([^>/]+)>\s+`)

	text = regexpObject.ReplaceAllString(text, "<$1>")

	regexpObject = getRegexp(`\s+</([^>/]+)>`)

	text = regexpObject.ReplaceAllString(text, "<$1>")

	regexpObject = getRegexp(`(?i)(<br>)|(<(/|)p>)|(<(/|)div>)|(<(/|)pre>)|(<(/|)blockquote>)`)

	text = regexpObject.ReplaceAllString(text, "\n")

	regexpObject = getRegexp(`(?i)(<(/|)i>)|(<(/|)em>)`)

	text = regexpObject.ReplaceAllString(text, "*")

	regexpObject = getRegexp(`(?i)\*\*\*\*`)

	text = regexpObject.ReplaceAllString(text, "* *")

	regexpObject = getRegexp(`(?i)\*\*`)

	text = regexpObject.ReplaceAllString(text, "*")

	regexpObject = getRegexp(`(?i)(<(/|)b>)|(<(/|)strong>)`)

	text = regexpObject.ReplaceAllString(text, "**")

	regexpObject = getRegexp(`(?i)(<h([0-9])>)`)

	text = regexpObject.ReplaceAllString(text, "\n**")

	regexpObject = getRegexp(`(?i)(</h([0-9])>)`)

	text = regexpObject.ReplaceAllString(text, "**\n")

	regexpObject = getRegexp(`(?i)\*\*\*\*([^\*]+)\*\*\*\*`)

	text = regexpObject.ReplaceAllString(text, "**$1**")

	regexpObject = getRegexp(`(?i)(<(/|)s>)|(<(/|)strike>)`)

	text = regexpObject.ReplaceAllString(text, "~~")

	regexpObject = getRegexp(`(?i)(<(/|)s>)|(<(/|)strike>)`)

	text = regexpObject.ReplaceAllString(text, "~~")

	text = removeTags(text)

	text = removeSpaces(text)

	regexpObject = getRegexp(`(?i)\*\* \*\*`)

	text = regexpObject.ReplaceAllString(text, " ")

	regexpObject = getRegexp(`(?i)\* \*`)

	text = regexpObject.ReplaceAllString(text, " ")

	text = removeSpaces(text)

	regexpObject = getRegexp(`(?i)(\n)`)

	return regexpObject.ReplaceAllString(text, "\n\n")
}

func removeMarkdown(text string) string {
	text = html.UnescapeString(text)

	regexpObject := getRegexp(`(\*|\~|#)`)

	text = regexpObject.ReplaceAllString(text, "")

	return removeSpaces(text)
}
