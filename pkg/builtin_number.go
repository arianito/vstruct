package vstruct

import (
	"reflect"
)

func loadInt() bool {
	var numbers = Combine(reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64)

	RegisterRule("required", numbers, func(ctx *Context) string {
		if ctx.FieldValue.IsZero() {
			return Translate("required", Attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("min", numbers, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) {
			return Translate("min.numeric", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("max", numbers, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) {
			return Translate("max.numeric", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("between", numbers, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) || a > parseFloat(ctx.Args[1]) {
			return Translate("between.numeric", Attribute(ctx.AliasName), ctx.Args[0], ctx.Args[1])
		}
		return ""
	})

	RegisterRule("lt", numbers, func(ctx *Context) string {
		if convertFloat(ctx.FieldValue.Interface()) >= parseFloat(ctx.Args[0]) {
			return Translate("lt.numeric", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("lte", numbers, func(ctx *Context) string {
		if convertFloat(ctx.FieldValue.Interface()) > parseFloat(ctx.Args[0]) {
			return Translate("lte.numeric", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("gt", numbers, func(ctx *Context) string {
		if convertFloat(ctx.FieldValue.Interface()) <= parseFloat(ctx.Args[0]) {
			return Translate("lte.numeric", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("gte", numbers, func(ctx *Context) string {
		if convertFloat(ctx.FieldValue.Interface()) < parseFloat(ctx.Args[0]) {
			return Translate("lte.numeric", Attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("in", numbers, func(ctx *Context) string {
		val := convertFloat(ctx.FieldValue.Interface())
		for _, item := range ctx.Args {
			if parseFloat(item) == val {
				return ""
			}
		}
		return Translate("in", Attribute(ctx.AliasName))
	})
	RegisterRule("not_in", numbers, func(ctx *Context) string {
		val := convertFloat(ctx.FieldValue.Interface())
		for _, item := range ctx.Args {
			if parseFloat(item) == val {
				return Translate("not_in", Attribute(ctx.AliasName))
			}
		}
		return ""
	})
	return true
}
