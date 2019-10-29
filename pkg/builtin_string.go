package vstruct

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
)


func loadString() bool {
	RegisterRule("required", Combine(reflect.String), func(ctx *Context) string {
		if len(ctx.FieldValue.String()) < 1 {
			return translate("required", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("filled", Combine(reflect.String), func(ctx *Context) string {
		if len(ctx.FieldValue.String()) < 1 {
			return translate("filled", attribute(ctx.AliasName))
		}
		return ""
	})
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
	RegisterRule("email", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().Email(ctx.FieldValue.String()) {
			return translate("email", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("phone_iran", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().PhoneIran(ctx.FieldValue.String()) {
			return translate("phone", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("mobile_iran", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().PhoneIran(ctx.FieldValue.String()) {
			return translate("phone", attribute(ctx.AliasName))
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
		if ln < parseInt(ctx.Args[0]) || ln > parseInt(ctx.Args[1]) {
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
	RegisterRule("confirmed", Combine(reflect.String), func(ctx *Context) string {
		if ctx.FieldValue.String() != ctx.InstanceValue.FieldByName(ctx.Args[0]).String() {
			return translate("confirmed", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("regex", Combine(reflect.String), func(ctx *Context) string {
		if !regexp.MustCompile(ctx.FieldValue.String()).MatchString(ctx.Args[0]) {
			return translate("regex", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("not_regex", Combine(reflect.String), func(ctx *Context) string {
		if regexp.MustCompile(ctx.FieldValue.String()).MatchString(ctx.Args[0]) {
			return translate("not_regex", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("string", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().String(ctx.FieldValue.String()) {
			return translate("string", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("persian", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().Persian(ctx.FieldValue.String()) {
			return translate("persian", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("ends_with", Combine(reflect.String), func(ctx *Context) string {
		val := ctx.FieldValue.String()
		nv := len(val)
		for _, arg := range ctx.Args {
			if strings.Index(val, arg) == nv-len(arg) {
				return ""
			}
		}
		return translate("ends_with", attribute(ctx.AliasName), strings.Join(ctx.Args, ", "))
	})
	RegisterRule("starts_with", Combine(reflect.String), func(ctx *Context) string {
		val := ctx.FieldValue.String()
		for _, arg := range ctx.Args {
			if strings.Index(val, arg) == 0 {
				return ""
			}
		}
		return translate("starts_with", attribute(ctx.AliasName), strings.Join(ctx.Args, ", "))
	})
	RegisterRule("national_code_iran", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().IranNationalCode(ctx.FieldValue.String()) {
			return translate("ssn", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("in", Combine(reflect.String), func(ctx *Context) string {
		val := ctx.FieldValue.String()
		for _, item := range ctx.Args {
			if item == val {
				return ""
			}
		}

		return translate("in", attribute(ctx.AliasName))
	})
	RegisterRule("not_in", Combine(reflect.String), func(ctx *Context) string {
		val := ctx.FieldValue.String()
		for _, item := range ctx.Args {
			if item == val {
				return translate("not_in", attribute(ctx.AliasName))
			}
		}
		return ""
	})
	return true
}



//noinspection GoUnusedGlobalVariable
var local_string = loadString()