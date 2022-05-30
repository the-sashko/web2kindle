package parser

func normalizeText(text string) string {
	text = removeSpaces(text)

	regexpObject := getRegexp(`­`)

	text = regexpObject.ReplaceAllString(text, "")

	regexpObject = getRegexp(`(?is)([«»‹›‟”„“]+)`)

	text = regexpObject.ReplaceAllString(text, " \" ")

	regexpObject = getRegexp(`(?is)"([^"]+)"`)

	text = regexpObject.ReplaceAllString(text, "«$1»")

	regexpObject = getRegexp(`(?is)([\(|\[|\)|\]])`)

	text = regexpObject.ReplaceAllString(text, " $1 ")

	text = removeSpaces(text)

	regexpObject = getRegexp(`(?is)([«\(|\[])\s`)

	text = regexpObject.ReplaceAllString(text, "$1")

	regexpObject = getRegexp(`(?is)\s([»|\)|\]])`)

	text = regexpObject.ReplaceAllString(text, "$1")

	regexpObject = getRegexp(`(?is)\s([,\.:;])`)

	text = regexpObject.ReplaceAllString(text, "$1")

	return removeSpaces(text)
}
