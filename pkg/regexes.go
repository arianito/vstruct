package vstruct

import "regexp"

type RegexFactory struct {
	alpha *regexp.Regexp
}


var regexFactory *RegexFactory

func GetRegex() *RegexFactory {
	if regexFactory == nil {
		regexFactory = new(RegexFactory)
		regexFactory.alpha = regexp.MustCompile("^[a-zA-Z\\s]+$")
	}
	return regexFactory
}

func (r *RegexFactory) ValidateEnglishLetters(value string) bool {
	return r.alpha.MatchString(value)
}
