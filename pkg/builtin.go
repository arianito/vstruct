package vstruct

import (
	"reflect"
	"strconv"
)

//noinspection GoUnusedGlobalVariable
var local = loadString()

func loadString() bool {
	RegisterRule("alpha", Combine(reflect.String), func(fieldName string, value reflect.Value, args ...string) string {
		if !GetRegex().ValidateEnglishLetters(value.String()) {
			return translate("alpha", fieldName)
		}
		return ""
	})
	RegisterRule("min", Combine(reflect.String), func(fieldName string, value reflect.Value, args ...string) string {
		if len(value.String()) < parseInt(args[0]) {
			return translate("min.string", fieldName, args[0])
		}
		return ""
	})
	return true
}

func parseInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
