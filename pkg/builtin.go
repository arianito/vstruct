package vstruct

import (
	"encoding/json"
	"reflect"
	"strconv"
)

//noinspection GoUnusedGlobalVariable
var local = loadString()

func loadString() bool {
	RegisterRule("alpha", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().Alpha(ctx.FieldValue.String()) {
			return translate("alpha", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("alpha_num", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().AlphaNum(ctx.FieldValue.String()) {
			return translate("alpha_num", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("alpha_dash", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().AlphaDash(ctx.FieldValue.String()) {
			return translate("alpha_dash", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("min", Combine(reflect.String), func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln < parseInt(ctx.Args[0]) {
			return translate("min.string", attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("max", Combine(reflect.String), func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln < parseInt(ctx.Args[0]) {
			return translate("min.string", attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("between", Combine(reflect.String), func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if !(ln >= parseInt(ctx.Args[0]) && ln <= parseInt(ctx.Args[1])) {
			return translate("between.string", attribute(ctx.AliasName), ctx.Args[0], ctx.Args[1])
		}
		return ""
	})
	RegisterRule("json", Combine(reflect.String), func(ctx *Context) string {
		if !json.Valid([]byte(ctx.FieldValue.String())) {
			return translate("json", attribute(ctx.AliasName))
		}
		return ""
	})
	return true
}

func parseInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
