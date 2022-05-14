package vstruct

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
)

func LoadBuiltin() {

	RegisterLanguage("en", map[string]interface{}{
		"accepted":    "The :attribute must be accepted.",
		"alpha":       "The :attribute may only contain letters.",
		"alpha_dash":  "The :attribute may only contain letters, numbers, dashes and underscores.",
		"alpha_num":   "The :attribute may only contain letters and numbers.",
		"username":    "The :attribute may only contain letters, underscores, dots and numbers.",
		"boolean":     "The :attribute field must be true or false.",
		"confirmed":   "The :attribute confirmation does not match.",
		"email":       "The :attribute must be a valid email address.",
		"ends_with":   "The :attribute must end with one of the following, :values",
		"filled":      "The :attribute field must have a value.",
		"in":          "The selected :attribute is invalid.",
		"json":        "The :attribute must be a valid JSON string.",
		"not_in":      "The selected :attribute is invalid.",
		"not_regex":   "The :attribute format is invalid.",
		"regex":       "The :attribute format is invalid.",
		"required":    "The :attribute field is required.",
		"starts_with": "The :attribute must start with one of the following, :values",
		"string":      "The :attribute must be a string.",
		"phone":       "The :attribute must be a valid phone number.",
		"mobile":      "The :attribute must be a valid mobile number.",
		"persian":     "The :attribute may only contain persian letters.",
		"ssn":         "The :attribute must be a valid social security number.",
		"max": map[string]interface{}{
			"numeric": "The :attribute may not be greater than :max.",
			"string":  "The :attribute may not be greater than :max characters.",
		},
		"min": map[string]interface{}{
			"numeric": "The :attribute must be at least :min.",
			"string":  "The :attribute must be at least :min characters.",
		},
		"between": map[string]interface{}{
			"numeric": "The :attribute must be between :min and :max.",
			"string":  "The :attribute must be between :min and :max characters.",
		},
		"size": map[string]interface{}{
			"string": "The :attribute must be :value characters.",
		},
	})

	anyKind := Combine(reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Array, reflect.Chan, reflect.Map)
	booleanKind := Combine(reflect.Bool)
	stringKind := Combine(reflect.String)
	numberKind := Combine(reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64)

	// any validators
	RegisterRule("required", anyKind, func(ctx *Context) string {
		if ctx.FieldValue.IsNil() || !ctx.FieldValue.IsValid() {
			return Translate("required", Attribute(ctx.AliasName))
		}
		return ""
	})

	// boolean validators
	RegisterRule("accepted", booleanKind, func(ctx *Context) string {
		if !ctx.FieldValue.Bool() {
			return Translate("accepted", Attribute(ctx.AliasName))
		}
		return ""
	})

	// number validators
	RegisterRule("required", numberKind, func(ctx *Context) string {
		if ctx.FieldValue.IsZero() {
			return Translate("required", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("min", numberKind, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) {
			return Translate("min.numeric", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("max", numberKind, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) {
			return Translate("max.numeric", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("between", numberKind, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) || a > parseFloat(ctx.Args[1]) {
			return Translate("between.numeric", Attribute(ctx.AliasName), ctx.Args[0], ctx.Args[1])
		}
		return ""
	})
	RegisterRule("in", numberKind, func(ctx *Context) string {
		val := convertFloat(ctx.FieldValue.Interface())
		for _, item := range ctx.Args {
			if parseFloat(item) == val {
				return ""
			}
		}
		return Translate("in", Attribute(ctx.AliasName))
	})
	RegisterRule("not_in", numberKind, func(ctx *Context) string {
		val := convertFloat(ctx.FieldValue.Interface())
		for _, item := range ctx.Args {
			if parseFloat(item) == val {
				return Translate("not_in", Attribute(ctx.AliasName))
			}
		}
		return ""
	})

	// string validators
	RegisterRule("required", stringKind, func(ctx *Context) string {
		if len(ctx.FieldValue.String()) < 1 {
			return Translate("required", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("filled", stringKind, func(ctx *Context) string {
		if len(ctx.FieldValue.String()) < 1 {
			return Translate("filled", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("alpha", stringKind, func(ctx *Context) string {
		if !GetRegex().Alpha(ctx.FieldValue.String()) {
			return Translate("alpha", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("alpha_num", stringKind, func(ctx *Context) string {
		if !GetRegex().AlphaNum(ctx.FieldValue.String()) {
			return Translate("alpha_num", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("alpha_dash", stringKind, func(ctx *Context) string {
		if !GetRegex().AlphaDash(ctx.FieldValue.String()) {
			return Translate("alpha_dash", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("username", stringKind, func(ctx *Context) string {
		if !GetRegex().Username(ctx.FieldValue.String()) {
			return Translate("username", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("email", stringKind, func(ctx *Context) string {
		if !GetRegex().Email(ctx.FieldValue.String()) {
			return Translate("email", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("min", stringKind, func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln < parseInt(ctx.Args[0]) {
			return Translate("min.string", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("max", stringKind, func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln < parseInt(ctx.Args[0]) {
			return Translate("min.string", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("size", stringKind, func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln == parseInt(ctx.Args[0]) {
			return Translate("size.string", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("between", stringKind, func(ctx *Context) string {
		ln := len(ctx.FieldValue.String())
		if ln < parseInt(ctx.Args[0]) || ln > parseInt(ctx.Args[1]) {
			return Translate("between.string", Attribute(ctx.AliasName), ctx.Args[0], ctx.Args[1])
		}
		return ""
	})
	RegisterRule("json", stringKind, func(ctx *Context) string {
		if !json.Valid([]byte(ctx.FieldValue.String())) {
			return Translate("json", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("confirmed", stringKind, func(ctx *Context) string {
		if ctx.FieldValue.String() != ctx.InstanceValue.FieldByName(ctx.Args[0]).String() {
			return Translate("confirmed", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("regex", stringKind, func(ctx *Context) string {
		if !regexp.MustCompile(ctx.FieldValue.String()).MatchString(ctx.Args[0]) {
			return Translate("regex", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("not_regex", stringKind, func(ctx *Context) string {
		if regexp.MustCompile(ctx.FieldValue.String()).MatchString(ctx.Args[0]) {
			return Translate("not_regex", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("ends_with", stringKind, func(ctx *Context) string {
		val := ctx.FieldValue.String()
		nv := len(val)
		for _, arg := range ctx.Args {
			if strings.Index(val, arg) == nv-len(arg) {
				return ""
			}
		}
		return Translate("ends_with", Attribute(ctx.AliasName), strings.Join(ctx.Args, ", "))
	})
	RegisterRule("starts_with", stringKind, func(ctx *Context) string {
		val := ctx.FieldValue.String()
		for _, arg := range ctx.Args {
			if strings.Index(val, arg) == 0 {
				return ""
			}
		}
		return Translate("starts_with", Attribute(ctx.AliasName), strings.Join(ctx.Args, ", "))
	})
	RegisterRule("in", stringKind, func(ctx *Context) string {
		val := ctx.FieldValue.String()
		for _, item := range ctx.Args {
			if item == val {
				return ""
			}
		}

		return Translate("in", Attribute(ctx.AliasName))
	})
	RegisterRule("not_in", stringKind, func(ctx *Context) string {
		val := ctx.FieldValue.String()
		for _, item := range ctx.Args {
			if item == val {
				return Translate("not_in", Attribute(ctx.AliasName))
			}
		}
		return ""
	})
}
