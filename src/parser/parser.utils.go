package parser

import "regexp"

func getRegexp(pattern string) *regexp.Regexp {
	regexpObject, err := regexp.Compile(pattern)

	if err != nil {
		doError(err)
	}

	return regexpObject
}
