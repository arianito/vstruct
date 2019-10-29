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
			return Translate("required", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("filled", Combine(reflect.String), func(ctx *Context) string {
		if len(ctx.FieldValue.String()) < 1 {
			return Translate("filled", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("alpha", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().Alpha(ctx.FieldValue.String()) {
			return Translate("alpha", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("alpha_num", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().AlphaNum(ctx.FieldValue.String()) {
			return Translate("alpha_num", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("alpha_dash", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().AlphaDash(ctx.FieldValue.String()) {
			return Translate("alpha_dash", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("email", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().Email(ctx.FieldValue.String()) {
			return Translate("email", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("phone_iran", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().PhoneIran(ctx.FieldValue.String()) {
			return Translate("phone", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("mobile_iran", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().PhoneIran(ctx.FieldValue.String()) {
			return Translate("phone", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("min", Combine(reflect.String), func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln < parseInt(ctx.Args[0]) {
			return Translate("min.string", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("max", Combine(reflect.String), func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln < parseInt(ctx.Args[0]) {
			return Translate("min.string", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("size", Combine(reflect.String), func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln == parseInt(ctx.Args[0]) {
			return Translate("size.string", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("between", Combine(reflect.String), func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln < parseInt(ctx.Args[0]) || ln > parseInt(ctx.Args[1]) {
			return Translate("between.string", Attribute(ctx.AliasName), ctx.Args[0], ctx.Args[1])
		}
		return ""
	})
	RegisterRule("json", Combine(reflect.String), func(ctx *Context) string {
		if !json.Valid([]byte(ctx.FieldValue.String())) {
			return Translate("json", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("confirmed", Combine(reflect.String), func(ctx *Context) string {
		if ctx.FieldValue.String() != ctx.InstanceValue.FieldByName(ctx.Args[0]).String() {
			return Translate("confirmed", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("regex", Combine(reflect.String), func(ctx *Context) string {
		if !regexp.MustCompile(ctx.FieldValue.String()).MatchString(ctx.Args[0]) {
			return Translate("regex", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("not_regex", Combine(reflect.String), func(ctx *Context) string {
		if regexp.MustCompile(ctx.FieldValue.String()).MatchString(ctx.Args[0]) {
			return Translate("not_regex", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("string", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().String(ctx.FieldValue.String()) {
			return Translate("string", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("persian", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().Persian(ctx.FieldValue.String()) {
			return Translate("persian", Attribute(ctx.AliasName))
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
		return Translate("ends_with", Attribute(ctx.AliasName), strings.Join(ctx.Args, ", "))
	})
	RegisterRule("starts_with", Combine(reflect.String), func(ctx *Context) string {
		val := ctx.FieldValue.String()
		for _, arg := range ctx.Args {
			if strings.Index(val, arg) == 0 {
				return ""
			}
		}
		return Translate("starts_with", Attribute(ctx.AliasName), strings.Join(ctx.Args, ", "))
	})
	RegisterRule("national_code_iran", Combine(reflect.String), func(ctx *Context) string {
		if !GetRegex().IranNationalCode(ctx.FieldValue.String()) {
			return Translate("ssn", Attribute(ctx.AliasName))
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

		return Translate("in", Attribute(ctx.AliasName))
	})
	RegisterRule("not_in", Combine(reflect.String), func(ctx *Context) string {
		val := ctx.FieldValue.String()
		for _, item := range ctx.Args {
			if item == val {
				return Translate("not_in", Attribute(ctx.AliasName))
			}
		}
		return ""
	})
	return true
}
