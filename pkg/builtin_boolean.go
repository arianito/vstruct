package vstruct

import (
	"reflect"
)

func loadBoolean() bool {
	RegisterRule("accepted", Combine(reflect.Bool), func(ctx *Context) string {
		if !ctx.FieldValue.Bool() {
			return Translate("accepted", Attribute(ctx.AliasName))
		}
		return ""
	})
	return true
}
