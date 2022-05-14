package vstruct

import (
	"fmt"
	"strings"
)

var lang = "en"
var messages = map[string]map[string]interface{}{}

func SetLocale(locale string) error {
	if _, ok := messages[locale]; !ok {
		return fmt.Errorf("locale not found")
	}
	lang = locale
	return nil
}
func RegisterLanguage(locale string, data map[string]interface{}) {
	messages[locale] = data
}

func RegisterAttributes(keys map[string]string) {
	obj := messages[lang]

	inner, ok := obj["attributes"]
	if !ok {
		obj["attributes"] = keys
		return
	}
	mp := inner.(map[string]interface{})
	for key, value := range keys {
		mp[key] = value
	}
}

func Append(keys map[string]interface{}) {
	obj := messages[lang]
	for key, value := range keys {
		obj[key] = value
	}
}

func Attribute(name string) string {
	obj := messages[lang]
	ok := eval("attributes."+name, obj)
	if ok == "" {
		return name
	}
	return ok
}

//Translate: translate using builtin dictionary
func Translate(key string, args ...interface{}) string {
	obj := messages[lang]
	msg := eval(key, obj)
	if msg == "" {
		return ""
	}
	j := 0
	li := 0
	ln := len(msg)
	st := false
	ot := ""
	aln := len(args)
	for i := 0; i < ln; i++ {
		if msg[i] == ':' {
			st = true
			ot += msg[li:i]
			li = i
		} else if (msg[i] == ' ' || msg[i] == ',' || msg[i] == '.' || i == ln-1) && st {
			if j < aln {
				ot += fmt.Sprintf("%v", args[j])
				li += len(msg[li:i])
				j++
			}
			st = false
		}
	}
	if li != ln-1 {
		ot += msg[li:ln]
	}
	return ot
}

func eval(path string, from map[string]interface{}) string {
	firstIndex := strings.Index(path, ".")
	var first string
	var rest string
	if firstIndex > -1 {
		first = path[:firstIndex]
		rest = path[firstIndex+1:]
	} else {
		first = path
		rest = ""
	}

	ifc := from[first]
	if value, ok := ifc.(string); ok {
		return value
	}
	if value, ok := ifc.(map[string]interface{}); ok && rest != "" {
		return eval(rest, value)
	}
	if value, ok := ifc.(map[string]string); ok && rest != "" {
		return value[rest]
	}
	return ""
}
