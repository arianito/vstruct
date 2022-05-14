package vstruct

import "regexp"

type RegexFactory struct {
	alpha      *regexp.Regexp
	alpha_num  *regexp.Regexp
	alpha_dash *regexp.Regexp
	username   *regexp.Regexp
	email      *regexp.Regexp
}

var regexFactory *RegexFactory

func GetRegex() *RegexFactory {
	if regexFactory == nil {
		regexFactory = new(RegexFactory)
		regexFactory.alpha = regexp.MustCompile(`^[a-zA-Z\s]+$`)
		regexFactory.alpha_num = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
		regexFactory.alpha_dash = regexp.MustCompile(`^[a-zA-Z]+[\-_a-zA-Z0-9]+$`)
		regexFactory.username = regexp.MustCompile("^[a-zA-Z]+[_a-zA-Z0-9.]+$")
		regexFactory.email = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	}
	return regexFactory
}

func (r *RegexFactory) Alpha(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.alpha.MatchString(value)
}
func (r *RegexFactory) AlphaNum(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.alpha_num.MatchString(value)
}
func (r *RegexFactory) AlphaDash(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.alpha_dash.MatchString(value)
}
func (r *RegexFactory) Username(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.username.MatchString(value)
}
func (r *RegexFactory) Email(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.email.MatchString(value)
}
