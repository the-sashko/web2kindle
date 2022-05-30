package parser

func normalizeHtml(html string) string {
	html = removeSpaces(html)

	html = normalizeTags(html)

	html = removeCommentTags(html)
	html = removeDoctypeTags(html)
	html = removeScriptTags(html)
	html = removeStyleTags(html)
	html = removeButtonTags(html)
	html = removeNavTags(html)
	html = removeIconTags(html)
	html = removeMapTags(html)
	html = removeCanvasTags(html)
	html = removeAsideTags(html)
	html = removeBaseTags(html)
	html = removeDialogTags(html)
	html = removeEmbedTags(html)
	html = removeFigureTags(html)
	html = removeFooterTags(html)
	html = removeLinkTags(html)
	html = removeObjectTags(html)
	html = removeProgressTags(html)
	html = removeTemplateTags(html)

	html = removeForms(html)
	html = removeMedia(html)
	html = removeTables(html)
	html = removeLists(html)
	html = removeFrames(html)

	html = removeDeprecatedTags(html)

	return removeSpaces(html)
}

func normalizeTags(html string) string {
	regexpObject := getRegexp(`<\s*/\s*(.*?)\s*>`)

	html = regexpObject.ReplaceAllString(html, "</$1>")

	regexpObject = getRegexp(`<\s*(.*?)\s*>`)

	html = regexpObject.ReplaceAllString(html, "<$1>")

	regexpObject = getRegexp(`<(.*?)/\s*>`)

	html = regexpObject.ReplaceAllString(html, "<$1/>")

	regexpObject = getRegexp(`<(.*?)\s*/>`)

	html = regexpObject.ReplaceAllString(html, "<$1/>")

	regexpObject = getRegexp(`<(.*?)/>`)

	html = regexpObject.ReplaceAllString(html, "<$1>")

	return html
}

func removeCommentTags(html string) string {
	regexpObject := getRegexp(`<!\-\-((.*?)|)\-\->`)

	return regexpObject.ReplaceAllString(html, " ")
}

func removeDoctypeTags(html string) string {
	return removeTag(html, `!doctype`)
}

func removeScriptTags(html string) string {
	html = removeTag(html, `script`)

	return removeTag(html, `noscript`)
}

func removeStyleTags(html string) string {
	return removeTag(html, `style`)
}

func removeButtonTags(html string) string {
	return removeTag(html, `button`)
}

func removeNavTags(html string) string {
	return removeTag(html, `nav`)
}

func removeIconTags(html string) string {
	regexpObject := getRegexp(`(?is)<i (.*?)>.*?<\/i>`)

	return regexpObject.ReplaceAllString(html, " ")
}

func removeMapTags(html string) string {
	html = removeTag(html, `map`)

	return removeTag(html, `area`)
}

func removeCanvasTags(html string) string {
	return removeTag(html, `canvas`)
}

func removeAsideTags(html string) string {
	return removeTag(html, `aside`)
}

func removeBaseTags(html string) string {
	return removeTag(html, `base`)
}

func removeDialogTags(html string) string {
	return removeTag(html, `dialog`)
}

func removeEmbedTags(html string) string {
	return removeTag(html, `embed`)
}

func removeFigureTags(html string) string {
	html = removeTag(html, `figure`)

	return removeTag(html, `figcaption`)
}

func removeFooterTags(html string) string {
	return removeTag(html, `footer`)
}

func removeLinkTags(html string) string {
	return removeTag(html, `link`)
}

func removeObjectTags(html string) string {
	html = removeTag(html, `object`)

	return removeTag(html, `param`)
}

func removeProgressTags(html string) string {
	return removeTag(html, `progress`)
}

func removeTemplateTags(html string) string {
	return removeTag(html, `template`)
}

func removeForms(html string) string {
	html = removeTag(html, `form`)
	html = removeTag(html, `input`)
	html = removeTag(html, `output`)
	html = removeTag(html, `label`)
	html = removeTag(html, `textarea`)
	html = removeTag(html, `select`)
	html = removeTag(html, `fieldset`)
	html = removeTag(html, `legend`)

	return removeTag(html, `option`)
}

func removeMedia(html string) string {
	html = removeTag(html, `img`)
	html = removeTag(html, `svg`)
	html = removeTag(html, `picture`)
	html = removeTag(html, `source`)
	html = removeTag(html, `track`)
	html = removeTag(html, `audio`)

	return removeTag(html, `video`)
}

func removeTables(html string) string {
	html = removeTag(html, `table`)
	html = removeTag(html, `tr`)
	html = removeTag(html, `td`)
	html = removeTag(html, `colgroup`)
	html = removeTag(html, `col`)
	html = removeTag(html, `th`)
	html = removeTag(html, `tbody`)
	html = removeTag(html, `thead`)

	return removeTag(html, `tfoot`)
}

func removeLists(html string) string {
	html = removeTag(html, `ol`)
	html = removeTag(html, `ul`)
	html = removeTag(html, `li`)

	return removeTag(html, `datalist`)
}

func removeFrames(html string) string {
	html = removeTag(html, `iframe`)
	html = removeTag(html, `frame`)

	return removeTag(html, `noframe`)
}

func removeDeprecatedTags(html string) string {
	html = removeTag(html, `applet`)
	html = removeTag(html, `basefont`)
	html = removeTag(html, `dir`)

	return removeTag(html, `font`)
}
