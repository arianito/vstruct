package vstruct

import (
	"reflect"
)

func loadAny() bool {
	RegisterRule("required", Combine(reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Array, reflect.Chan, reflect.Map), func(ctx *Context) string {
		if ctx.FieldValue.IsNil() || !ctx.FieldValue.IsValid() {
			return Translate("required", Attribute(ctx.AliasName))
		}
		return ""
	})
	return true
}
