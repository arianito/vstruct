package vstruct

import "regexp"

type RegexFactory struct {
	alpha       *regexp.Regexp
	alpha_num   *regexp.Regexp
	alpha_dash  *regexp.Regexp
	username  *regexp.Regexp
	email       *regexp.Regexp
	phone_iran  *regexp.Regexp
	mobile_iran *regexp.Regexp
	string      *regexp.Regexp
	persian     *regexp.Regexp
}

var regexFactory *RegexFactory

func GetRegex() *RegexFactory {
	if regexFactory == nil {
		regexFactory = new(RegexFactory)
		regexFactory.alpha = regexp.MustCompile("^[a-zA-Z\\s]+$")
		regexFactory.alpha_num = regexp.MustCompile("^[a-zA-Z0-9\\s]+$")
		regexFactory.alpha_dash = regexp.MustCompile("^[a-zA-Z]+[\\-_a-zA-Z0-9]+$")
		regexFactory.username = regexp.MustCompile("^[a-zA-Z]+[_a-zA-Z0-9.]+$")
		regexFactory.email = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		regexFactory.phone_iran = regexp.MustCompile("^[0][1-8][0-9]{9}$")
		regexFactory.mobile_iran = regexp.MustCompile("^[0][9][0-9]{9}$")
		regexFactory.string = regexp.MustCompile("^[a-zA-Z0-9؀-ۿ\\s\\-_.*@!']+$")
		regexFactory.persian = regexp.MustCompile("^[0-9\u0600-\u06FF\\s\\-_.*@!']+$")
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
func (r *RegexFactory) PhoneIran(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.phone_iran.MatchString(value)
}
func (r *RegexFactory) MobileIran(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.mobile_iran.MatchString(value)
}
func (r *RegexFactory) String(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.string.MatchString(value)
}
func (r *RegexFactory) Persian(value string) bool {
	if len(value) < 1 {
		return true
	}
	return r.persian.MatchString(value)
}
func (r *RegexFactory) IranNationalCode(value string) bool {
	for i := 0; i < 10; i++ {
		if value[i] < '0' || value[i] > '9' {
			return false
		}
	}
	check := int(value[9] - '0')
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(value[i]-'0') * (10 - i)
	}
	sum %= 11
	return (sum < 2 && check == sum) || (sum >= 2 && check+sum == 11)
}
