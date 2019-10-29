package vstruct

import (
	"reflect"
)

func loadInt() bool {
	var numbers = Combine(reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64)

	RegisterRule("required", numbers, func(ctx *Context) string {
		if ctx.FieldValue.IsZero() {
			return translate("required", attribute(ctx.AliasName))
		}
		return ""
	})
	RegisterRule("min", numbers, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) {
			return translate("min.numeric", attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("max", numbers, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) {
			return translate("max.numeric", attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("between", numbers, func(ctx *Context) string {
		a := convertFloat(ctx.FieldValue.Interface())
		if a < parseFloat(ctx.Args[0]) || a > parseFloat(ctx.Args[1]) {
			return translate("between.numeric", attribute(ctx.AliasName), ctx.Args[0], ctx.Args[1])
		}
		return ""
	})


	RegisterRule("lt", numbers, func(ctx *Context) string {
		if convertFloat(ctx.FieldValue.Interface()) >= parseFloat(ctx.Args[0]) {
			return translate("lt.numeric", attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("lte", numbers, func(ctx *Context) string {
		if convertFloat(ctx.FieldValue.Interface()) > parseFloat(ctx.Args[0]) {
			return translate("lte.numeric", attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("gt", numbers, func(ctx *Context) string {
		if convertFloat(ctx.FieldValue.Interface()) <= parseFloat(ctx.Args[0]) {
			return translate("lte.numeric", attribute(ctx.AliasName), ctx.Args[0])
		}
		return ""
	})
	RegisterRule("gte", numbers, func(ctx *Context) string {
		if convertFloat(ctx.FieldValue.Interface()) < parseFloat(ctx.Args[0]) {
			return translate("lte.numeric", attribute(ctx.AliasName), ctx.Args[0])
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
		return translate("in", attribute(ctx.AliasName))
	})
	RegisterRule("not_in", numbers, func(ctx *Context) string {
		val := convertFloat(ctx.FieldValue.Interface())
		for _, item := range ctx.Args {
			if parseFloat(item) == val {
				return translate("not_in", attribute(ctx.AliasName))
			}
		}
		return ""
	})
	return true
}

//noinspection GoUnusedGlobalVariable
var local_int = loadInt()
